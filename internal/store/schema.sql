-- One row per person, matched case-insensitively so "Wilhelm" and "wilhelm" are the
-- same player on the leaderboard.
CREATE TABLE IF NOT EXISTS players (
    id         INTEGER PRIMARY KEY,
    name       TEXT    NOT NULL UNIQUE COLLATE NOCASE,
    created_at INTEGER NOT NULL
);

-- A quiz run. payload holds the questions as they were served, answer keys included,
-- so grading survives a restart and a generated question can be marked long after the
-- generator forgot about it.
CREATE TABLE IF NOT EXISTS sessions (
    id          TEXT    PRIMARY KEY,
    player_id   INTEGER NOT NULL REFERENCES players(id),
    mode        TEXT    NOT NULL,
    topic       TEXT    NOT NULL DEFAULT '',
    total       INTEGER NOT NULL,
    score       INTEGER NOT NULL DEFAULT 0,
    payload     TEXT    NOT NULL,
    started_at  INTEGER NOT NULL,
    finished_at INTEGER,
    duration_ms INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS answers (
    id          INTEGER PRIMARY KEY,
    session_id  TEXT    NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    question_id TEXT    NOT NULL,
    topic       TEXT    NOT NULL,
    given       TEXT    NOT NULL,
    correct     INTEGER NOT NULL,
    confident   INTEGER NOT NULL,
    ms          INTEGER NOT NULL,
    answered_at INTEGER NOT NULL,
    UNIQUE (session_id, question_id)
);

CREATE INDEX IF NOT EXISTS answers_by_question ON answers (question_id);
CREATE INDEX IF NOT EXISTS sessions_by_player  ON sessions (player_id, finished_at);
CREATE INDEX IF NOT EXISTS sessions_finished   ON sessions (finished_at);
