package action

import (
	"exclusive_base_qz/common/db"
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"
	"exclusive_base_qz/util"
	"fmt"
	"log"
	"time"
)

func CommonSignout(req *exclusive_base_qz.CommonSignOutRequest) *exclusive_base_qz.CommonSignOutResponse {
	resp := &exclusive_base_qz.CommonSignOutResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 登出失败")
		return resp
	}

	isSuccess, _ := db.UpdateUserInfo(fmt.Sprintf("logintoken='%s'", *req.Token), fmt.Sprintf("logintoken=''"))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func CommonUpdateUserInfo(req *exclusive_base_qz.CommonUpdateUserInfoRequest) *exclusive_base_qz.CommonUpdateUserInfoResponse {

	resp := &exclusive_base_qz.CommonUpdateUserInfoResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	if req.Token == nil && req.Ping == nil {
		resp.BaseResp = util.ProcessBaseResp(102, "token不存在, 修改失败")
		return resp
	}
	if req.Token != nil {
		if !CheckIsLoginIn(*req.Token) {
			resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 修改失败")
			return resp
		}
	}
	if req.Ping != nil {
		userInfo := GetUserInfoByPing(*req.Ping, *req.User)
		if len(userInfo) == 0 {
			resp.BaseResp = util.ProcessBaseResp(103, "Ping不正确, 修改失败")
			return resp
		}
		req.Token = &userInfo[0].LoginToken
	}

	isSuccess, _ := db.UpdateUserInfo(fmt.Sprintf("logintoken='%s'", *req.Token), fmt.Sprintf("password='%s'", *req.NewPassword_))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败,请重试")
		return resp
	}
	return resp
}

func CommonLogin(req *exclusive_base_qz.CommonLoginRequest) *exclusive_base_qz.CommonLoginResponse {
	cookie := ""
	resp := &exclusive_base_qz.CommonLoginResponse{
		Token:    &cookie,
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	if req.Token != nil {
		if CheckIsLoginIn(*req.Token) {
			resp.Token = req.Token
			return resp
		}
	}

	isCheckPass, err := CheckAccount(req)
	if err != "" {
		resp.BaseResp = util.ProcessBaseResp(1, err)
	}
	if isCheckPass {
		cookie = SetCookie(req)
		resp.Token = &cookie
	}
	return resp
}
func CommonIsLogin(req *exclusive_base_qz.CommonIsLoginRequest) *exclusive_base_qz.CommonIsLoginResponse {
	IsLogin := "false"
	resp := &exclusive_base_qz.CommonIsLoginResponse{
		IsLogin:  &IsLogin,
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	if req.Token != nil {
		if CheckIsLoginIn(*req.Token) {
			IsLogin = "true"
			resp.IsLogin = &IsLogin
			return resp
		}
	}
	return resp
}
func CheckAccount(req *exclusive_base_qz.CommonLoginRequest) (checkResult bool, err string) {
	isCheckPass := false
	err = ""
	UserInfo := db.QueryUserInfo(fmt.Sprintf("user='%s'", *req.User))
	if len(UserInfo) == 0 {
		err = "用户名不存在"
		return isCheckPass, err
	}
	if UserInfo[0].PassWord != *req.Password {
		err = "密码不正确"
		return isCheckPass, err
	}
	isCheckPass = true
	return isCheckPass, ""
}
func SetCookie(req *exclusive_base_qz.CommonLoginRequest) string {
	key := []byte("2fa6c1e9")
	timestp := time.Now().Unix()
	cookie := ""
	strEncrypted, err := util.Encrypt(util.ToString(timestp), key)
	if err != nil {
		log.Fatal(err)
	}
	isSuccess, _ := db.UpdateUserInfo(fmt.Sprintf("user='%s'", *req.User), fmt.Sprintf("logintoken='%s'", strEncrypted))
	if isSuccess {
		cookie = strEncrypted
		fmt.Println("Encrypted:", strEncrypted)
	}
	return cookie
}
func CheckIsLoginIn(strEncrypted string) bool {
	// key := []byte("2fa6c1e9")
	// strDecrypted, err := Decrypt(strEncrypted, key)
	isLogin := false
	UserInfo := db.QueryUserInfo(fmt.Sprintf("logintoken='%s'", strEncrypted))
	if len(UserInfo) > 0 {
		fmt.Printf("用户已登陆")
		isLogin = true
		return isLogin
	}
	return isLogin
}

func CheckPingCode(pingCode string) bool {
	// key := []byte("2fa6c1e9")
	// strDecrypted, err := Decrypt(strEncrypted, key)
	isTrue := false
	UserInfo := db.QueryUserInfo(fmt.Sprintf("pinCode='%s'", pingCode))
	if len(UserInfo) > 0 {
		isTrue = true
		return isTrue
	}
	return isTrue
}

func GetUserInfoByPing(ping string, user string) []*db.UserInfo {
	// key := []byte("2fa6c1e9")
	// strDecrypted, err := Decrypt(strEncrypted, key)
	UserInfo := db.QueryUserInfo(fmt.Sprintf("pinCode='%s' and user='%s'", ping, user))
	if len(UserInfo) > 0 {
		return UserInfo
	}
	return UserInfo
}

func GetUserInfoByToken(token string) []*db.UserInfo {
	// key := []byte("2fa6c1e9")
	// strDecrypted, err := Decrypt(strEncrypted, key)
	UserInfo := db.QueryUserInfo(fmt.Sprintf("logintoken='%s'", token))
	if len(UserInfo) > 0 {
		return UserInfo
	}
	return UserInfo
}
