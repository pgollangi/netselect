package main

import (
	"log"

	"github.com/pgollangi/netselect/commands"

	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(commands.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
