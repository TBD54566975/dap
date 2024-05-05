CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE daps (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id CHAR(26) NOT NULL UNIQUE,
    did VARCHAR(255) NOT NULL,
    handle VARCHAR(64) NOT NULL UNIQUE,
    proof TEXT,
    date_created TEXT NOT NULL CHECK(date_created LIKE '____-__-__T__:__:__Z')
);
CREATE TABLE challenges (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id CHAR(26) NOT NULL UNIQUE,
    challenge TEXT NOT NULL UNIQUE,
    date_created TEXT NOT NULL CHECK(date_created LIKE '____-__-__T__:__:__Z')
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240408133638');
