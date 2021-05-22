CREATE TABLE IF NOT EXISTS job(
    id TEXT NOT NULL UNIQUE CHECK(id <> ''),
    name TEXT NOT NULL CHECK(name <> ''),
    status TEXT NOT NULL CHECK(status <> ''),
    input TEXT NOT NULL CHECK(json_type(input) == 'object'),
    succeed_at TEXT NULL,
    errored_at TEXT NULL,
    claimed_at TEXT NULL,
    scheduled_for TEXT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(created_at <> ''),
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(updated_at <> '')
);

CREATE INDEX IF NOT EXISTS job_name ON job(name);

CREATE INDEX IF NOT EXISTS job_status ON job(status);

CREATE INDEX IF NOT EXISTS job_created_at ON job(created_at);

CREATE INDEX IF NOT EXISTS job_updated_at ON job(updated_at);

CREATE INDEX IF NOT EXISTS job_scheduled_for ON job(scheduled_for);

CREATE TRIGGER IF NOT EXISTS set_job_updated_at
AFTER
UPDATE
    ON job BEGIN
UPDATE
    job
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;