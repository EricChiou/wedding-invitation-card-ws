package apply

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../constants"
	"../database"
	"../router"
	"../util"
	"../vo"
)

// INIT initial api
func INIT() {
	if db.MySQL == nil {
		return
	}
	// init api
	router.POST("/apply/add", add)
}

func add(context *router.Context) {
	context.Res.Header().Set("Access-Control-Allow-Origin", "www.calicomoo.ml, calicomoo.ml")

	body, err := ioutil.ReadAll(context.Req.Body)
	defer context.Req.Body.Close()
	if err != nil {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "no have req body")))
		return
	}

	var applicant vo.Applicant
	err = json.Unmarshal(body, &applicant)
	if err != nil {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "json parser fail")))
		return
	}

	if len(applicant.Name) == 0 {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "沒輸入姓名")))
		return
	}
	if applicant.Number <= 0 {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "沒輸入人數")))
		return
	}
	if len(applicant.Phone) == 0 && len(applicant.Email) == 0 && len(applicant.Line) == 0 {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "沒輸入聯絡方式")))
		return
	}
	if len(applicant.Relation) == 0 {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "沒輸入與新人關係")))
		return
	}
	if applicant.Card && len(applicant.Address) == 0 {
		context.Res.Write([]byte(util.ResultHandler(cons.Result.FormatError, "沒輸入寄送地址")))
		return
	}

	_, err = db.MySQL.Exec("INSERT INTO user_list (name, number, vegetarian, phone, email, line, relation, card, address, other) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		applicant.Name,
		applicant.Number,
		applicant.Vegetarian,
		applicant.Phone,
		applicant.Email,
		applicant.Line,
		applicant.Relation,
		applicant.Card,
		applicant.Address,
		applicant.Other)
	if err != nil {
		fmt.Println(err)
		context.Res.Write([]byte(util.ResultHandler(cons.Result.DBError, "insert db fail")))
		return
	}

	context.Res.Write([]byte(util.ResultHandler(cons.Result.Success, "")))

}
