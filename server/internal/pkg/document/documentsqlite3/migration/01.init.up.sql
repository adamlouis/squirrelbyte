CREATE TABLE IF NOT EXISTS document(
    id TEXT NOT NULL UNIQUE CHECK(id <> ''),
    header TEXT NOT NULL CHECK(json_type(header) == 'object'),
    body TEXT NOT NULL CHECK(json_type(body) == 'object'),
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(created_at <> ''),
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(updated_at <> '')
);

CREATE INDEX IF NOT EXISTS document_created_at ON document(created_at);

CREATE INDEX IF NOT EXISTS document_updated_at ON document(updated_at);

CREATE TRIGGER IF NOT EXISTS set_document_updated_at
AFTER
UPDATE
    ON document BEGIN
UPDATE
    document
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;

CREATE TABLE IF NOT EXISTS path(name TEXT NOT NULL UNIQUE CHECK(name <> ''));