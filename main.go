package main

import (
	"fmt"
	"os"

	"netselect/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
