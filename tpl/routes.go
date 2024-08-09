package tpl

func Routes() string {
	return `package {{.APIName}}

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	g := r.Group("/{{.Mod}}")
	{
		g.GET("/index", IndexAction)
	}
}
	`
}
