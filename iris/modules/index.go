package modules

import (
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin-rbac/iris/perm"
	"github.com/snowlyg/iris-admin-rbac/iris/role"
	"github.com/snowlyg/iris-admin-rbac/iris/user"
	"github.com/snowlyg/iris-admin/migration"
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
)

// 加载模块
var PartyFunc = func(wi *web_iris.WebServer) {
	// 初始化驱动
	err := multi.InitDriver(&multi.Config{DriverType: "jwt", HmacSecret: nil})
	if err != nil {
		zap_server.ZAPLOG.Panic("err")
	}
	wi.AddModule(web_iris.Party{
		Perfix:    "/api/v1",
		PartyFunc: rbac.Party(),
	})
}

//  填充数据
var SeedFunc = func(wi *web_iris.WebServer, mc *migration.MigrationCmd) {
	mc.AddMigration(perm.GetMigration(), role.GetMigration(), user.GetMigration(), operation.GetMigration())
	routes, _ := wi.GetSources()
	// 权鉴模块全部为管理员权限
	authorityTypes := map[string]int{}
	for _, route := range routes {
		authorityTypes[route["path"]] = multi.AdminAuthority
	}
	// notice : 注意模块顺序
	mc.AddSeed(perm.New(routes), role.Source, user.Source)
}
