// Package store is the only thing that talks to SQLite. It holds attempts and the
// leaderboard; the study material itself stays in files.
package store

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite" // pure-Go driver: the release build cross-compiles with CGO off
)

//go:embed schema.sql
var schema string

// ErrNoSession is returned when a session id does not exist.
var ErrNoSession = errors.New("no such quiz session")

// DB wraps the connection with the handful of queries the app needs.
type DB struct{ sql *sql.DB }

// Open connects to path, creating and migrating it if needed.
func Open(path string) (*DB, error) {
	// WAL keeps a reader on the dashboard from blocking someone finishing a quiz,
	// and busy_timeout turns the rare write collision into a short wait instead of
	// an error.
	conn, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	if _, err := conn.Exec(schema); err != nil {
		conn.Close()
		return nil, fmt.Errorf("applying schema: %w", err)
	}
	return &DB{sql: conn}, nil
}

func (d *DB) Close() error { return d.sql.Close() }

// Player returns the id for a name, creating the player on first sight. Names are
// trimmed and capped, because they end up on a shared leaderboard.
func (d *DB) Player(name string) (int64, string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, "", errors.New("a name is required")
	}
	if len([]rune(name)) > 24 {
		name = string([]rune(name)[:24])
	}
	if _, err := d.sql.Exec(
		`INSERT INTO players (name, created_at) VALUES (?, ?) ON CONFLICT (name) DO NOTHING`,
		name, time.Now().Unix()); err != nil {
		return 0, "", err
	}
	var id int64
	var stored string
	// The stored spelling wins, so someone typing "WILHELM" tonight does not rename
	// yesterday's entries.
	err := d.sql.QueryRow(`SELECT id, name FROM players WHERE name = ?`, name).Scan(&id, &stored)
	return id, stored, err
}

// Session is a quiz run as the API reports it.
type Session struct {
	ID         string `json:"id"`
	Player     string `json:"player"`
	Mode       string `json:"mode"`
	Topic      string `json:"topic,omitempty"`
	Total      int    `json:"total"`
	Score      int    `json:"score"`
	DurationMS int64  `json:"durationMs"`
	StartedAt  int64  `json:"startedAt"`
	FinishedAt *int64 `json:"finishedAt,omitempty"`
}

// CreateSession records a run and the exact questions it was served.
func (d *DB) CreateSession(id string, playerID int64, mode, topic string, total int, payload string) error {
	_, err := d.sql.Exec(
		`INSERT INTO sessions (id, player_id, mode, topic, total, payload, started_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, playerID, mode, topic, total, payload, time.Now().UnixMilli())
	return err
}

// Payload returns the questions a session was served, with their answer keys.
func (d *DB) Payload(sessionID string) (string, error) {
	var payload string
	err := d.sql.QueryRow(`SELECT payload FROM sessions WHERE id = ?`, sessionID).Scan(&payload)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrNoSession
	}
	return payload, err
}

// RecordAnswer saves one graded answer. Answering the same question twice in a session
// keeps the first attempt: the reader has already seen the explanation, so a second
// try would measure nothing.
func (d *DB) RecordAnswer(sessionID, questionID, topic, given string, correct, confident bool, ms int64) error {
	_, err := d.sql.Exec(
		`INSERT INTO answers (session_id, question_id, topic, given, correct, confident, ms, answered_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		 ON CONFLICT (session_id, question_id) DO NOTHING`,
		sessionID, questionID, topic, given, correct, confident, ms, time.Now().UnixMilli())
	return err
}

// Answered reports how many questions of a session have been answered already.
func (d *DB) Answered(sessionID string) (int, error) {
	var n int
	err := d.sql.QueryRow(`SELECT count(*) FROM answers WHERE session_id = ?`, sessionID).Scan(&n)
	return n, err
}

// Finish closes a session and scores it from the answers on record, so the score on the
// leaderboard is one the server computed rather than one the browser reported.
func (d *DB) Finish(sessionID string) (Session, error) {
	var started int64
	var s Session
	err := d.sql.QueryRow(
		`SELECT sessions.id, players.name, mode, topic, total, started_at
		   FROM sessions JOIN players ON players.id = sessions.player_id
		  WHERE sessions.id = ?`, sessionID).
		Scan(&s.ID, &s.Player, &s.Mode, &s.Topic, &s.Total, &started)
	if errors.Is(err, sql.ErrNoRows) {
		return s, ErrNoSession
	}
	if err != nil {
		return s, err
	}

	var score int
	if err := d.sql.QueryRow(
		`SELECT count(*) FROM answers WHERE session_id = ? AND correct = 1`, sessionID).Scan(&score); err != nil {
		return s, err
	}
	now := time.Now().UnixMilli()
	// A session already finished keeps its original time: refreshing the results page
	// must not inflate anyone's duration.
	if _, err := d.sql.Exec(
		`UPDATE sessions SET score = ?, finished_at = coalesce(finished_at, ?),
		        duration_ms = CASE WHEN finished_at IS NULL THEN ? - started_at ELSE duration_ms END
		  WHERE id = ?`, score, now, now, sessionID); err != nil {
		return s, err
	}
	err = d.sql.QueryRow(
		`SELECT score, duration_ms, started_at, finished_at FROM sessions WHERE id = ?`, sessionID).
		Scan(&s.Score, &s.DurationMS, &s.StartedAt, &s.FinishedAt)
	return s, err
}

// Get returns one finished or in-flight session.
func (d *DB) Get(sessionID string) (Session, error) {
	var s Session
	err := d.sql.QueryRow(
		`SELECT sessions.id, players.name, mode, topic, total, score, duration_ms, started_at, finished_at
		   FROM sessions JOIN players ON players.id = sessions.player_id
		  WHERE sessions.id = ?`, sessionID).
		Scan(&s.ID, &s.Player, &s.Mode, &s.Topic, &s.Total, &s.Score, &s.DurationMS, &s.StartedAt, &s.FinishedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return s, ErrNoSession
	}
	return s, err
}

// Leaderboard returns finished runs, best first. Accuracy ranks above raw score so a
// clean ten does not lose to a sloppy forty, and speed breaks the tie.
//
// Runs shorter than five questions are left out: they are too easy to farm.
func (d *DB) Leaderboard(mode string, limit int) ([]Session, error) {
	query := `SELECT sessions.id, players.name, mode, topic, total, score, duration_ms, started_at, finished_at
	            FROM sessions JOIN players ON players.id = sessions.player_id
	           WHERE finished_at IS NOT NULL AND total >= 5`
	args := []any{}
	if mode != "" && mode != "all" {
		query += ` AND mode = ?`
		args = append(args, mode)
	}
	query += ` ORDER BY (CAST(score AS REAL) / total) DESC, score DESC, duration_ms ASC LIMIT ?`
	args = append(args, limit)
	return d.sessions(query, args...)
}

// Recent returns the latest finished runs, whatever they scored, so the dashboard shows
// that other people are actually studying.
func (d *DB) Recent(limit int) ([]Session, error) {
	return d.sessions(
		`SELECT sessions.id, players.name, mode, topic, total, score, duration_ms, started_at, finished_at
		   FROM sessions JOIN players ON players.id = sessions.player_id
		  WHERE finished_at IS NOT NULL
		  ORDER BY finished_at DESC LIMIT ?`, limit)
}

func (d *DB) sessions(query string, args ...any) ([]Session, error) {
	rows, err := d.sql.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Session{}
	for rows.Next() {
		var s Session
		if err := rows.Scan(&s.ID, &s.Player, &s.Mode, &s.Topic, &s.Total, &s.Score,
			&s.DurationMS, &s.StartedAt, &s.FinishedAt); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, rows.Err()
}

// TopicStat is one row of a player's coverage heatmap.
type TopicStat struct {
	Topic    string `json:"topic"`
	Seen     int    `json:"seen"`
	Correct  int    `json:"correct"`
	LastSeen int64  `json:"lastSeen"`
}

// TopicStats returns per-topic accuracy for a player, counting each question once
// (its most recent attempt) so repeating one item does not dominate the heatmap.
func (d *DB) TopicStats(name string) ([]TopicStat, error) {
	rows, err := d.sql.Query(
		`SELECT topic, count(*), sum(correct), max(answered_at) FROM (
		     SELECT a.topic AS topic, a.correct AS correct, a.answered_at AS answered_at,
		            row_number() OVER (PARTITION BY a.question_id ORDER BY a.answered_at DESC) AS rn
		       FROM answers a
		       JOIN sessions s ON s.id = a.session_id
		       JOIN players  p ON p.id = s.player_id
		      WHERE p.name = ?
		 ) WHERE rn = 1 GROUP BY topic`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []TopicStat{}
	for rows.Next() {
		var t TopicStat
		if err := rows.Scan(&t.Topic, &t.Seen, &t.Correct, &t.LastSeen); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// Miss is a question a player has got wrong.
type Miss struct {
	QuestionID string `json:"questionId"`
	Topic      string `json:"topic"`
	Wrong      int    `json:"wrong"`
	Attempts   int    `json:"attempts"`
	// Confident counts the times they were sure and still wrong. Those beliefs are
	// the expensive ones, so weak-spot practice puts them first.
	Confident int   `json:"confident"`
	LastSeen  int64 `json:"lastSeen"`
}

// Misses returns a player's wrong answers, the confidently-wrong ones first. Generated
// drill questions are excluded — they never repeat, so there is nothing to revisit.
func (d *DB) Misses(name string, limit int) ([]Miss, error) {
	rows, err := d.sql.Query(
		`SELECT a.question_id, a.topic,
		        sum(CASE WHEN a.correct = 0 THEN 1 ELSE 0 END) AS wrong,
		        count(*) AS attempts,
		        sum(CASE WHEN a.correct = 0 AND a.confident = 1 THEN 1 ELSE 0 END) AS confident,
		        max(a.answered_at)
		   FROM answers a
		   JOIN sessions s ON s.id = a.session_id
		   JOIN players  p ON p.id = s.player_id
		  WHERE p.name = ? AND a.question_id NOT LIKE 'gen:%'
		  GROUP BY a.question_id
		 HAVING wrong > 0
		  ORDER BY confident DESC, wrong DESC, max(a.answered_at) DESC
		  LIMIT ?`, name, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Miss{}
	for rows.Next() {
		var m Miss
		if err := rows.Scan(&m.QuestionID, &m.Topic, &m.Wrong, &m.Attempts, &m.Confident, &m.LastSeen); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// Mastered returns the questions a player has answered correctly and confidently, so
// fresh practice can prefer material they have not nailed yet.
func (d *DB) Mastered(name string) (map[string]bool, error) {
	rows, err := d.sql.Query(
		`SELECT a.question_id FROM answers a
		   JOIN sessions s ON s.id = a.session_id
		   JOIN players  p ON p.id = s.player_id
		  WHERE p.name = ? AND a.correct = 1 AND a.confident = 1
		  GROUP BY a.question_id`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string]bool{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out[id] = true
	}
	return out, rows.Err()
}

// Listen is how far through an episode someone is.
type Listen struct {
	Episode     string  `json:"episode"`
	PositionSec float64 `json:"positionSec"`
	ListenedSec float64 `json:"listenedSec"`
	UpdatedAt   int64   `json:"updatedAt"`
}

// RecordListening moves the resume point and adds to the time listened.
//
// The delta is accumulated rather than set, because the honest measure of listening is
// time actually spent playing — scrubbing to the end of an episode is not listening to
// it. The client sends small deltas while playing, so a seek adds nothing.
func (d *DB) RecordListening(playerID int64, episode string, position, delta float64) error {
	if delta < 0 {
		delta = 0
	}
	_, err := d.sql.Exec(
		`INSERT INTO listening (player_id, episode, position_sec, listened_sec, updated_at)
		 VALUES (?, ?, ?, ?, ?)
		 ON CONFLICT (player_id, episode) DO UPDATE SET
		     position_sec = excluded.position_sec,
		     listened_sec = listening.listened_sec + excluded.listened_sec,
		     updated_at   = excluded.updated_at`,
		playerID, episode, position, delta, time.Now().UnixMilli())
	return err
}

// Listening returns a player's progress through every episode they have started.
func (d *DB) Listening(name string) ([]Listen, error) {
	rows, err := d.sql.Query(
		`SELECT episode, position_sec, listened_sec, updated_at
		   FROM listening l JOIN players p ON p.id = l.player_id
		  WHERE p.name = ?`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Listen{}
	for rows.Next() {
		var l Listen
		if err := rows.Scan(&l.Episode, &l.PositionSec, &l.ListenedSec, &l.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, l)
	}
	return out, rows.Err()
}
