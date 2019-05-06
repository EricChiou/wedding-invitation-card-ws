package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

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
	http.Serve(autocert.NewListener("www.calicomoo.ml", "calicomoo.ml"), nil)
	fmt.Println("start server at port 6200")
	if err := http.ListenAndServeTLS(":6200", "", "", nil); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}
