package generator

import (
	"fmt"
	"strings"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/tpl"
)

type ActionGenerator struct {
	Mod     string
	Name    string
	APIName string
}

func GenAction(xargs []string) {
	if len(xargs) != 2 {
		helper.Help("airway g action [api] [action]")
	}

	GenerateAPIAction(xargs[0], xargs[1])
}

func GenerateAPIAction(mod, name string) {
	apiName := apiDirName(mod)

	targetFileName := strings.Join(
		[]string{
			".",
			"app",
			"api",
			apiName,
			name + "_action.go",
		},
		"/",
	)

	err := helper.ExecTemplate(
		tpl.Action(),
		targetFileName,
		ActionGenerator{Mod: mod, Name: helper.ToCamel(name), APIName: apiName},
	)

	if err != nil {
		fmt.Println(err)
	}

}
