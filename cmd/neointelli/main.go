// cmd/neointelli/main.go
package main

import (
	"flag"
	"log"
	"os"

	"neointelli/internal/neointelli"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := neointelli.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
