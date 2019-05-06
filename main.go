package main

import (
	"crypto/tls"
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

	// ssl setting
	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("www.calicomoo.ml", "calicomoo.ml"),
	// 	Cache:      autocert.DirCache("/opt/WS/ssl"),
	// }
	cert, err := tls.LoadX509KeyPair("/opt/ssl/domain.crt", "/opt/ssl/domain.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair error: ", err)
		return
	}

	// start https server
	s := &http.Server{
		Addr: ":6200",
		// TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}
	fmt.Println("start server at port 6200")
	err = s.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("start server error: ", err)
	}
}
