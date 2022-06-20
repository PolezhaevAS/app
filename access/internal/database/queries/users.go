package queries

const (
	USERS_LIST = `
	SELECT user_id
	FROM "access".user_groups WHERE group_id = $1;
	`

	USERS_ADD = `
	INSERT INTO "access".user_groups
	(user_id, group_id)
	VALUES($1, $2);
	`

	USERS_REMOVE = `
	DELETE FROM "access".user_groups
	WHERE group_id = $1 and user_id = $2;
	`

	USERS_ACCESS = `
	select s."name" , m."name"  from "access".user_groups ug
	join "access".group_methods gm on ug.group_id = gm.group_id 
	join "access".methods m on gm.method_id = m.id 
	join "access".services s on m.service_id = s.id 
	where user_id  = $1
	`
)
