package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type (
	Sql struct {
		Db *sqlx.DB
		Host string
		Port int
		UserName string
		Password string
		DbName string
	}
)

func (sql *Sql)Connection()  {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",sql.Host,sql.Port,sql.UserName,sql.Password,sql.DbName)
	sql.Db =  sqlx.MustConnect("postgres",dataSource)
	if err:= sql.Db.Ping(); err!= nil {
		log.Error(err.Error())
	}
	fmt.Println("connect db success !")
}

func (sql *Sql)Close()  {
	sql.Db.Close()
}