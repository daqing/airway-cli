package tpl

func Action() string {
	return `package {{.APIName}}

import (
	"github.com/gin-gonic/gin"
)

type {{.Name}}Params struct {
}

func {{.Name}}Action(c *gin.Context) {
	var p {{.Name}}Params

	if err := c.BindJSON(&p); err != nil {
		resp.Error(c, err)
		return
	}
}
	`
}
