package queries

const (
	SERVICES_LIST = "SELECT id, name FROM access.services ;"

	METHODS_LIST = "SELECT id, name FROM access.methods where service_id = $1 ;"
)
