package main

import (
	"ip-finder/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.GerarCLI()

	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
