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

	router.SetHeader("Access-Control-Allow-Origin", "https://www.calicomoo.ml")
	router.SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// start https server
	fmt.Println("start server at port 6200")
	err = http.ListenAndServeTLS(":6200", "/etc/letsencrypt/live/www.calicomoomoo.ml/fullchain.pem", "/etc/letsencrypt/live/www.calicomoo.ml/privkey.pem", nil)
	if err != nil {
		fmt.Println("start server error: ", err)
	}
}
