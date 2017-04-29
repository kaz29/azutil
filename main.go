package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
	commands "github.com/kaz29/azutil/commands"
)

func main() {

	parser := flags.NewParser(&commands.Azutil, flags.Default)
	parser.Name = "azutil"

	parser.Parse()

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

}
