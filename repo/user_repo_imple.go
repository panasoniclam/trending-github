package repo

import (
	"database/sql"
)


type UserRepoImpl struct {
	Db *sql.DB
}

