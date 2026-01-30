-- sql
-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 1. USERS (For Admin/Auth)
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       role TEXT NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 2. COUNTIES (Dimension: The "Where")
CREATE TABLE counties (
                          id INTEGER PRIMARY KEY,
                          name TEXT NOT NULL,
                          code TEXT UNIQUE NOT NULL,
                          former_province TEXT,
                          area_sq_km NUMERIC,
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 3. INDICATORS (Dimension: The "What")
CREATE TABLE indicators (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            code TEXT UNIQUE NOT NULL,
                            name TEXT NOT NULL,
                            description TEXT,
                            unit TEXT NOT NULL,
                            source TEXT NOT NULL,
                            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 4. OBSERVATIONS (Fact: The Data)
CREATE TABLE observations (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              county_id INTEGER NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
                              indicator_id UUID NOT NULL REFERENCES indicators(id) ON DELETE CASCADE,
                              year INTEGER NOT NULL,
                              value NUMERIC NOT NULL,
                              source_document TEXT,
                              created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                              UNIQUE (county_id, indicator_id, year)
);

CREATE INDEX idx_observations_county ON observations(county_id);
CREATE INDEX idx_observations_indicator ON observations(indicator_id);
CREATE INDEX idx_observations_year ON observations(year);

-- +goose Down
DROP INDEX IF EXISTS idx_observations_year;
DROP INDEX IF EXISTS idx_observations_indicator;
DROP INDEX IF EXISTS idx_observations_county;

DROP TABLE IF EXISTS observations;
DROP TABLE IF EXISTS indicators;
DROP TABLE IF EXISTS counties;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";
