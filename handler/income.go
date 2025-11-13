package handler

import (
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"

	"encoding/json"

	"exclusive_base_qz/src/action"
	"exclusive_base_qz/util"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateIncome(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CreateIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CreateIncomeRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CreateIncome(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func UpdateIncome(r *ghttp.Request) string {

	resp := &exclusive_base_qz.UpdateIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.UpdateIncomeRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.UpdateIncome(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func DeleteIncome(r *ghttp.Request) string {

	resp := &exclusive_base_qz.DeleteIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.DeleteIncomeRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.DeleteIncome(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func QueryIncome(r *ghttp.Request) string {

	resp := &exclusive_base_qz.QueryIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.QueryIncomeRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.QueryIncome(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}
