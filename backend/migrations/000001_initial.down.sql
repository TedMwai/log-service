BEGIN;
DROP INDEX IF EXISTS "fk_logs_microservices_microservice_id";
DROP TABLE IF EXISTS "logs";
DROP TABLE IF EXISTS "microservices";
COMMIT;