package main

import (
	"path/filepath"

	"github.com/snowlyg/iris-admin-example/iris/g"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

func main() {
	// 初始化项目路径,统一 go run ,go build ,go test 三种方式下项目绝对路径
	err := g.InitRootDir()
	if err != nil {
		panic(err)
	}
	// 初始化 gin web 项目
	wi := web_iris.Init()
	// 添加静态页面
	wi.AddWebStatic("/admin", "/static", filepath.Join(g.Root, "dist"))
	// 增加权鉴api
	wi.AddModule(web_iris.Party{
		Perfix:    "/api/v1",
		PartyFunc: rbac.Party(),
	})
	// 启动项目
	web.Start(wi)
}
