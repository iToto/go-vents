/* SQLEditor (Postgres)*/

/**
 DROP TABLE IF EXISTS
*/
DROP TABLE IF EXISTS events;

/**
 CREATE TABLE
*/

CREATE TABLE events
(
id UUID UNIQUE ,
name VARCHAR(255) UNIQUE ,
properties JSONB,
"created_on" TIMESTAMPTZ,
"updated_on" TIMESTAMPTZ,
"deleted_on" TIMESTAMPTZ,
CONSTRAINT "events_pkey" PRIMARY KEY (id)
);

/**
 CREATE INDEX
 ADD FOREIGN KEY
*/

CREATE INDEX "events_id_idx" ON events(id);
CREATE INDEX "events_name_idx" ON events(name);