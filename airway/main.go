package main

import (
	"fmt"
	"os"

	"github.com/daqing/airway-cli/lib/generator"
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
	default:
		showHelp()
	}

}

func showHelp() {
	fmt.Println("airway g [what] [params]")
}
