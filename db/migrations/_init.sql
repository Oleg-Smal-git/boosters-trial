-- Instantiate an empty database.
CREATE DATABASE "boosters-trial";
\connect "boosters-trial";
CREATE SCHEMA IF NOT EXISTS "public";

-- Create a new full-privilege user.
CREATE USER "boosters-trial-user" WITH PASSWORD 'secret123';
GRANT ALL ON DATABASE "boosters-trial" TO "boosters-trial-user";
GRANT USAGE ON SCHEMA "public" TO "boosters-trial-user";
GRANT ALL ON ALL TABLES IN SCHEMA "public" TO "boosters-trial-user";
ALTER DEFAULT PRIVILEGES IN SCHEMA "public" GRANT ALL ON TABLES TO "boosters-trial-user";

-- Create a new readonly user.
CREATE USER "boosters-trial-readonly-user" WITH PASSWORD 'secret123';
GRANT CONNECT ON DATABASE "boosters-trial" TO "boosters-trial-readonly-user";
GRANT USAGE ON SCHEMA "public" TO "boosters-trial-readonly-user";
GRANT SELECT ON ALL TABLES IN SCHEMA "public" TO "boosters-trial-readonly-user";
ALTER DEFAULT PRIVILEGES IN SCHEMA "public" GRANT SELECT ON TABLES TO "boosters-trial-readonly-user";

-- Create version table that would manage migrations.
CREATE TABLE versions (
    id VARCHAR(255) PRIMARY KEY
);
