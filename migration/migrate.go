package migration

import (
	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/gomigrate/migrate_up"
)

func Migrate(args []string) {
	dir := "./db/migrate"
	var version string = ""

	if len(args) > 0 {
		version = args[0]
	}

	dsn := helper.GetDSN()

	if version == "" {
		migrate_up.All(dir, dsn)
	} else {
		migrate_up.Version(dir, version, dsn)
	}
}
