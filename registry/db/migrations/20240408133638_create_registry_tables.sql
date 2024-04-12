-- migrate:up
CREATE TABLE daps (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id CHAR(26) NOT NULL UNIQUE,
    did VARCHAR(255) NOT NULL,
    handle VARCHAR(64) NOT NULL UNIQUE,
    proof TEXT NOT NULL,
    date_created TEXT NOT NULL CHECK(date_created LIKE '____-__-__T__:__:__Z')
);

CREATE TABLE challenges (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id CHAR(26) NOT NULL UNIQUE,
    challenge TEXT NOT NULL UNIQUE,
    date_created TEXT NOT NULL CHECK(date_created LIKE '____-__-__T__:__:__Z')
);

-- migrate:down
DROP TABLE daps;
DROP TABLE challenges;