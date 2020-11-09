package queries

const (
	GetUserList = `
		SELECT
			*
		FROM users
		WHERE
			id = $1
	`
)
