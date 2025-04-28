package generator

import (
	"fmt"

	gomigrate "github.com/daqing/gomigrate/generator"
)

func GenMigration(xargs []string) {
	dir := "./db/migrate"
	if len(xargs) != 1 {
		fmt.Println("Usage: airway g migration [name]")
	}

	name := xargs[0]

	gomigrate.Generate(name, dir)
}
