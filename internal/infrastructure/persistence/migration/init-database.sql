-- public.sessions definition
-- Drop table
-- DROP TABLE public.sessions;
CREATE TABLE public.sessions (
  id text NOT NULL,
  created_at timestamptz NULL,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL,
  patients varchar(255) NULL,
  professionals varchar(255) NULL,
  start_date timestamptz NULL,
  time_in_minutes int8 NULL,
  amount numeric NULL,
  end_date timestamptz NULL,
  notes text NULL,
  forms text NULL,
  CONSTRAINT sessions_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_sessions_deleted_at ON public.sessions USING btree (deleted_at);

-- public.cids definition
-- Drop table
-- DROP TABLE public.cids;
CREATE TABLE public.cids (
  id bigserial NOT NULL,
  created_at timestamptz NULL,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL,
  patient_id text NULL,
  code text NULL,
  name text NULL,
  observation text NULL,
  session_id text NULL,
  CONSTRAINT cids_pkey PRIMARY KEY (id),
  CONSTRAINT fk_sessions_cids FOREIGN KEY (session_id) REFERENCES public.sessions(id)
);

CREATE INDEX idx_cids_deleted_at ON public.cids USING btree (deleted_at);