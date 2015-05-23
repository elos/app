package main

import (
	"log"

	"github.com/elos/aeolus/builtin/ego"
)

func main() {
	if err := ego.Generate("./definitions/hosts/app.json", "../"); err != nil {
		log.Print(err)
	}
}
