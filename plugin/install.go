package plugin

import (
	"fmt"
	"time"

	"github.com/daqing/airway-cli/helper"
)

func Install(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: awcli plugin:install [/path/to/project]")
		return
	}

	projectPath := args[0]

	fmt.Println("Install current plugin to", projectPath)

	// copy ./app to [projectPath]/app
	err := helper.System(fmt.Sprintf("cp -R ./app/* %s/app/", projectPath), false)
	if err != nil {
		fmt.Println("Copy ./app to project failed:", err)
		return
	}

	// copy ./cmd to [projectPath]/cmd
	err = helper.System(fmt.Sprintf("cp -R ./cmd/* %s/cmd/", projectPath), false)
	if err != nil {
		fmt.Println("Copy ./cmd to project failed:", err)
		return
	}

	// copy ./db/migrate to [projectPath]/db/migrate
	// we need to prepend a timestamp to the migration file name
	timestamp := fmt.Sprintf("%s", time.Now().Format("20060102150405"))

	files, err := helper.Glob("./db/migrate", "*.sql")
	if err != nil {
		fmt.Println("Glob migration files failed:", err)
		return
	}

	for _, file := range files {
		basename := helper.Basename(file)
		newBaseName := fmt.Sprintf("%s_%s", timestamp, basename)

		cmd := fmt.Sprintf("cp ./db/migrate/%s %s/db/migrate/%s", basename, projectPath, newBaseName)
		err := helper.System(cmd, false)
		if err != nil {
			fmt.Println("Copy migration file failed:", err)
			return
		}
	}
}
