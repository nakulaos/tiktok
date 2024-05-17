package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func finalizer(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	// connect to the database first.
	db, err := sqlx.Connect("mysql", "root:asdasd@tcp(127.0.0.1:3306)/tiktok")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	// need to control by user, not the package
	runtime.SetFinalizer(db, finalizer)

	// Initialize a Sqlx adapter and use it in a Casbin enforcer:
	// The adapter will use the Sqlite3 table name "casbin_rule_test",
	// the default table name is "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := sqlxadapter.NewAdapter(db, "casbin_rule")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	if err != nil {
		panic(err)
	}

	// Load the policy from DB.
	if err = e.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}
	//rules := [][]string{
	//	{
	//		"alice", "data1", "read",
	//	},
	//}
	//e.AddPolicies(rules)

	// Check the permission.
	res, err := e.GetRolesForUser("lisi")
	fmt.Println(res)
	allSubjects := e.GetAllSubjects()
	fmt.Println(allSubjects)
	has, err := e.Enforce("zhangsan", "/user/info", "POST")

	if err != nil {
		log.Println("Enforce failed, err: ", err)
	}
	if !has {
		log.Println("do not have permission")
	} else {
		log.Println(" have permission")
	}

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	if err = e.SavePolicy(); err != nil {
		log.Println("SavePolicy failed, err: ", err)
	}
}
