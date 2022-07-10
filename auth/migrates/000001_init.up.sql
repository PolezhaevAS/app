

-- Create schema
CREATE SCHEMA users;

-- Create table users
CREATE TABLE users.users (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	name varchar NOT NULL,
	login varchar NOT NULL,
	"password" varchar NOT NULL,
	email varchar NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un UNIQUE (login)
);