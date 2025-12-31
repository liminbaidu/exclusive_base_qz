package handler

import (
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"

	"encoding/json"

	"exclusive_base_qz/src/action"
	"exclusive_base_qz/util"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CommonLogin(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CommonLoginResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CommonLoginRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CommonLogin(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}
func CommonIsLogin(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CommonIsLoginResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CommonIsLoginRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CommonIsLogin(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func CommonSignout(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CommonSignOutResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CommonSignOutRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CommonSignout(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func CommonUpdateUserInfo(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CommonUpdateUserInfoResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CommonUpdateUserInfoRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CommonUpdateUserInfo(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}
