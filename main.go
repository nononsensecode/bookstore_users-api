package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nononsensecode/bookstore_users-api/app"
)

func main() {
	app.StartApplication()
}
