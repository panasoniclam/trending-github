package Db

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

func (s *Sql)Connection()     {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",s.Host,s.Port,s.UserName,s.Password,s.DbName)
	s.Db =   sqlx.MustConnect("postgres",dataSource)
	if err:= s.Db.Ping() ;err!= nil {
		log.Error(err.Error())

	}
	
	fmt.Println("connect db success!")

}

func (s *Sql)Close()  {
	s.Db.Close()
}