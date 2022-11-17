-- Instantiate an empty database.
CREATE DATABASE "boosters-trial";
\connect "boosters-trial";
CREATE SCHEMA IF NOT EXISTS "public";

-- Create a new full-privilege and a new readonly users.
CREATE USER "boosters-trial-user" WITH PASSWORD 'secret123';
GRANT ALL ON DATABASE "boosters-trial" TO "boosters-trial-user";
CREATE USER "boosters-trial-readonly-user" WITH PASSWORD 'secret123';
GRANT SELECT ON ALL TABLES IN SCHEMA "public" TO "boosters-trial-readonly-user";

-- Create version table that would manage migrations.
CREATE TABLE versions (
    id VARCHAR(255) PRIMARY KEY
);
