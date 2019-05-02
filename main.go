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
	err := db.MySQLInit(util.DBConnect.Account, util.DBConnect.Password, util.DBConnect.DBName)
	if err != nil {
		fmt.Println(err)
		return
	}

	apply.INIT()

	router.INIT(6200, nil)
	router.RUN()
}
