package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func QueryUserInfo(condition string) []*UserInfo {
	rows, err := dbs.Query("SELECT * FROM UserInfo where " + condition) // ignore_security_alert SQL_INJECTION
	fmt.Printf("SELECT * FROM UserInfo where " + condition)
	var ret []*UserInfo
	for rows.Next() {
		var UserInfo UserInfo
		err := rows.Scan(&UserInfo.Id, &UserInfo.User, &UserInfo.PassWord, &UserInfo.PinCode, &UserInfo.Createtime, &UserInfo.Updatetime, &UserInfo.LoginToken)
		if err != nil {
			fmt.Println(err.Error())
		}
		ret = append(ret, &UserInfo)
	}
	checkErr(err)
	return ret
}

func UpdateUserInfo(condition string, update_data string) (isSuccess bool, msg string) {
	ret := ""
	isSuccess_ := false
	stmt, err := dbs.Prepare("Update UserInfo set " + update_data + " where " + condition + ";") // ignore_security_alert SQL_INJECTION
	res, err := stmt.Exec()
	fmt.Printf("Update UserInfo set " + update_data + " where " + condition)
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

type UserInfo struct {
	Id         string `thrift:"Id,1,optional" frugal:"1,optional,string" json:"Id,omitempty"`
	User       string `thrift:"user,2,optional" frugal:"2,optional,string" json:"user,omitempty"`
	PassWord   string `thrift:"passWord,3,optional" frugal:"3,optional,string" json:"passWord,omitempty"`
	PinCode    string `thrift:"pinCode,1,optional" frugal:"1,optional,string" json:"pinCode,omitempty"`
	Createtime string `thrift:"createtime,1,optional" frugal:"1,optional,string" json:"createtime,omitempty"`
	Updatetime string `thrift:"updatetime,1,optional" frugal:"1,optional,string" json:"updatetime,omitempty"`
	LoginToken string `thrift:"logintoken,1,optional" frugal:"1,optional,string" json:"logintoken,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// // 插入数据
// stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
// checkErr(err)

// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// checkErr(err)

// id, err := res.LastInsertId()
// checkErr(err)

// fmt.Println(id)
// // 更新数据
// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
// checkErr(err)

// res, err = stmt.Exec("astaxieupdate", id)
// checkErr(err)

// affect, err := res.RowsAffected()
// checkErr(err)

// fmt.Println(affect)

// // 查询数据
// rows, err := db.Query("SELECT * FROM userinfo")
// checkErr(err)

// for rows.Next() {
//     var uid int
//     var username string
//     var department string
//     var created time.Time
//     err = rows.Scan(&uid, &username, &department, &created)
//     checkErr(err)
//     fmt.Println(uid)
//     fmt.Println(username)
//     fmt.Println(department)
//     fmt.Println(created)
// }

// // 删除数据
// stmt, err = db.Prepare("delete from userinfo where uid=?")
// checkErr(err)

// res, err = stmt.Exec(id)
// checkErr(err)

// affect, err = res.RowsAffected()
// checkErr(err)

// fmt.Println(affect)

// db.Close()
