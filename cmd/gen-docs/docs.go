package main

import (
	"log"

	"netselect/commands"

	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(commands.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
