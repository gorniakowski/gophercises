package main

import (
	"gophercises/ex7/task/cmd"
	"gophercises/ex7/task/db"
)

func main() {
	err := db.Initialize("task.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cmd.Execute()
}
