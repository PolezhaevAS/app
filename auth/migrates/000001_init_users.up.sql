
-- Create new schema
CREATE SCHEMA "users"
;

-- Create users table
CREATE TABLE "users".users (
    id int NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" varchar NOT NULL,
	login varchar NOT NULL,
	"password" varchar NOT NULL
);