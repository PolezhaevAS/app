package queries

const (
	GROUP = "SELECT id, name, description FROM access.groups WHERE id = $1;"

	GROUP_LIST = "SELECT id, name, description FROM access.groups;"

	GROUP_CREATE = `
	INSERT INTO "access"."groups"
	("name", description)
	VALUES($1, $2)
	RETURNING id;
	`

	GROUP_UPDATE = `
	UPDATE "access"."groups"
	SET "name"=$2, description=$3
	WHERE id=$1;
	`

	GROUP_DELETE = "DELETE FROM access.groups WHERE id=$1;"

	GROUP_ADD_METHOD = `
	INSERT INTO "access".group_methods
	(group_id, method_id)
	VALUES($1, $2);
	`

	GROUP_REMOVE_METHOD = "DELETE FROM access.group_methods where group_id = $1 and method_id = $2;"
)
