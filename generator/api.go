package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/daqing/airway-cli/helper"
	"github.com/daqing/airway-cli/tpl"
)

func GenAPI(xargs []string) {
	if len(xargs) != 1 {
		helper.Help("airway g api [name]")
	}

	GenerateAPI(xargs[0])
}

type APIGenerator struct {
	Mod       string
	ModelName string
	APIName   string
}

func GenerateAPI(name string) {
	dir := apiDirName(name)
	dirPath := fmt.Sprintf("./app/api/%s", dir)

	if err := os.Mkdir(dirPath, 0755); err != nil {
		panic(err)
	}

	GenerateAPIAction(name, "index")
	generateAPIRoutes(name)

}

func generateAPIRoutes(name string) {
	apiName := apiDirName(name)

	targetPath := strings.Join([]string{
		".",
		"app",
		"api",
		apiName,
		"routes.go",
	}, "/")

	helper.ExecTemplate(
		tpl.Routes(),
		targetPath,
		APIGenerator{name, helper.ToCamel(name), apiName},
	)
}

// func generateAPIModels(topDir, name string) {
// 	apiName := apiDirName(name)

// 	targetPath := strings.Join([]string{
// 		".",
// 		topDir,
// 		"api",
// 		apiName,
// 		"models.go",
// 	}, "/")

// 	helper.ExecTemplate(
// 		"./cli/template/api/models.txt",
// 		targetPath,
// 		APIGenerator{name, helper.ToCamel(name), apiName},
// 	)
// }

// func generateAPIServices(topDir, name string) {
// 	apiName := apiDirName(name)

// 	targetPath := strings.Join([]string{
// 		".",
// 		topDir,
// 		"api",
// 		apiName,
// 		"services.go",
// 	}, "/")

// 	helper.ExecTemplate(
// 		"./cli/template/api/services.txt",
// 		targetPath,
// 		APIGenerator{name, helper.ToCamel(name), apiName},
// 	)
// }

func apiDirName(name string) string {
	return fmt.Sprintf("%s_api", name)
}
