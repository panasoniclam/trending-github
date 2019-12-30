package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type (
	Sql struct {
		Db *sqlx.DB
		Host string
		Port int
		Username string
		Password string
		Dbname string

	}

)

func (s *Sql)Connection()   {
	datasource := fmt.Sprint("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable")
	s.Db  =  sqlx.MustConnect("postgres",datasource)

	if err := s.Db.Ping(); err != nil{

		return
	}
}