package queries

const (
	SIGN_IN = "SELECT id, name, login, password FROM users.users WHERE login = $1 and password = $2"
	USER    = "SELECT id, name, login, password FROM users.users WHERE id = $1"
	LIST    = "SELECT id, name, login, password FROM users.users"
	CREATE  = "INSERT INTO users.users(name, login, password) VALUES($1, $2, $3)"
	UPDATE  = "UPDATE users.users set name = $2, login = $3, password = $4 WHERE id = $1"
	DELETE  = "DELETE ROM users.users WHERE id = $1"
)
