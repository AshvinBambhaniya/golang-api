-- +migrate Up
CREATE TABLE IF NOT EXISTS titles (
    id VARCHAR(10) PRIMARY KEY,
    title VARCHAR(255),
    type VARCHAR(20),
    description TEXT,
    release_year INTEGER,
    age_certification VARCHAR(10),
    runtime INTEGER,
    genres TEXT[],
    production_countries TEXT[],
    seasons INTEGER,
    imdb_id VARCHAR(20),
    imdb_score FLOAT,
    imdb_votes FLOAT,
    tmdb_popularity FLOAT,
    tmdb_score FLOAT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS credits (
    person_id INTEGER,
    id VARCHAR(10) REFERENCES titles(id),
    name VARCHAR(255),
    character VARCHAR(255),
    role VARCHAR(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COPY titles FROM '/path/to/titles.csv' WITH CSV HEADER;
COPY credits FROM '/path/to/credits.csv' WITH CSV HEADER;
