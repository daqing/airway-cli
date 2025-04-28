package migration

import (
	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/gomigrate/status"
)

func Status(args []string) {
	dir := "./db/migrate"

	status.Show(dir, helper.GetDSN())
}
