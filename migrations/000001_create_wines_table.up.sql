CREATE TABLE IF NOT EXISTS wines (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    version integer NOT NULL DEFAULT 1
);