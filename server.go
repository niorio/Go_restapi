package main

import(
	"github.com/go-martini/martini"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
)

func SetupDB() *sql.DB{
	db, err := sql.Open("postgres", "dbname=todos sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error){
	if err != nil{
		panic(err)
	}
}

func main() {
	m:= martini.Classic()
	m.Map(SetupDB())
	m.Use(render.Renderer())
	
	m.Get("/", func() string{
		return "Hello!"
	})
	m.Get("/todos", TodoIndex)
	// m.Post("/todos", TodoCreate)
	// m.Get("/todos/:id", TodoShow)
	// m.Patch("/todos/:id", TodoUpdate)
	// m.Delete("/todos/:id", TodoDestroy)
	
	m.Run()
}