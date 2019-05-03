package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// GetMySQLCfg get MySQL config from setting.txt
func GetMySQLCfg() (string, string, string, error) {
	binary, err := ioutil.ReadFile("setting.txt")
	if err != nil {
		fmt.Println(err)
		return "", "", "", err
	}
	var account, password, dbName string
	str := string(binary)
	strAry := strings.Split(str, "\n")
	for i := 0; i < len(strAry); i++ {
		s := strings.Replace(strAry[i], "\r", "", -1)
		s = strings.Replace(s, " ", "", -1)
		if s == "[MySQL]" {
			if len(strAry) < i+4 {
				return "", "", "", errors.New("string format error")
			}
			for j := 1; j < 4; j++ {
				key, value, err := dataHandler(strAry[i+j])
				if err == nil {
					if key == "account" {
						account = value
					} else if key == "password" {
						password = value
					} else if key == "dbName" {
						dbName = value
					}
				}
			}
		}
	}

	if len(account) == 0 || len(password) == 0 || len(dbName) == 0 {
		return "", "", "", errors.New("string format error")
	}
	return account, password, dbName, nil
}

func dataHandler(str string) (string, string, error) {
	s := strings.Replace(str, " ", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	ss := strings.Split(s, "=")
	if len(ss) > 1 {
		return ss[0], ss[1], nil
	}
	return "", "", errors.New("string format error")
}
