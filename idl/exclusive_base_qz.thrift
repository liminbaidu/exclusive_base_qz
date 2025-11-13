include "base.thrift"
include "common.thrift"

namespace go exclusive_base_qz

service ItemService {
    CommonLoginResponse CommonLogin (1: CommonLoginRequest req)
    CommonUpdateUserInfoResponse CommonUpdateUserInfo (1: CommonUpdateUserInfoRequest req)
    CommonSignOutResponse CommonSignOut (1: CommonSignOutRequest req)

    CreateIncomeResponse CreateIncome (1: CreateIncomeRequest req)
    UpdateIncomeResponse UpdateIncome (1: UpdateIncomeRequest req)
    DeleteIncomeResponse DeleteIncome (1: DeleteIncomeRequest req)
    QueryIncomeResponse QueryIncome (1: QueryIncomeRequest req)


    CreateDiaryResponse CreateDiary (1: CreateDiaryRequest req)
    UpdateDiaryResponse UpdateDiary (1: UpdateDiaryRequest req)
    DeleteDiaryResponse DeleteDiary (1: DeleteDiaryRequest req)
    QueryDiaryResponse QueryDiary (1: QueryDiaryRequest req)
}

struct CommonLoginRequest {
    1: optional string user,
    2: optional string password,
    3: optional string token
}

struct CommonLoginResponse {
    1: optional string token,
   255: optional base.BaseResp BaseResp,
}

struct CommonUpdateUserInfoRequest {
    1: optional string ping,
    2: optional string new_password,
    3: optional string user,
    4: optional string token
}

struct CommonUpdateUserInfoResponse {
   255: optional base.BaseResp BaseResp,
}

struct CommonSignOutRequest {
    1: optional string token
}

struct CommonSignOutResponse {
   255: optional base.BaseResp BaseResp,
}



struct CreateIncomeRequest {
    1: optional string token
    2: optional string amount
    3: optional string spendType
    4: optional string spendTime
    5: optional string remark
}

struct CreateIncomeResponse {
    1: optional string incomeId
   255: optional base.BaseResp BaseResp,
}

struct UpdateIncomeRequest {
    1: optional string token
    2: optional string amount
    3: optional string spendType
    4: optional string spendTime
    5: optional string remark
    6: optional string incomeId
}

struct UpdateIncomeResponse {
    1: optional string incomeId
   255: optional base.BaseResp BaseResp,
}

struct DeleteIncomeRequest {
    1: optional string token
    2: optional string incomeId
}

struct DeleteIncomeResponse {
   255: optional base.BaseResp BaseResp,
}

struct QueryIncomeRequest {
    1: optional string token
    2: optional string minAmount
    3: optional string maxAmount
    4: optional string spendType
    5: optional string spendStartTime
    6: optional string spendEndTime
    7: optional string incomeId
    8: optional string remark
    9: optional string page
    10: optional string size
}

struct QueryIncomeResponse {
    1: optional list<IncomeInfo> IncomeInfo
    2: optional string total
    3: optional IncomeCount IncomeCount
   255: optional base.BaseResp BaseResp
}

struct IncomeInfo{
    1: optional string amount
    2: optional string spendType
    3: optional string spendTime
    4: optional string incomeId
    5: optional string remark
}

struct IncomeCount{
    1: optional map<string, string> month
    2: optional map<string, string> week
}


struct CreateDiaryRequest {
    1: optional string token
    2: optional string content
}

struct CreateDiaryResponse {
    1: optional string diaryId
   255: optional base.BaseResp BaseResp,
}

struct UpdateDiaryRequest {
    1: optional string token
    2: optional string content
    3: optional string DiaryId
}

struct UpdateDiaryResponse {
    1: optional string diaryId
   255: optional base.BaseResp BaseResp,
}

struct DeleteDiaryRequest {
    1: optional string token
    2: optional string diaryId
}

struct DeleteDiaryResponse {
   255: optional base.BaseResp BaseResp,
}

struct QueryDiaryRequest {
    1: optional string token
    2: optional string postStartTime
    3: optional string postEndTime
    4: optional string updateStartTime
    5: optional string updateEndTime
    6: optional string diaryId
    7: optional string page
    8: optional string size
}

struct QueryDiaryResponse {
    1: optional list<DiaryInfo> diaryList
    2: optional string total
   255: optional base.BaseResp BaseResp
}


struct DiaryInfo {
    1: optional string diaryId
    2: optional string content
    3: optional string postTime
    4: optional string updateTime
}

