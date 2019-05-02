package apply

import (
	"encoding/json"
	"io/ioutil"

	"../database"
	"../router"
	"../util"
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
	body, err := ioutil.ReadAll(context.Req.Body)
	defer context.Req.Body.Close()
	if err != nil {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "no have req body")))
		return
	}

	var applicant Applicant
	err = json.Unmarshal(body, &applicant)
	if err != nil {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "json parser fail")))
		return
	}

	if len(applicant.Name) == 0 {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "沒輸入姓名")))
		return
	}
	if applicant.Number <= 0 {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "沒輸入人數")))
		return
	}
	if len(applicant.Phone) == 0 && len(applicant.Email) == 0 && len(applicant.Line) == 0 {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "沒輸入聯絡方式")))
		return
	}
	if len(applicant.Relation) == 0 {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "沒輸入與新人關係")))
		return
	}
	if applicant.Card && len(applicant.Address) == 0 {
		context.Res.Write([]byte(util.ResultHandler(util.Result.FormatError, "沒輸入寄送地址")))
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
		context.Res.Write([]byte(util.ResultHandler(util.Result.DBError, "insert db fail")))
		return
	}

	context.Res.Write([]byte(util.ResultHandler(util.Result.Success, "")))

}
