package main

import (
	"fmt"

	"github.com/sautmanurung1/BRI-api/infrastructure/http/server"
)

func main() {
	app := server.Server()
	err := app.Start(":5000")

	if err != nil {
		fmt.Println("This is error : ", err)
	}
}
