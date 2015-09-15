package main

import(
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  _ "fmt"
)

type User struct {
  Id int64
  Description string
  Name string
  Age int64
  Bio string
  Email string
}

func main() {
  db, _ := gorm.Open("postgres", "user=postgres dbname=phoenix_crud_dev sslmode=disable")

  m := martini.Classic()
  m.Use(render.Renderer())
  m.Map(&db)

  m.Group("/books", func(r martini.Router) {
    r.Get("", ListBooks)
    r.Get("/:id", GetBook)
    r.Post("/new", NewBook)
    r.Put("/update/:id", UpdateBook)
    r.Delete("/delete/:id", DeleteBook)
  })

  m.Run()
}

func ListBooks(r render.Render, db *gorm.DB) {
  users := []User{}
  db.Find(&users)

  r.HTML(200, "list", users)
}

func GetBook() string {
  return "Get"
}

func NewBook() string {
  return "Post"
}

func UpdateBook() string {
  return "Put"
}

func DeleteBook() string {
  return "Delete"
}
