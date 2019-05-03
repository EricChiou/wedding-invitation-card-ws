package main

import (
	"fmt"

	"./apply"
	"./database"
	"./router"
	"./util"

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
	router.GET("/helloWorld", helloWorld)
	apply.INIT()

	// init http server
	router.INIT(6200, nil)
	router.RUNSSL("server.crt", "server.key")
}

func helloWorld(context *router.Context) {
	context.Res.Write([]byte("Hello World!!"))
}
