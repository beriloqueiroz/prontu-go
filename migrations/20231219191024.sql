-- Create "sessions" table
CREATE TABLE "public"."sessions" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "patients" character varying(255) NULL,
  "professionals" character varying(255) NULL,
  "start_date" timestamptz NULL,
  "time_in_minutes" bigint NULL,
  "amount" numeric NULL,
  "end_date" timestamptz NULL,
  "notes" text NULL,
  "forms" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_sessions_deleted_at" to table: "sessions"
CREATE INDEX "idx_sessions_deleted_at" ON "public"."sessions" ("deleted_at");
-- Create "cids" table
CREATE TABLE "public"."cids" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "patient_id" text NULL,
  "code" text NULL,
  "name" text NULL,
  "observation" text NULL,
  "session_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_sessions_cids" FOREIGN KEY ("session_id") REFERENCES "public"."sessions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_cids_deleted_at" to table: "cids"
CREATE INDEX "idx_cids_deleted_at" ON "public"."cids" ("deleted_at");
