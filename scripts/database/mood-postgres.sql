CREATE DATABASE mood;
\c mood

CREATE TABLE teams (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	score INTEGER NOT NULL,
	created TIMESTAMP NOT NULL
);

CREATE INDEX idx_teams_created ON teams(created);
