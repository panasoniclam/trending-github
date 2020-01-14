package main

import (
	"fmt"
	"github.com/panasoniclam/trending-github/db"

)
const schema = `
   CREATE TABLE TEST(
     name text,
     email text
    )
`
type person struct {
	Name string `db:"name"`
	Email string `db:"email"`
}
func main()  {
   db := 	db.Sql{

		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "3414pic2ck2pi",
		DbName:   "postgres",
	}
	db.Connection()

    //db.Db.MustExec(schema)
    //tx:= db.Db.MustBegin()
    //tx.MustExec("INSERT INTO test(name,email) VALUES ($1,$2)","nguyen ngoc lam","lamnn8@fpt.com")
    //tx.NamedExec("INSERT INTO test(name,email) VALUES (:name,:email)",&person{
	//	"dung",
	//	"dungtd",
	//})
    //tx.Commit()

    peope:=[]person{}
    query := db.Db
    query.Select(&peope,"SELECT * FROM test")
    fmt.Println(peope)
    defer db.Close()

}