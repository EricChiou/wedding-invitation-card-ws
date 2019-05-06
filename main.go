package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"./apply"
	"./database"
	"./router"
	"./util"
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

	// set api
	apply.INIT()

	// init api
	router.INIT()

	// start https server
	fmt.Println("start server at port 6200")
	err = http.ListenAndServeTLS(":6200", "/opt/ssl/crt.pem", "/opt/ssl/key.pem", nil)
	if err != nil {
		fmt.Println("start server error: ", err)
	}
}
