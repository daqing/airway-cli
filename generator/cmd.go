package generator

import (
	"fmt"
	"strings"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/tpl"
)

type CmdData struct {
	Name      string
	LowerName string
	Fields    string
	Args      string
	Args1     string
}

func GenCmd(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: awcli g cmd <name> <field> <field>...")
		return
	}

	var data = CmdData{
		Name:      helper.ToCamel(args[0]),
		LowerName: args[0],
	}

	var fields []string
	var ns []string
	var ns1 []string

	var n int
	for _, arg := range args[1:] {
		fields = append(fields, fmt.Sprintf("<%s>", arg))
		ns = append(ns, fmt.Sprintf("args[%d]", n))
		n++
		ns1 = append(ns1, fmt.Sprintf("args[%d]", n))
	}

	data.Fields = strings.Join(fields, " ")
	data.Args = strings.Join(ns, ", ")
	data.Args1 = strings.Join(ns1, ", ")

	name := args[0]

	helper.ExecTemplate(
		tpl.Cmd(),
		strings.Join([]string{
			".",
			"cmd",
			name + ".go",
		}, "/"),
		data,
	)
}
