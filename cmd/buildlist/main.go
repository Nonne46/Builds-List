package main

import (
	"log"

	"github.com/Nonne46/Builds-List/internal/app/buildlist"
)

func main() {
	if err := buildlist.Start(); err != nil {
		log.Fatal(err)
	}
}
