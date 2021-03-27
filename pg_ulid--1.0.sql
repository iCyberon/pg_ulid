-- use CREATE EXTENSION
\echo Use "CREATE EXTENSION pg_ulid" to load this file. \quit

CREATE OR REPLACE FUNCTION ulid() RETURNS TEXT
  AS 'pg_ulid', 'Ulid' LANGUAGE C;

CREATE OR REPLACE FUNCTION ulid_to_time(text) RETURNS timestamp
  AS 'pg_ulid', 'UlidToTime' LANGUAGE C;

CREATE OR REPLACE FUNCTION ulid_to_local_time(text) RETURNS timestamptz
  AS 'pg_ulid', 'UlidToLocalTime' LANGUAGE C;