package models

type (
	InsertUserArguments struct {
		Name  string
		Email string
	}

	UpdateUserArguments struct {
		ID    int64
		Name  string
		Email string

		Updates []string
	}
)
