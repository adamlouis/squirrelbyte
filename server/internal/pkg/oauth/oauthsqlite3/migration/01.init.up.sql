-- create config table
CREATE TABLE IF NOT EXISTS config(
    name TEXT NOT NULL UNIQUE CHECK(name <> ''),
    client_id TEXT NOT NULL,
    client_secret TEXT NOT NULL,
    auth_url TEXT NOT NULL,
    auth_url_params TEXT NOT NULL,
    token_url TEXT NOT NULL,
    redirect_url TEXT NOT NULL,
    scopes TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(created_at <> ''),
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(updated_at <> '')
);

-- create indexes
CREATE INDEX IF NOT EXISTS config_created_at ON config(created_at);

CREATE INDEX IF NOT EXISTS config_updated_at ON config(updated_at);

-- create updated_at trigger
CREATE TRIGGER IF NOT EXISTS set_config_updated_at
AFTER
UPDATE
    ON config BEGIN
UPDATE
    config
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    name = OLD.name;

END;