/*
 Cmd is a tool for iris-admin.
 You can use it easy, like this: `iris-admin init`.
 It has some helpful commands for build your program. Example  init , migrate , rollback and so on.

 COMMAND
 Init is a command for initialize your program.
 Migrate is a command for migrate your database.
 Rollback is a command for rollback your migrations which you executed .
*/

package main

import (
	"github.com/snowlyg/iris-admin/modules/migration"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
	"github.com/spf13/cobra"
)

var MigrationId string

func main() {
	// cmdInit 初始化项目
	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "initialize program and set config",
		Long:  `initialize this program by set it's config param before migrate migration and seed data`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return web_iris.InitConfig()
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "migrate",
		Short: "exec run migration",
		Long:  `exec run  migrations which are you writed in  migrate.go file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return baseMigration().Migrate()
		},
	}

	var cmdRefresh = &cobra.Command{
		Use:   "refresh",
		Short: "exec refresh migration",
		Long:  `exec refresh  migrations which are you writed in  migrate.go file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return baseMigration().Refresh()
		},
	}

	var cmdRollback = &cobra.Command{
		Use:   "rollback",
		Short: "exec rollback",
		Long:  `exec rollback migrate command which are you execed`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return baseMigration().Rollback(MigrationId)
		},
	}
	cmdRollback.PersistentFlags().StringVarP(&MigrationId, "to", "t", "", "Rollback to migration id")

	var cmdSeed = &cobra.Command{
		Use:   "seed",
		Short: "exec seed",
		Long:  `exec seed  command which are you execed`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return baseMigration().Seed()
		},
	}
	cmdRollback.PersistentFlags().StringVarP(&MigrationId, "to", "t", "", "Rollback to migration id")

	var rootCmd = &cobra.Command{Use: "iris-admin"}
	rootCmd.AddCommand(cmdInit, cmdRun, cmdRollback, cmdRefresh, cmdSeed)
	rootCmd.Execute()
}

// baseMigration 实现自己的迁移逻辑
func baseMigration() *migration.MigrationCmd {
	mc := migration.New(true)
	return mc
}
