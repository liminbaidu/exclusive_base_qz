package handler

import (
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"

	"encoding/json"

	"exclusive_base_qz/src/action"
	"exclusive_base_qz/util"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateDiary(r *ghttp.Request) string {

	resp := &exclusive_base_qz.CreateDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.CreateDiaryRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.CreateDiary(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func UpdateDiary(r *ghttp.Request) string {

	resp := &exclusive_base_qz.UpdateDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.UpdateDiaryRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.UpdateDiary(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func DeleteDiary(r *ghttp.Request) string {

	resp := &exclusive_base_qz.DeleteDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.DeleteDiaryRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.DeleteDiary(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}

func QueryDiary(r *ghttp.Request) string {

	resp := &exclusive_base_qz.QueryDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, ""),
	}
	req := &exclusive_base_qz.QueryDiaryRequest{}

	if err := r.Parse(req); err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "请求解析失败")
	}
	resp = action.QueryDiary(req)

	obj, err := json.Marshal(resp)
	if err != nil {
		resp.BaseResp = util.ProcessBaseResp(100, "返回转义失败")
	}
	return string(obj)
}
