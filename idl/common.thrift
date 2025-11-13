include "base.thrift"

namespace go ecom.mmc.qatools

enum PriceTag {
	NotSupport     = -2
	NoSameSku      = 1
	MinPriceSku    = 2
	NotMinPriceSku = 3
}

enum PriceType {
	OutLowestOnlinePrice = 1 // 最低价(凑单)
	OutLowestSupplyPrice = 2 // 供货价(凑单)
	OutSingleBuyPrice = 3 // 最低价(单品)
	OutSingleSupplyPrice = 4 // 供货价(单品)
	InnerPrice  = 5 // 站内最低价
	LowestSameProduct30dPriceOnline = 6 // 同店同款30天最低价
	NoSupport = 7 // 不支持比价
}

enum RoleType {
    Shop = 1
    Author = 2
    Operator = 3
}

enum PriceObject {
    ApplyPrice = 1
    EstimatePrice = 2
}

enum SignRecordStatus {
	unknown = 0
	draft = 1
	wait_first_audit = 2
	wait_second_audit = 3
	wait_third_audit = 4
	apply_succeed = 6
	audit_rejected = 7
	wait_deposit = 5
	wait_author_confirm = 8
	canceled = 11
	wait_publish = 12
	wait_set_subsidy = 13
	subsidy_audit_in_progress = 14
	wait_set_schedule = 15
	price_risk_wait_audit = 16
	wait_bid_publish = 17
	draft_with_algo_schedule = 18
}

