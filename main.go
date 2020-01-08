package main

import (
	"database/sql"
	"fmt"
	"github.com/panasoniclam/trending-github/Db"
)

func init()  {
	fmt.Println("lll")
}

type (
	Person struct {
		FirstName string  `db:"first_name""`
		LastName string   `db:"last_name"`
		Email string

	}
	Place struct {
		Conutry string
		City sql.NullString
		Telcode int

	}
)
var Schema =`
	 CREATE TABLE PERSON (
      first_name text,
      last_name text,
      email text
    );
    CREATE TABLE PLACE (
     country text,
     city text null,
     telcode integer
   )
`
func main()  {
	db:=  Db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "3414pic2ck2pi",
		DbName:   "postgres",
	}
	db.Connection()
	defer  db.Close()
	//db.Db.MustExec(Schema)
    //tx:= db.Db.MustBegin()
    //tx.MustExec("INSERT INTO person(first_name,last_name,email) VALUES ($1,$2,$3)","duc","nguyen","nguyenduc@gmail.com")
    //tx.MustExec("INSERT INTO person(first_name,last_name,email) VALUES ($1,$2,$3)","phong","nguyen","nguyenphong@gmail.com")
    //tx.NamedExec("INSERT INTO person(first_name,last_name,email) VALUES (:first_name,:last_name,:email)", &Person{
	//	FirstName: "phuong",
	//	LastName:  "nguyen",
	//	Email:     "nguyenphuong@gmail.com",
	//} )
    //tx.NamedExec("INSERT INTO place(country,city,telcode) VALUES (:country,:city,:telcode)",&Place{
	//	Conutry: "china",
	//	City:    sql.NullString{},
	//	Telcode: 999,
	//})
    //tx.Commit()

    query:= db.Db
    people := []Person{}

    query.Select(&people,"SELECT  * FROM person order by first_name ASC")
    fmt.Println(people)
    lam:= Person{}
    query.Get(&lam,"SELECT * from person WHERE first_name=$1","lam")
    fmt.Println(lam)



	//e:= echo.New()
	//e.GET("/", func(context echo.Context) error {
	//	return  context.String(http.StatusOK,"wellcome to db")
	//})
	//e.Logger.Fatal(e.Start(":8000"))
}
