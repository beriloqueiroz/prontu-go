-- Create "sessions" table
CREATE TABLE "public"."sessions" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "start_date" timestamptz NULL,
  "time_in_minutes" bigint NULL,
  "amount" numeric NULL,
  "end_date" timestamptz NULL,
  "notes" text NULL,
  "cids" text NULL,
  "forms" text NULL,
  "origin" text NULL,
  "external_id" text NULL,
  "location" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_sessions_deleted_at" to table: "sessions"
CREATE INDEX "idx_sessions_deleted_at" ON "public"."sessions" ("deleted_at");
-- Create "patients" table
CREATE TABLE "public"."patients" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "patient_id" text NULL,
  "session_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_sessions_patients" FOREIGN KEY ("session_id") REFERENCES "public"."sessions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_patients_deleted_at" to table: "patients"
CREATE INDEX "idx_patients_deleted_at" ON "public"."patients" ("deleted_at");
-- Create "professionals" table
CREATE TABLE "public"."professionals" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "professional_id" text NULL,
  "session_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_sessions_professionals" FOREIGN KEY ("session_id") REFERENCES "public"."sessions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_professionals_deleted_at" to table: "professionals"
CREATE INDEX "idx_professionals_deleted_at" ON "public"."professionals" ("deleted_at");
