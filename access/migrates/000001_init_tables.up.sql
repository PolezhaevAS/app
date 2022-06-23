
-- Init tables


-- Create new schema
CREATE SCHEMA "access"

-- Create table groups. This table saved groups by creating users.
CREATE TABLE IF NOT EXISTS "access"."groups" (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" varchar NOT NULL,
	description varchar NULL,
	CONSTRAINT groups_pk PRIMARY KEY (id)
);
CREATE INDEX groups_id_idx ON access.groups USING btree (id);

-- Create services table. This table need to save service name. Service name using in JWT-token claims.
CREATE TABLE "access".services (
	id int4 NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT services_pk PRIMARY KEY (id)
);
CREATE INDEX services_id_idx ON access.services USING btree (id);

-- Create methods table. This table saved service methods with name. Methods name using in JWT-token claims.
CREATE TABLE "access".methods (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	service_id int4 NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT methods_pk PRIMARY KEY (id)
);
CREATE INDEX methods_id_idx ON access.methods USING btree (id);

ALTER TABLE "access".methods ADD CONSTRAINT methods_fk FOREIGN KEY (service_id) REFERENCES "access".services(id);

-- Create group methods table. This table need to add method into group.
CREATE TABLE "access".group_methods (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	group_id int4 NOT NULL,
	method_id int4 NOT NULL,
	CONSTRAINT group_methods_pk PRIMARY KEY (id)
);
CREATE INDEX group_methods_group_id_idx ON access.group_methods USING btree (group_id);

ALTER TABLE "access".group_methods ADD CONSTRAINT group_methods_fk FOREIGN KEY (group_id) REFERENCES "access"."groups"(id);
ALTER TABLE "access".group_methods ADD CONSTRAINT group_methods_fk_1 FOREIGN KEY (method_id) REFERENCES "access".methods(id);

-- Create user groups table. This table need to add user in new group.
CREATE TABLE "access".user_groups (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	user_id int4 NOT NULL,
	group_id int4 NOT NULL,
	CONSTRAINT user_groups_pk PRIMARY KEY (id)
);
CREATE INDEX user_groups_group_id_idx ON access.user_groups USING btree (group_id);
CREATE INDEX user_groups_user_id_idx ON access.user_groups USING btree (user_id);


ALTER TABLE "access".user_groups ADD CONSTRAINT user_groups_fk FOREIGN KEY (group_id) REFERENCES "access"."groups"(id);
