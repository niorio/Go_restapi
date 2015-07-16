package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
)

func TodoIndex(db *sql.DB, r render.Render) {
	rows, err := db.Query("SELECT id, name, completed FROM todos")
	PanicIf(err)
	defer rows.Close()
	
	var id int
	var name string
	var completed bool
	var todos Todos
	
	for rows.Next() {
		err:= rows.Scan(&id, &name, &completed)
		PanicIf(err)
		
		var todo Todo
		todo.Id = id
		todo.Name = name
		todo.Completed = completed
		todos = append(todos, todo)
	}
	
	r.JSON(200, todos)
}

func TodoCreate(db *sql.DB, p martini.Params, r render.Render, req *http.Request, todo Todo) {
	
	var id int
	err := db.QueryRow("INSERT INTO todos (name, completed) VALUES ($1, $2) RETURNING id", todo.Name, todo.Completed).Scan(&id)
	PanicIf(err)
	todo.Id = int(id)
	r.JSON(200, todo)

}

func TodoShow(db *sql.DB, p martini.Params, r render.Render) {
	row := db.QueryRow("SELECT id, name, completed FROM todos WHERE id = $1", p["id"])
	
	var id int
	var name string
	var completed bool
	err := row.Scan(&id, &name, &completed)
	PanicIf(err)
	
	todo := Todo{Id: id, Name: name, Completed: completed}
	
	r.JSON(200, todo)
}

func TodoUpdate(db *sql.DB, p martini.Params, r render.Render, req *http.Request, todo Todo) {
	var id int
	err := db.QueryRow("UPDATE todos SET name=$1, completed=$2 WHERE id=$3 RETURNING id", todo.Name, todo.Completed, p["id"]).Scan(&id)
	PanicIf(err)
	todo.Id = id
	r.JSON(200,todo)
}

func TodoDestroy(db *sql.DB, p martini.Params, r render.Render) {
	_, err := db.Exec("DELETE FROM todos WHERE id=$1", p["id"])
	PanicIf(err)
	r.JSON(200, map[string]interface{}{"success": true})
}