-- create scheduler table
CREATE TABLE IF NOT EXISTS scheduler(
    id TEXT NOT NULL UNIQUE CHECK(id <> ''),
    schedule TEXT NOT NULL CHECK(schedule <> ''),
    job_name TEXT NOT NULL CHECK(job_name <> ''),
    input TEXT NOT NULL CHECK(json_type(input) == 'object'),
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(created_at <> ''),
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(updated_at <> '')
);

-- create indexes
CREATE INDEX IF NOT EXISTS scheduler_created_at ON scheduler(created_at);

CREATE INDEX IF NOT EXISTS scheduler_updated_at ON scheduler(updated_at);

-- create updated_at trigger
CREATE TRIGGER IF NOT EXISTS set_scheduler_updated_at
AFTER
UPDATE
    ON scheduler BEGIN
UPDATE
    scheduler
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;