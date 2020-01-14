 package main

 import "fmt"

type Role int
 const (
  ADMIN  Role = iota
  MEMBER
  AGENT
 )

 func (r Role) String() string {
  return []string{"ADMIN","MEMBER","AGANT"}[r]
 }
 func main()  {
   fmt.Println(ADMIN.String())
 }