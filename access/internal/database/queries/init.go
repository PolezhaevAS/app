package queries

const (
	SERVICE_EXISTS = `
	SELECT id FROM "access".services WHERE id = $1 
	`

	SERVICE_ADD = `
	INSERT INTO "access".services
	(id, "name")
	VALUES($1, $2);
	`

	METHOD_EXISTS = `
	SELECT id FROM "access".methods WHERE service_id = $1 and "name" = $2
	`

	METHOD_ADD = `
	INSERT INTO "access".methods
	(service_id, "name")
	VALUES($1, $2);
	`
)
