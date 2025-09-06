-- Database initialization script

-- Create application user if it doesn't exist
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles WHERE rolname = 'taskapi_user'
   ) THEN
      CREATE USER taskapi_user WITH PASSWORD 'password';
   END IF;
END
$do$;

-- Create application database if it doesn't exist
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_database WHERE datname = 'coco_db'
   ) THEN
      CREATE DATABASE coco_db;
   END IF;
END
$do$;

-- Grant privileges to app user
GRANT ALL PRIVILEGES ON DATABASE coco_db TO taskapi_user;

-- Switch to target database
\c coco_db;

-- Schema + default privileges
GRANT ALL ON SCHEMA public TO taskapi_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO taskapi_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO taskapi_user;

-- Confirmation
SELECT 'Database initialization completed!' AS status;
