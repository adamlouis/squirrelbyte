CREATE TABLE IF NOT EXISTS kv(
    key TEXT UNIQUE NOT NULL CHECK(key <> ''),
    value TEXT NOT NULL CHECK(value <> ''),
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(created_at <> ''),
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(updated_at <> '')
);

CREATE INDEX IF NOT EXISTS kv_created_at ON kv(created_at);

CREATE INDEX IF NOT EXISTS kv_updated_at ON kv(updated_at);

CREATE TRIGGER IF NOT EXISTS set_kv_updated_at
AFTER
UPDATE
    ON kv BEGIN
UPDATE
    kv
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    key = OLD.key;

END;