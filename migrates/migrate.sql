/* SQLEditor (Postgres)*/

/**
 DROP TABLE IF EXISTS
*/
DROP TABLE IF EXISTS public.events;
DROP TABLE IF EXISTS public.tracked_events;

/**
 CREATE TABLE
*/

CREATE TABLE public.events
(
id UUID UNIQUE ,
name VARCHAR(255) UNIQUE ,
properties JSONB,
"created_on" TIMESTAMPTZ,
"updated_on" TIMESTAMPTZ,
"deleted_on" TIMESTAMPTZ,
CONSTRAINT "events_pkey" PRIMARY KEY (id)
);

CREATE TABLE public.tracked_events
(
id SERIAL NOT NULL UNIQUE ,
name VARCHAR(255) NOT NULL,
properties JSONB,
"created_on" TIMESTAMPTZ,
"tracked_on" TIMESTAMPTZ NOT NULL,
CONSTRAINT "tracked_events_pkey" PRIMARY KEY (id)
);


/**
 CREATE INDEX
 ADD FOREIGN KEY
*/

CREATE INDEX "events_id_idx" ON public.events(id);
CREATE INDEX "events_name_idx" ON public.events(name);
CREATE INDEX "tracked_events_id_idx" ON "tracked_events"(id);
CREATE INDEX "tracked_events_name_idx" ON "tracked_events"(name);