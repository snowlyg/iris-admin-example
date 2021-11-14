package main

import (
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

func main() {
	wi := web_gin.Init()
	web.Start(wi)
}
