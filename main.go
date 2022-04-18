package main

import (
	"log"

	"github.com/fprotimaru/go-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
