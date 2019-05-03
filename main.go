package main

import (
	"fmt"
	"wedding-invitation-card-ws/util"

	"./apply"
	"./database"
	"./router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	account, password, dbName, err := util.GetMySQLCfg()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.MySQLInit(account, password, dbName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// init api
	apply.INIT()

	// init http server
	router.INIT(6200, nil)
	router.RUN()
}
