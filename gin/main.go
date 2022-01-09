package main

import (
	"path/filepath"

	"github.com/snowlyg/iris-admin-example/gin/g"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

func main() {
	// 初始化项目路径,统一 go run ,go build ,go test 三种方式下项目绝对路径
	err := g.InitRootDir()
	if err != nil {
		panic(err)
	}
	// 初始化 gin web 项目

	wi := web_gin.Init()
	// 添加静态页面
	wi.AddWebStatic(filepath.Join(g.Root, "dist/admin"), "/admin", "/admin_static")
	wi.AddWebStatic(filepath.Join(g.Root, "dist/client"), "/client", "/client_static")
	// 增加权鉴api
	v1 := wi.GetRouterGroup("/api/v1")
	{
		rbac.Party(v1)
	}
	// 启动项目
	web.Start(wi)
}
