package models

import (
	"github.com/jmoiron/sqlx"

	"github.com/koinworks/asgard-heimdal/libs/serror"
)

type (
	User struct {
		ID    int64  `json:"id" db:"user_id"`
		Code  string `json:"code" db:"user_code"`
		Name  string `json:"name" db:"user_name"`
		Email string `json:"email" db:"user_email"`
	}

	InsertUser struct {
		Code  string
		Name  string
		Email string
	}
)

func GetUserList(pn Pagination) (rows []User, errx serror.SError) {
	q := `
		SELECT *
		FROM users.users
		WHERE
			is_active = 1::BIT
		LIMIT $1
		OFFSET $2
	`

	rs, err := db.Queryx(q, pn.Limit, pn.Offset)
	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to query get user list")
		return rows, errx
	}

	defer rs.Close()
	for rs.Next() {
		var row User
		err = rs.StructScan(&row)
		if err != nil {
			errx = serror.NewFromErrorc(err, "Failed to struct scanning")
			return rows, errx
		}

		rows = append(rows, row)
	}

	return rows, errx
}

func (ox *User) GetUserByID(id int64) (errx serror.SError) {
	q := `
		SELECT
			user_id,
			user_code,
			user_name,
			user_email
		FROM users.users
		WHERE
			user_id = $1 AND
			is_active = 1::BIT
		LIMIT 1
	`
	err := db.QueryRowx(q, id).StructScan(ox)
	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to query get user by id")
		return errx
	}

	return errx
}

func (ox InsertUser) Execute(tx *sqlx.Tx) (id int64, errx serror.SError) {
	q := `
		INSERT INTO users.users (user_code user_name, user_email)
		VALUES ($1, $2, $3)
	`

	var err error
	switch {
	case tx != nil:
		err = tx.QueryRow(q, ox.Code, ox.Name, ox.Email).Scan(&id)

	default:
		err = db.QueryRow(q, ox.Code, ox.Name, ox.Email).Scan(&id)
	}

	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to execute insert user")
		return id, errx
	}

	return id, errx
}
