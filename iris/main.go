package main

import (
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
	v1 "github.com/snowlyg/iris-admin/server/web/web_iris/modules/v1"
)

func main() {
	wi := web_iris.Init()
	v1Party := web_iris.Party{
		Perfix:    "/api/v1",
		PartyFunc: v1.Party(),
	}
	wi.AddModule(v1Party)
	web.Start(wi)
}
