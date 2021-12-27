package main

import (
	"os"
	"path/filepath"

	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

func main() {

	// 获取静态文件绝对路径
	gwd, _ := os.Getwd()
	web.CONFIG.System.StaticAbsPath = filepath.Join(gwd, "dist")

	wi := web_iris.Init()
	web.Start(wi)
}
