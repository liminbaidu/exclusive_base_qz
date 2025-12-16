// ignore_security_alert_file SQL_INJECTION
package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func QueryRecordContent(condition string) []*RecordContent {
	rows, err := dbs.Query("SELECT * FROM RecordContent where " + condition + ";")
	fmt.Printf("SELECT * FROM RecordContent where " + condition)
	var ret []*RecordContent
	for rows.Next() {
		var RecordContent RecordContent
		err := rows.Scan(&RecordContent.Id, &RecordContent.UserId, &RecordContent.Content, &RecordContent.PostTime, &RecordContent.Updatetime, &RecordContent.Createtime)
		if err != nil {
			fmt.Println(err.Error())
		}
		ret = append(ret, &RecordContent)
	}
	checkErr(err)
	return ret
}

func UpdateRecordContent(condition string, update_data string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("Update RecordContent set " + update_data + " where " + condition + ";")
	fmt.Printf("Update RecordContent set " + update_data + " where " + condition)
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

func DeleteRecordContent(condition string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("Delete from RecordContent where " + condition + ";")
	fmt.Printf("Delete from RecordContent where " + condition)
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

func InsertRecordContent(insertDate string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("INSERT INTO RecordContent(userId, content, postTime, updateTime) values" + insertDate + ";")
	fmt.Printf("INSERT INTO RecordContent(userId, content, postTime, updateTime) values" + insertDate)
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

type RecordContent struct {
	Id         string `thrift:"Id,1,optional" frugal:"1,optional,string" json:"Id,omitempty"`
	UserId     string `thrift:"userId,2,optional" frugal:"2,optional,string" json:"userId,omitempty"`
	Content    string `thrift:"content,3,optional" frugal:"3,optional,string" json:"content,omitempty"`
	PostTime   string `thrift:"postTime,1,optional" frugal:"1,optional,string" json:"postTime,omitempty"`
	Updatetime string `thrift:"updatetime,1,optional" frugal:"1,optional,string" json:"updatetime,omitempty"`
	Createtime string `thrift:"createtime,1,optional" frugal:"1,optional,string" json:"createtime,omitempty"`
}
