package generator

import (
	"strings"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/tpl"
)

func GenModel(xargs []string) {
	if len(xargs) == 0 {
		helper.Help("airway g model [name] [field:type]...")
	}

	GenerateModel(xargs[0], xargs[1:])
}

type Data struct {
	Name      string
	TableName string
	Fields    string
}

func GenerateModel(name string, fields []string) {
	name = helper.ToLower(name)

	targetPath := strings.Join([]string{
		".",
		"app",
		"models",
		name + ".go",
	}, "/")

	helper.ExecTemplate(
		tpl.Model(),
		targetPath,
		Data{
			Name:      helper.ToCamel(name),
			TableName: name + "s",
		},
	)
}
