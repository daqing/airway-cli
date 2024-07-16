package main

import (
	"fmt"
	"os"

	"github.com/daqing/airway-cli/generator"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		showHelp()
		return
	}

	switch args[0] {
	case "g":
		generator.Generate(args[1:])
	case "-v", "version":
		showVersion()
	default:
		showHelp()
	}

}

func showHelp() {
	fmt.Println("airway [-v | g] [what] [params]")
}

func showVersion() {
	fmt.Printf("%s\n", version())
}
