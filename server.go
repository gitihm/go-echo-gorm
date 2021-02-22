package main

import (
	"main/db"
	"main/routers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db.Init()

	e := routers.Init()
	e.Logger.Fatal(e.Start(":4000"))
}
