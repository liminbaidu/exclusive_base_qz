// ignore_security_alert_file SQL_INJECTION
package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func QueryRecordIncome(condition string) []*RecordIncome {
	rows, err := dbs.Query("SELECT * FROM RecordIncome where " + condition + ";")
	fmt.Printf("SELECT * FROM RecordIncome where " + condition)
	var ret []*RecordIncome
	for rows.Next() {
		var RecordIncome RecordIncome
		err := rows.Scan(&RecordIncome.Id, &RecordIncome.UserId, &RecordIncome.Amount, &RecordIncome.SpendType, &RecordIncome.SpendTime, &RecordIncome.Remark, &RecordIncome.Createtime, &RecordIncome.Updatetime)
		if err != nil {
			fmt.Println(err.Error())
		}
		ret = append(ret, &RecordIncome)
	}
	checkErr(err)
	return ret
}

func UpdateRecordIncome(condition string, update_data string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("Update RecordIncome set " + update_data + " where " + condition + ";")
	fmt.Printf("Update RecordIncome set " + update_data + " where " + condition)
	res, err := stmt.Exec()
	if err != nil {
		ret = "执行失败"
	}
	affect, err := res.RowsAffected()
	if affect > int64(0) {
		isSuccess_ = true
	}

	ret = "执行成功" + fmt.Sprint(affect) + "条"
	return isSuccess_, ret
}
func DeleteRecordIncome(condition string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("Delete From RecordIncome where " + condition + ";")
	fmt.Printf("Delete From RecordIncome where " + condition)
	res, err := stmt.Exec()
	if err != nil {
		ret = "执行失败"
	}
	affect, err := res.RowsAffected()
	if affect > int64(0) {
		isSuccess_ = true
	}

	ret = "执行成功" + fmt.Sprint(affect) + "条"
	return isSuccess_, ret
}

func InsertRecordIncome(insertDate string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("INSERT INTO RecordIncome(userId, amount, spendType, spendTime, remark) values" + insertDate + ";")
	fmt.Printf("INSERT INTO RecordIncome(userId, amount, spendType, spendTime, remark) values" + insertDate)
	res, err := stmt.Exec()
	if err != nil {
		ret = "执行失败"
	}
	affect, err := res.RowsAffected()
	if affect > int64(0) {
		isSuccess_ = true
	}

	ret = "执行成功" + fmt.Sprint(affect) + "条"
	return isSuccess_, ret
}

type RecordIncome struct {
	Id         string `thrift:"Id,1,optional" frugal:"1,optional,string" json:"Id,omitempty"`
	UserId     string `thrift:"userId,2,optional" frugal:"2,optional,string" json:"userId,omitempty"`
	Amount     string `thrift:"amount,3,optional" frugal:"3,optional,string" json:"amount,omitempty"`
	SpendType  string `thrift:"spendType,1,optional" frugal:"1,optional,string" json:"spendType,omitempty"`
	SpendTime  string `thrift:"spendTime,1,optional" frugal:"1,optional,string" json:"spendTime,omitempty"`
	Remark     string `thrift:"remark,1,optional" frugal:"1,optional,string" json:"remark,omitempty"`
	Updatetime string `thrift:"updatetime,1,optional" frugal:"1,optional,string" json:"updatetime,omitempty"`
	Createtime string `thrift:"createtime,1,optional" frugal:"1,optional,string" json:"createtime,omitempty"`
}
