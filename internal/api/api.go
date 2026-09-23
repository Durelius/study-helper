// Package api exposes the study material and the leaderboard as JSON.
//
// One rule shapes the whole package: an answer key never leaves the server before the
// reader has committed to an answer. Everything else — the leaderboard, the accuracy
// figures, the weak-spot list — is only meaningful because of that.
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wilhelmdurelius/chulastudy/internal/audio"
	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/quiz"
	"github.com/wilhelmdurelius/chulastudy/internal/store"
)

// Server holds everything the handlers need.
type Server struct {
	Set   *content.Set
	DB    *store.DB
	Log   *slog.Logger
	Exam  time.Time
	Audio *audio.Library
}

// New returns a Server.
func New(set *content.Set, db *store.DB, log *slog.Logger, exam time.Time, lib *audio.Library) *Server {
	return &Server{Set: set, DB: db, Log: log, Exam: exam, Audio: lib}
}

// Routes registers the API.
func (s *Server) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/course", s.handleCourse)
	mux.HandleFunc("GET /api/topics", s.handleTopics)
	mux.HandleFunc("GET /api/modes", s.handleModes)
	mux.HandleFunc("GET /api/notes/{topic}", s.handleNotes)
	mux.HandleFunc("GET /api/leaderboard", s.handleLeaderboard)
	mux.HandleFunc("GET /api/stats", s.handleStats)
	mux.HandleFunc("GET /api/plan", s.handlePlan)
	mux.HandleFunc("POST /api/quiz", s.handleStartQuiz)
	mux.HandleFunc("POST /api/quiz/{id}/answer", s.handleAnswer)
	mux.HandleFunc("POST /api/quiz/{id}/finish", s.handleFinish)
	mux.HandleFunc("GET /api/quiz/{id}", s.handleGetQuiz)
	mux.HandleFunc("POST /api/check", s.handleCheck)
	mux.HandleFunc("GET /api/episodes", s.handleEpisodes)
	mux.HandleFunc("POST /api/listen", s.handleListen)
}

// publicQuestion is a question as the browser is allowed to see it: no answer key, no
// explanation. Both arrive one at a time, after the answer is in.
type publicQuestion struct {
	ID         string          `json:"id"`
	Topic      string          `json:"topic"`
	TopicTitle string          `json:"topicTitle"`
	Type       string          `json:"type"`
	Stem       string          `json:"stem"`
	Choices    []string        `json:"choices"`
	Difficulty int             `json:"difficulty"`
	ExamFocus  bool            `json:"examFocus"`
	Data       json.RawMessage `json:"data,omitempty"`
}

func (s *Server) public(q content.Question) publicQuestion {
	return publicQuestion{
		ID: q.ID, Topic: q.Topic, TopicTitle: s.Set.Title(q.Topic), Type: q.Type,
		Stem: q.Stem, Choices: q.Choices, Difficulty: q.Difficulty,
		ExamFocus: q.ExamFocus, Data: q.Data,
	}
}

func (s *Server) handleModes(w http.ResponseWriter, r *http.Request) {
	// A mode with nothing to draw on is worse than a missing one: it looks available
	// and then fails. Report how many questions back each so the UI can grey it out.
	type modeInfo struct {
		quiz.Mode
		Available int `json:"available"`
	}
	out := []modeInfo{}
	for _, m := range quiz.Modes(s.Set.Course()) {
		out = append(out, modeInfo{Mode: m, Available: quiz.Available(s.Set, m)})
	}
	writeJSON(w, http.StatusOK, out)
}

// handleCourse tells the browser which course it is serving, so the title, the course
// code and the countdown are not compiled into the frontend.
func (s *Server) handleCourse(w http.ResponseWriter, r *http.Request) {
	c := s.Set.Course()
	writeJSON(w, http.StatusOK, map[string]any{
		"id": c.ID, "title": c.Title, "code": c.Code, "shortName": c.ShortName,
		"exam": s.Exam.Format(time.RFC3339),
	})
}

func (s *Server) handleTopics(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimSpace(r.URL.Query().Get("player"))
	type topicView struct {
		content.Topic
		Seen     int   `json:"seen"`
		Correct  int   `json:"correct"`
		LastSeen int64 `json:"lastSeen"`
	}
	stats := map[string]store.TopicStat{}
	if player != "" {
		rows, err := s.DB.TopicStats(player)
		if err != nil {
			s.fail(w, err)
			return
		}
		for _, row := range rows {
			stats[row.Topic] = row
		}
	}
	out := []topicView{}
	for _, t := range s.Set.Topics {
		st := stats[t.ID]
		out = append(out, topicView{Topic: t, Seen: st.Seen, Correct: st.Correct, LastSeen: st.LastSeen})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleNotes(w http.ResponseWriter, r *http.Request) {
	topic := r.PathValue("topic")
	lec, ok := s.Set.Lectures[topic]
	if !ok {
		writeError(w, http.StatusNotFound, "No notes for that topic yet.")
		return
	}
	// The inline "?check id=" markers point at questions; send those along so the
	// reader can answer them without a round trip per card.
	checks := map[string]publicQuestion{}
	for _, line := range strings.Split(lec.Body, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "?check id=") {
			continue
		}
		id := strings.TrimSpace(strings.TrimPrefix(line, "?check id="))
		if q, ok := s.Set.Questions[id]; ok {
			checks[id] = s.public(q)
		}
	}
	writeJSON(w, http.StatusOK, struct {
		content.Lecture
		Checks map[string]publicQuestion `json:"checks"`
	}{lec, checks})
}

func (s *Server) handleLeaderboard(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	limit := intParam(r, "limit", 20, 100)
	top, err := s.DB.Leaderboard(mode, limit)
	if err != nil {
		s.fail(w, err)
		return
	}
	recent, err := s.DB.Recent(10)
	if err != nil {
		s.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"top": top, "recent": recent})
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimSpace(r.URL.Query().Get("player"))
	if player == "" {
		writeError(w, http.StatusBadRequest, "Who is asking? Pass ?player=")
		return
	}
	topics, err := s.DB.TopicStats(player)
	if err != nil {
		s.fail(w, err)
		return
	}
	misses, err := s.DB.Misses(player, 40)
	if err != nil {
		s.fail(w, err)
		return
	}
	// A miss is only useful with the question attached — the point is to re-read the
	// thing you got wrong, not to see an id.
	type missView struct {
		store.Miss
		Question    publicQuestion `json:"question"`
		Explanation string         `json:"explanation"`
		Source      content.Source `json:"source"`
	}
	out := []missView{}
	for _, m := range misses {
		q, ok := s.Set.Questions[m.QuestionID]
		if !ok {
			continue
		}
		out = append(out, missView{Miss: m, Question: s.public(q), Explanation: q.Explanation, Source: q.Source})
	}
	writeJSON(w, http.StatusOK, map[string]any{"topics": topics, "misses": out})
}

type startRequest struct {
	Player string `json:"player"`
	Mode   string `json:"mode"`
	Topic  string `json:"topic"`
	Count  int    `json:"count"`
}

func (s *Server) handleStartQuiz(w http.ResponseWriter, r *http.Request) {
	var req startRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "That request did not parse.")
		return
	}
	mode, ok := quiz.ModeByID(s.Set.Course(), req.Mode)
	if !ok {
		writeError(w, http.StatusBadRequest, "Unknown quiz mode.")
		return
	}
	playerID, name, err := s.DB.Player(req.Player)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	hist := quiz.History{Mastered: map[string]bool{}}
	if misses, err := s.DB.Misses(name, 60); err == nil {
		for _, m := range misses {
			hist.Misses = append(hist.Misses, m.QuestionID)
		}
	}
	if mastered, err := s.DB.Mastered(name); err == nil {
		hist.Mastered = mastered
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	questions, err := quiz.Build(s.Set, mode, req.Topic, req.Count, hist, rnd)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	payload, err := json.Marshal(questions)
	if err != nil {
		s.fail(w, err)
		return
	}
	id := newID(rnd)
	if err := s.DB.CreateSession(id, playerID, mode.ID, req.Topic, len(questions), string(payload)); err != nil {
		s.fail(w, err)
		return
	}

	out := make([]publicQuestion, 0, len(questions))
	for _, q := range questions {
		out = append(out, s.public(q))
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"id": id, "player": name, "mode": mode, "questions": out,
	})
}

// handleGetQuiz re-serves a run after a refresh. It hands back the questions without
// their keys, exactly as the start call did, so reloading mid-quiz is not a way to see
// the answers.
func (s *Server) handleGetQuiz(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sess, err := s.DB.Get(id)
	if errors.Is(err, store.ErrNoSession) {
		writeError(w, http.StatusNotFound, "That quiz has expired or never existed.")
		return
	}
	if err != nil {
		s.fail(w, err)
		return
	}
	payload, err := s.DB.Payload(id)
	if err != nil {
		s.fail(w, err)
		return
	}
	var questions []content.Question
	if err := json.Unmarshal([]byte(payload), &questions); err != nil {
		s.fail(w, err)
		return
	}
	out := make([]publicQuestion, 0, len(questions))
	for _, q := range questions {
		out = append(out, s.public(q))
	}
	answered, err := s.DB.Answered(id)
	if err != nil {
		s.fail(w, err)
		return
	}
	mode, _ := quiz.ModeByID(s.Set.Course(), sess.Mode)
	writeJSON(w, http.StatusOK, map[string]any{
		"id": sess.ID, "player": sess.Player, "mode": mode,
		"questions": out, "answered": make([]string, answered), "session": sess,
	})
}

// handleCheck grades one question with no session behind it, for the retrieval prompts
// dropped into the lecture notes.
//
// Nothing is recorded: these exist to interrupt reading with a moment of recall, and
// counting them would let someone farm the same three questions into a perfect topic
// score without ever sitting a quiz.
func (s *Server) handleCheck(w http.ResponseWriter, r *http.Request) {
	var req answerRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "That request did not parse.")
		return
	}
	q, ok := s.Set.Questions[req.QuestionID]
	if !ok {
		writeError(w, http.StatusNotFound, "No such question.")
		return
	}
	correct, _ := s.grade(q, req)
	writeJSON(w, http.StatusOK, map[string]any{
		"correct":     correct,
		"answer":      q.Answer,
		"accept":      q.Accept,
		"explanation": q.Explanation,
		"source":      q.Source,
	})
}

type answerRequest struct {
	QuestionID string `json:"questionId"`
	Given      []int  `json:"given"`
	// Text carries a free-text answer. Graded liberally — see internal/content/text.go.
	Text      string `json:"text"`
	Confident bool   `json:"confident"`
	MS        int64  `json:"ms"`
}

func (s *Server) handleAnswer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req answerRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "That request did not parse.")
		return
	}
	q, err := s.question(id, req.QuestionID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	correct, given := s.grade(q, req)
	if err := s.DB.RecordAnswer(id, q.ID, q.Topic, given, correct, req.Confident, req.MS); err != nil {
		s.fail(w, err)
		return
	}
	// Only now does the key go out, together with the reason — which is the part that
	// actually changes what the reader knows.
	writeJSON(w, http.StatusOK, map[string]any{
		"correct":     correct,
		"answer":      q.Answer,
		"accept":      q.Accept,
		"explanation": q.Explanation,
		"source":      q.Source,
	})
}

func (s *Server) handleFinish(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sess, err := s.DB.Finish(id)
	if errors.Is(err, store.ErrNoSession) {
		writeError(w, http.StatusNotFound, "That quiz has expired or never existed.")
		return
	}
	if err != nil {
		s.fail(w, err)
		return
	}

	// The review list is the most useful screen in the app, so it carries everything:
	// the question, what the reader chose, the key and the explanation.
	payload, err := s.DB.Payload(id)
	if err != nil {
		s.fail(w, err)
		return
	}
	var questions []content.Question
	if err := json.Unmarshal([]byte(payload), &questions); err != nil {
		s.fail(w, err)
		return
	}
	type review struct {
		Question    publicQuestion `json:"question"`
		Answer      []int          `json:"answer"`
		Explanation string         `json:"explanation"`
		Source      content.Source `json:"source"`
	}
	byTopic := map[string]*struct {
		Topic   string `json:"topic"`
		Title   string `json:"title"`
		Seen    int    `json:"seen"`
		Correct int    `json:"correct"`
	}{}
	out := []review{}
	for _, q := range questions {
		out = append(out, review{s.public(q), q.Answer, q.Explanation, q.Source})
		if _, ok := byTopic[q.Topic]; !ok {
			byTopic[q.Topic] = &struct {
				Topic   string `json:"topic"`
				Title   string `json:"title"`
				Seen    int    `json:"seen"`
				Correct int    `json:"correct"`
			}{Topic: q.Topic, Title: s.Set.Title(q.Topic)}
		}
		byTopic[q.Topic].Seen++
	}
	breakdown := []any{}
	keys := make([]string, 0, len(byTopic))
	for k := range byTopic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		breakdown = append(breakdown, byTopic[k])
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"session": sess, "review": out, "breakdown": breakdown,
	})
}

// grade marks one answer and returns what the reader gave, for the record.
func (s *Server) grade(q content.Question, req answerRequest) (bool, string) {
	if q.Type == "text" {
		return q.CorrectText(req.Text), strings.TrimSpace(req.Text)
	}
	return q.Correct(req.Given), joinInts(req.Given)
}

// question pulls one question out of a session's stored payload. Reading it from there
// rather than from the live bank means a question generated on the fly can still be
// graded, and an edit to the content files cannot retroactively change a past answer.
func (s *Server) question(sessionID, questionID string) (content.Question, error) {
	payload, err := s.DB.Payload(sessionID)
	if err != nil {
		return content.Question{}, fmt.Errorf("that quiz has expired or never existed")
	}
	var questions []content.Question
	if err := json.Unmarshal([]byte(payload), &questions); err != nil {
		return content.Question{}, err
	}
	for _, q := range questions {
		if q.ID == questionID {
			return q, nil
		}
	}
	return content.Question{}, fmt.Errorf("that question is not part of this quiz")
}

func joinInts(in []int) string {
	parts := make([]string, 0, len(in))
	for _, i := range in {
		parts = append(parts, strconv.Itoa(i))
	}
	return strings.Join(parts, ",")
}

const idAlphabet = "abcdefghijkmnpqrstuvwxyz23456789"

func newID(rnd *rand.Rand) string {
	b := make([]byte, 12)
	for i := range b {
		b[i] = idAlphabet[rnd.Intn(len(idAlphabet))]
	}
	return string(b)
}

func intParam(r *http.Request, name string, fallback, max int) int {
	v, err := strconv.Atoi(r.URL.Query().Get(name))
	if err != nil || v <= 0 {
		return fallback
	}
	if v > max {
		return max
	}
	return v
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func (s *Server) fail(w http.ResponseWriter, err error) {
	s.Log.Error("request failed", "err", err)
	writeError(w, http.StatusInternalServerError, "Something broke on the server.")
}

// Secure adds the few headers worth setting on a small self-hosted app.
func Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "same-origin")
		next.ServeHTTP(w, r)
	})
}


// handleEpisodes lists the audiobook with the caller's progress through it.
func (s *Server) handleEpisodes(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimSpace(r.URL.Query().Get("player"))
	progress := map[string]store.Listen{}
	if player != "" {
		rows, err := s.DB.Listening(player)
		if err != nil {
			s.fail(w, err)
			return
		}
		for _, row := range rows {
			progress[row.Episode] = row
		}
	}
	type episodeView struct {
		audio.Episode
		PositionSec float64 `json:"positionSec"`
		ListenedSec float64 `json:"listenedSec"`
	}
	out := []episodeView{}
	for _, e := range s.Audio.Episodes {
		p := progress[e.Topic]
		out = append(out, episodeView{Episode: e, PositionSec: p.PositionSec, ListenedSec: p.ListenedSec})
	}
	writeJSON(w, http.StatusOK, map[string]any{"episodes": out, "minutes": s.Audio.Minutes()})
}

type listenRequest struct {
	Player   string  `json:"player"`
	Episode  string  `json:"episode"`
	Position float64 `json:"position"`
	// Delta is seconds played since the last report, not a running total.
	Delta float64 `json:"delta"`
}

func (s *Server) handleListen(w http.ResponseWriter, r *http.Request) {
	var req listenRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 2<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "That request did not parse.")
		return
	}
	// A delta larger than the reporting interval means a stall or a tab left open, not
	// listening, so it is capped rather than trusted.
	if req.Delta > 60 {
		req.Delta = 60
	}
	playerID, _, err := s.DB.Player(req.Player)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := s.DB.RecordListening(playerID, req.Episode, req.Position, req.Delta); err != nil {
		s.fail(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
