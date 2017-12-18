-- use CREATE EXTENSION
\echo Use "CREATE EXTENSION pg_ulid" to load this file. \quit

CREATE OR REPLACE FUNCTION ulid() RETURNS TEXT
  AS 'pg_ulid', 'Ulid' LANGUAGE C;