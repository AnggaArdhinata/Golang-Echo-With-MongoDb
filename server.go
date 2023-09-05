package main

import (
	"github.com/AnggaArdhinata/drillingfazz005/src/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	e := routes.Route()

	e.Logger.Fatal(e.Start(":8080"))
}
