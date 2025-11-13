package action

import (
	"exclusive_base_qz/common/db"
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"
	"exclusive_base_qz/util"
	"fmt"
	"strings"
	"time"
)

func CreateDiary(req *exclusive_base_qz.CreateDiaryRequest) *exclusive_base_qz.CreateDiaryResponse {
	resp := &exclusive_base_qz.CreateDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}
	userinfo := GetUserInfoByToken(*req.Token)
	postTime := time.Now().Unix()

	isSuccess, _ := db.InsertRecordContent(fmt.Sprintf("('%s','%s','%s','%s')", userinfo[0].Id, *req.Content, util.ToString(postTime), util.ToString(postTime)))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func UpdateDiary(req *exclusive_base_qz.UpdateDiaryRequest) *exclusive_base_qz.UpdateDiaryResponse {
	resp := &exclusive_base_qz.UpdateDiaryResponse{
		DiaryId:  req.DiaryId,
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 编辑失败")
		return resp
	}
	if !CheckDiaryId(*req.DiaryId) {
		resp.BaseResp = util.ProcessBaseResp(103, "内容id不存在, 编辑失败")
		return resp
	}

	userinfo := GetUserInfoByToken(*req.Token)

	updateDataList := []string{}
	updateData := ""
	updateTime := time.Now().Unix()
	updateDataList = append(updateDataList, fmt.Sprintf("updateTime='%s'", util.ToString(updateTime)))
	if req.Content != nil && *req.Content != "" {
		updateDataList = append(updateDataList, fmt.Sprintf("content='%s'", *req.Content))
	}

	if len(updateDataList) > 0 {
		updateData = strings.Join(updateDataList, " , ")
	}

	if updateData == "" {
		resp.BaseResp = util.ProcessBaseResp(102, "无数据需要操作")
		return resp
	}

	isSuccess, _ := db.UpdateRecordContent(fmt.Sprintf("userId='%s' and Id=%s", userinfo[0].Id, *req.DiaryId), updateData)
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func DeleteDiary(req *exclusive_base_qz.DeleteDiaryRequest) *exclusive_base_qz.DeleteDiaryResponse {
	resp := &exclusive_base_qz.DeleteDiaryResponse{
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}

	if !CheckDiaryId(*req.DiaryId) {
		resp.BaseResp = util.ProcessBaseResp(103, "内容id不存在, 编辑失败")
		return resp
	}

	isSuccess, _ := db.DeleteRecordContent(fmt.Sprintf("Id=%s", *req.DiaryId))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func QueryDiary(req *exclusive_base_qz.QueryDiaryRequest) *exclusive_base_qz.QueryDiaryResponse {
	DiaryInfoList := []*exclusive_base_qz.DiaryInfo{}
	total := "0"
	resp := &exclusive_base_qz.QueryDiaryResponse{
		Total:     &total,
		DiaryList: DiaryInfoList,
		BaseResp:  util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}
	condition := ""
	conditionList := []string{}
	userinfo := GetUserInfoByToken(*req.Token)
	conditionList = append(conditionList, fmt.Sprintf("userId=%s", *&userinfo[0].Id))

	if req.DiaryId != nil && *req.DiaryId != "" {
		conditionList = append(conditionList, fmt.Sprintf("Id=%s", *req.DiaryId))
	}

	if req.PostEndTime != nil && *req.PostEndTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("postTime<=%s", *req.PostEndTime))
	}

	if req.PostStartTime != nil && *req.PostStartTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("postTime>=%s", *req.PostStartTime))
	}

	if req.UpdateStartTime != nil && *req.UpdateStartTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("updateTime>=%s", *req.UpdateStartTime))
	}

	if req.UpdateEndTime != nil && *req.UpdateEndTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("updateTime<=%s", *req.UpdateEndTime))
	}

	if len(conditionList) > 0 {
		condition = strings.Join(conditionList, " and ")
	}
	if condition == "" {
		return resp
	}

	pagelimit := " limit 1,20"
	if req.Page != nil && req.Size != nil {
		startPage := (util.StrToInt64(*req.Page) - 1) * util.StrToInt64(*req.Page)
		if *req.Page == "0" || *req.Page == "1" {
			startPage = int64(0)
		}
		pagelimit = fmt.Sprintf(" limit %s,%s", util.ToString(startPage), *req.Size)
	}
	RecordDiaryInfo := db.QueryRecordContent(condition + " order by createtime desc" + pagelimit)

	total = util.ToString(len(RecordDiaryInfo))
	resp.Total = &total
	for _, info := range RecordDiaryInfo {
		DiaryInfo := &exclusive_base_qz.DiaryInfo{
			Content:    &info.Content,
			PostTime:   &info.PostTime,
			UpdateTime: &info.Updatetime,
			DiaryId:    &info.Id,
		}
		DiaryInfoList = append(DiaryInfoList, DiaryInfo)
	}
	resp.DiaryList = DiaryInfoList
	return resp
}

func CheckDiaryId(DiaryId string) bool {
	isPass := false
	DiaryInfo := db.QueryRecordContent(fmt.Sprintf("Id=%s", DiaryId))
	if len(DiaryInfo) > 0 {
		fmt.Printf("用户已登陆")
		isPass = true
		return isPass
	}
	return isPass
}
