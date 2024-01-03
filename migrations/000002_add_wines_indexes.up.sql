CREATE INDEX IF NOT EXISTS wines_name_idx ON wines USING GIN (to_tsvector('simple', name));
