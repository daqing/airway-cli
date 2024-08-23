package generator

import (
	"fmt"

	"github.com/daqing/airway-cli/helper"
)

func Generate(args []string) {
	if len(args) == 0 {
		helper.Help("airway g [what] [params]")
	}

	thing := args[0]
	xargs := args[1:]

	switch thing {
	case "action":
		GenAction(xargs)
	case "api":
		GenAPI(xargs)
	case "model":
		GenModel(xargs)
	default:
		panic("unknown generator")
	}

	fmt.Println("done.")
}
