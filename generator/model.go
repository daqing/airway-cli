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
			Fields:    buildFields(fields),
		},
	)
}

func buildFields(fields []string) string {
	if len(fields) == 0 {
		return ""
	}

	// fields are in the form of "email:string age:int"
	var result []string
	for _, f := range fields {
		parts := strings.Split(f, ":")
		result = append(result, helper.ToCamel(parts[0])+" "+parts[1]+" `json:\""+parts[0]+"\"`")
	}

	return strings.Join(result, "\n")
}
