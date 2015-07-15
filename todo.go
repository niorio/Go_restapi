package main

type Todo struct{
	Id 			int			`json:"id"`
	Name		string		`json:"name" binding: "required"`
	Completed	bool		`json:"completed"`
}

type Todos []Todo