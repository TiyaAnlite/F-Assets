package main

import (
	"flag"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	. "github.com/TiyaAnlite/F-Assests/types"
	"github.com/TiyaAnlite/FocotServicesCommon/dbx"
	"github.com/TiyaAnlite/FocotServicesCommon/echox"
	"github.com/TiyaAnlite/FocotServicesCommon/envx"
	"github.com/TiyaAnlite/FocotServicesCommon/utils"
	"github.com/duke-git/lancet/v2/xerror"
	"k8s.io/klog/v2"
	"testing"
)

type config struct {
	dbx.DBConfig
	echox.EchoConfig
}

var (
	cfg       = &config{}
	db        = &dbx.GormHelper{}
	snowFlake *snowflake.Snowflake
)

func init() {
	testing.Init()
	// klog.InitFlags(nil)
	flag.Parse()
	envx.MustLoadEnv(cfg)
	if err := db.Open(&cfg.DBConfig, dbx.MySQLProvider); err != nil {
		klog.Fatalf("cannot connect to MySQL: %s", err.Error())
	}
	// Migrate
	migrator := db.DB().Migrator()
	xerror.TryUnwrap("", migrator.AutoMigrate(&Position{}))
	xerror.TryUnwrap("", migrator.AutoMigrate(&Asset{}))
	xerror.TryUnwrap("", migrator.AutoMigrate(&Record{}))
	xerror.TryUnwrap("", migrator.AutoMigrate(&Book{}))
	xerror.TryUnwrap("", migrator.AutoMigrate(&CD{}))
	snowFlake, _ = snowflake.NewSnowflake(int64(0), int64(0))
}

func main() {
	go echox.Run(&cfg.EchoConfig, setupRoutes)
	klog.Infof("fire...")
	utils.Wait4CtrlC()
	klog.Infof("closing...")
	echox.Shutdown(&cfg.EchoConfig)
}
