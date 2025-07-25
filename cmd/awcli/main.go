package main

import (
	"fmt"
	"os"

	"github.com/daqing/airway-cli/generator"
	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/migration"
	"github.com/daqing/airway-cli/plugin"
	"github.com/daqing/gomigrate/lib"
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
	case "migrate":
		lib.CheckOrCreateTable(helper.GetDSN())
		migration.Migrate(args[1:])
	case "rollback":
		lib.CheckOrCreateTable(helper.GetDSN())
		migration.Rollback(args[1:])
	case "migrate:status":
		lib.CheckOrCreateTable(helper.GetDSN())
		migration.Status(args[1:])
	case "plugin:install":
		plugin.Install(args[1:])
	case "-v", "version":
		showVersion()
	default:
		showHelp()
	}

}

func showHelp() {
	fmt.Println("awcli -v")
	fmt.Println("awcli g [what] [params]")
	fmt.Println("awcli migrate|rollback|migrate:status")
	fmt.Println("awcli plugin:install [/path/to/project]")
}

func showVersion() {
	fmt.Printf("%s\n", version())
}
