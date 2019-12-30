package driver


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	SQL *sql.DB
}

var Postgres = &PostgresDb{}

func Connect(host,port,user,password,dbname string) (*PostgresDb)  {
	connStr:= fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db,err :=sql.Open("postgres",connStr)
	if err !=nil {
		panic(err)
	}
	Postgres.SQL = db
	return Postgres
}