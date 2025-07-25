package generator

import (
	"fmt"
	"strings"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/tpl"
)

type ServiceData struct {
	Name   string
	Fields string
	SQLH   string
}

func GenService(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: awcli g service <name> <field:type> <field:type>...")
		return
	}

	var data = ServiceData{
		Name: helper.ToCamel(args[0]),
	}

	name := args[0]

	var fields []string
	var sqlh []string
	for _, arg := range args[1:] {
		parts := strings.Split(arg, ":")
		fields = append(fields, fmt.Sprintf("%s %s", parts[0], parts[1]))
		sqlh = append(sqlh, fmt.Sprintf("\"%s\": %s", parts[0], parts[0]))
	}

	data.Fields = strings.Join(fields, ", ")
	data.SQLH = "{" + strings.Join(sqlh, ", ") + "}"

	helper.ExecTemplate(
		tpl.Service(),
		strings.Join([]string{
			".",
			"app",
			"services",
			name + ".go",
		}, "/"),
		data,
	)
}
