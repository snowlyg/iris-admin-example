package modules

import (
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/admin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin/migration"
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
)

// 加载模块
var PartyFunc = func(wi *web_gin.WebServer) {
	// 初始化驱动
	err := multi.InitDriver(&multi.Config{DriverType: "jwt", HmacSecret: nil})
	if err != nil {
		zap_server.ZAPLOG.Panic("err")
	}
	v1 := wi.GetRouterGroup("/api/v1")
	{
		rbac.Party(v1)
	}
}

//  填充数据
var SeedFunc = func(wi *web_gin.WebServer, mc *migration.MigrationCmd) {
	mc.AddMigration(api.GetMigration(), authority.GetMigration(), admin.GetMigration(), operation.GetMigration())
	routes, _ := wi.GetSources()
	// 权鉴模块全部为管理员权限
	authorityTypes := map[string]int{}
	for _, route := range routes {
		authorityTypes[route["path"]] = multi.AdminAuthority
	}
	// notice : 注意模块顺序
	mc.AddSeed(api.New(routes, authorityTypes), authority.Source, admin.Source)
}
