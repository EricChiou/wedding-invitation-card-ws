package db

import "database/sql"

// MySQL mysql
var MySQL *sql.DB

var err error

// MySQLConnect connect mysql
func MySQLInit(account string, password string, dbName string) error {
	MySQL, err = sql.Open("mysql", account+":"+password+"@/"+dbName)
	return err
}
