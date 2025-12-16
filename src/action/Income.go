package action

import (
	"exclusive_base_qz/common/db"
	"exclusive_base_qz/kitex_gen/exclusive_base_qz"
	"exclusive_base_qz/util"
	"fmt"
	"strings"
	"time"
)

func CreateIncome(req *exclusive_base_qz.CreateIncomeRequest) *exclusive_base_qz.CreateIncomeResponse {
	resp := &exclusive_base_qz.CreateIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}
	userinfo := GetUserInfoByToken(*req.Token)

	isSuccess, _ := db.InsertRecordIncome(fmt.Sprintf("('%s','%s','%s','%s','%s')", userinfo[0].Id, *req.Amount, *req.SpendType, *req.SpendTime, *req.Remark))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func UpdateIncome(req *exclusive_base_qz.UpdateIncomeRequest) *exclusive_base_qz.UpdateIncomeResponse {
	resp := &exclusive_base_qz.UpdateIncomeResponse{
		IncomeId: req.IncomeId,
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 编辑失败")
		return resp
	}
	if !CheckIncomeId(*req.IncomeId) {
		resp.BaseResp = util.ProcessBaseResp(103, "账单id不存在, 编辑失败")
		return resp
	}

	userinfo := GetUserInfoByToken(*req.Token)

	updateDataList := []string{}
	updateData := ""
	if req.Amount != nil && *req.Amount != "" {
		updateDataList = append(updateDataList, fmt.Sprintf("amount='%s'", *req.Amount))
	}

	if req.SpendType != nil && *req.SpendType != "" {
		updateDataList = append(updateDataList, fmt.Sprintf("spendType='%s'", *req.SpendType))
	}

	if req.SpendTime != nil && *req.SpendTime != "" {
		updateDataList = append(updateDataList, fmt.Sprintf("spendTime='%s'", *req.SpendTime))
	}
	if req.Remark != nil && *req.Remark != "" {
		updateData = updateData + fmt.Sprintf("remark='%s'", *req.Remark)
	}

	if len(updateDataList) > 0 {
		updateData = strings.Join(updateDataList, " , ")
	}

	if updateData == "" {
		resp.BaseResp = util.ProcessBaseResp(102, "无数据需要操作")
		return resp
	}

	isSuccess, _ := db.UpdateRecordIncome(fmt.Sprintf("userId='%s' and Id=%s", userinfo[0].Id, *req.IncomeId), updateData)
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func DeleteIncome(req *exclusive_base_qz.DeleteIncomeRequest) *exclusive_base_qz.DeleteIncomeResponse {
	resp := &exclusive_base_qz.DeleteIncomeResponse{
		BaseResp: util.ProcessBaseResp(0, "操作成功"),
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}

	if !CheckIncomeId(*req.IncomeId) {
		resp.BaseResp = util.ProcessBaseResp(103, "账单id不存在, 编辑失败")
		return resp
	}

	isSuccess, _ := db.DeleteRecordIncome(fmt.Sprintf("Id=%s", *req.IncomeId))
	if !isSuccess {
		resp.BaseResp = util.ProcessBaseResp(101, "Db操作失败, 请重试")
		return resp
	}
	return resp
}

func QueryIncome(req *exclusive_base_qz.QueryIncomeRequest) *exclusive_base_qz.QueryIncomeResponse {
	IncomeInfoList := []*exclusive_base_qz.IncomeInfo{}
	total := "0"
	resp := &exclusive_base_qz.QueryIncomeResponse{
		Total:       &total,
		IncomeInfo:  IncomeInfoList,
		BaseResp:    util.ProcessBaseResp(0, "操作成功"),
		IncomeCount: &exclusive_base_qz.IncomeCount{},
	}

	if !CheckIsLoginIn(*req.Token) {
		resp.BaseResp = util.ProcessBaseResp(100, "未识别登录态, 新建失败")
		return resp
	}
	condition := ""
	conditionList := []string{}
	userinfo := GetUserInfoByToken(*req.Token)
	conditionList = append(conditionList, fmt.Sprintf("userId=%s", *&userinfo[0].Id))

	if req.IncomeId != nil && *req.IncomeId != "" {
		conditionList = append(conditionList, fmt.Sprintf("Id=%s", *req.IncomeId))
	}

	if req.MaxAmount != nil && *req.MaxAmount != "" {
		conditionList = append(conditionList, fmt.Sprintf("amount<=%s", *req.MaxAmount))
	}

	if req.MinAmount != nil && *req.MinAmount != "" {
		conditionList = append(conditionList, fmt.Sprintf("amount>=%s", *req.MinAmount))
	}

	if req.SpendEndTime != nil && *req.SpendEndTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("spendTime<=%s", *req.SpendEndTime))
	}

	if req.SpendStartTime != nil && *req.SpendStartTime != "" {
		conditionList = append(conditionList, fmt.Sprintf("spendTime>=%s", *req.SpendStartTime))
	}
	if req.SpendType != nil && *req.SpendType != "" {
		conditionList = append(conditionList, fmt.Sprintf("spendType in (%s)", *req.SpendType))
	}

	if req.Remark != nil && *req.Remark != "" {
		conditionList = append(conditionList, "remark like '%"+util.ToString(*req.Remark)+"%'")
	}
	if len(conditionList) > 0 {
		condition = strings.Join(conditionList, " and ")
	}
	if condition == "" {
		return resp
	}

	pagelimit := " limit 0,19"
	if req.Page != nil && req.Size != nil {
		startPage := (util.StrToInt64(*req.Page) - 1) * util.StrToInt64(*req.Size)
		if *req.Page == "0" || *req.Page == "1" {
			startPage = int64(0)
		}
		pagelimit = fmt.Sprintf(" limit %s,%s", util.ToString(startPage), util.ToString(util.StrToInt64(*req.Size)-1))
	}
	RecordIncomeInfo := db.QueryRecordIncome(condition + " order by createtime desc" + pagelimit)
	allRecordIncomeInfo := db.QueryRecordIncome("Id!=0")

	total = util.ToString(len(allRecordIncomeInfo))
	resp.Total = &total
	IncomeCountWeek := map[string]string{}
	IncomeCountMonth := map[string]string{}
	for _, info := range allRecordIncomeInfo {
		IncomeCountWeekkey := UnixToWeek(util.StrToInt64(info.SpendTime))
		IncomeCountMonthkey := UnixToMonth(util.StrToInt64(info.SpendTime))
		if _, ok := IncomeCountWeek[IncomeCountWeekkey]; ok {
			WeekAmount := util.StrToInt64(info.Amount) + util.StrToInt64(IncomeCountWeek[IncomeCountWeekkey])
			IncomeCountWeek[IncomeCountWeekkey] = util.ToString(WeekAmount)
		} else {
			IncomeCountWeek[IncomeCountWeekkey] = info.Amount
		}
		if _, ok := IncomeCountMonth[IncomeCountMonthkey]; ok {
			MonthAmount := util.StrToInt64(info.Amount) + util.StrToInt64(IncomeCountMonth[IncomeCountMonthkey])
			IncomeCountMonth[IncomeCountMonthkey] = util.ToString(MonthAmount)
		} else {
			IncomeCountMonth[IncomeCountMonthkey] = info.Amount
		}
	}
	for _, info := range RecordIncomeInfo {
		IncomeInfo := &exclusive_base_qz.IncomeInfo{
			Amount:    &info.Amount,
			SpendType: &info.SpendType,
			SpendTime: &info.SpendTime,
			IncomeId:  &info.Id,
			Remark:    &info.Remark,
		}
		IncomeInfoList = append(IncomeInfoList, IncomeInfo)
	}
	resp.IncomeCount.Month = IncomeCountMonth
	resp.IncomeCount.Week = IncomeCountWeek
	resp.IncomeInfo = IncomeInfoList
	return resp
}

func CheckIncomeId(incomeId string) bool {
	isPass := false
	IncomeInfo := db.QueryRecordIncome(fmt.Sprintf("Id=%s", incomeId))
	if len(IncomeInfo) > 0 {
		fmt.Printf("用户已登陆")
		isPass = true
		return isPass
	}
	return isPass
}

func UnixToMonth(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	month := t.Month()
	year := t.Year()
	ret := fmt.Sprintf("%d-%d", year, month)
	return ret
}

func UnixToWeek(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	year1, week1 := t.ISOWeek()
	ret := fmt.Sprintf("%d-第%d周", year1, week1)
	return ret
}
