package migration

import (
	"fmt"
	"os"
	"strconv"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/gomigrate/rollback_to"
)

func Rollback(args []string) {
	dir := "./db/migrate"
	var step int = 1
	var err error

	if len(args) > 0 {
		step, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Err: invalid step")
			fmt.Printf("Usage: airway rollback [step]\n")
			os.Exit(1)
		}
	}

	dsn := helper.GetDSN()

	if step <= 1 {
		rollback_to.Latest(dir, dsn)
	} else {
		rollback_to.Step(dir, step, dsn)
	}
}
