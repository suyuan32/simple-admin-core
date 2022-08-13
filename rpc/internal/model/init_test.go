package model

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	db := GormMysql()
	//db.AutoMigrate(&User{}, &Role{}, &MenuParam{}, &Menu{})
	db.AutoMigrate(&Api{})

	//var r Role
	//_ = db.Preload("Menus").Preload("Menus.Children").Where(&Role{
	//	Model: gorm.Model{ID: 1},
	//}).First(&r)
	//fmt.Println(r.CreatedAt.UnixMilli())
	//fmt.Printf("%+v", r)
}

func TestCasbin(t *testing.T) {
	//db := GormMysql()
	//csb := InitCasbin(db)
	//policy, err := csb.AddPolicy()
	//if err != nil {
	//	return
	//}
}

func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "ryan:loQHLKBZztr1Isfs@tcp(101.132.124.135:3306)/go-easy-boot?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}

func InitCasbin(db *gorm.DB) *casbin.SyncedEnforcer {
	var syncedEnforcer *casbin.SyncedEnforcer
	a, _ := gormadapter.NewAdapterByDB(db)
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		logx.Error("InitCasbin: import model fail!", err)
		return nil
	}
	syncedEnforcer, err = casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		logx.Error("InitCasbin: NewSyncedEnforcer fail!", err)
		return nil
	}
	err = syncedEnforcer.LoadPolicy()
	if err != nil {
		logx.Error("InitCasbin: LoadPolicy fail!", err)
		return nil
	}
	return syncedEnforcer
}
