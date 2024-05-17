package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewCasbin(dataSource string, conf string, table string) *casbin.Enforcer {
	//db, err := sqlx.Connect("mysql", dataSource)
	//if err != nil {
	//	panic(err)
	//}
	//sqlxAdapt := sqlxadapter.NewAdapterFromOptions(&sqlxadapter.AdapterOptions{DB: db})

	//dataSource0 := strings.Split(dataSource, "?")[0]

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormadapter.TurnOffAutoMigrate(db)
	adapt, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	//sqlxAdapt, err := sqlxadapter.NewAdapter(db, table)
	//if err != nil {
	//	panic(err)
	//}

	//enforce, err := casbin.NewEnforcer(conf, sqlxAdapt)
	enforce, err := casbin.NewEnforcer(conf, adapt)
	if err != nil {
		panic(err)
	}

	return enforce
}
