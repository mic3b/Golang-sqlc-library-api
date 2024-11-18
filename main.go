package main

import (
	"example.com/m/internal/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app.Run()
}
