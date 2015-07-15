package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	// "github.com/martini-contrib/binding"
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

// func TodoCreate() {
//
// }
//
// func TodoShow() {
//
// }
//
// func TodoUpdate() {
//
// }
//
// func TodoDestroy() {
//
// }