BEGIN;

CREATE TABLE IF NOT EXISTS "microservices" (
    "id" VARCHAR NOT NULL,
    "created_by" VARCHAR NOT NULL DEFAULT 'SYSTEM',
    "updated_by" VARCHAR NOT NULL DEFAULT 'SYSTEM',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "deleted_at" TIMESTAMPTZ,
    "metadata" JSONB NOT NULL DEFAULT '{}',
    "name" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "logs" (
    "id" VARCHAR NOT NULL,
    "created_by" VARCHAR NOT NULL DEFAULT 'SYSTEM',
    "updated_by" VARCHAR NOT NULL DEFAULT 'SYSTEM',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "deleted_at" TIMESTAMPTZ,
    "metadata" JSONB NOT NULL DEFAULT '{}',
    "microservice_id" VARCHAR NOT NULL,
    "log_level" VARCHAR NOT NULL,
    "message" VARCHAR NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("microservice_id") REFERENCES "microservices" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "fk_logs_microservices_microservice_id" ON "logs" ("microservice_id")
WHERE
    deleted_at IS NULL;

COMMIT;