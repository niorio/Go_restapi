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
	
	_, err := db.Exec("INSERT INTO todos VALUES ($1, $2, $3)", todo.Id, todo.Name, todo.Completed)
	PanicIf(err)
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
	_, err := db.Exec("UPDATE todos SET name=$1, completed=$2 WHERE id=$3", todo.Name, todo.Completed, p["id"])
	PanicIf(err)
	r.JSON(200,todo)
}
//
// func TodoDestroy() {
//
// }