package response

type GetAttendanceVacationQuotaList struct {
	Response
	Result OapiLeaveQuotaUserListVo `json:"result"`
}

type OapiLeaveQuotaUserListVo struct {
	//是否存在更多记录
	HasMore     bool          `json:"has_more"`
	LeaveQuotas []LeaveQuotas `json:"leave_quotas"`
}

type LeaveQuotas struct {
	//员工的userId
	UserId string `json:"userid"`
	//假期类型唯一标识
	LeaveCode string `json:"leave_code"`
	//额度所对应的周期
	QuotaCycle string `json:"quota_cycle"`
	//配额的唯一标记
	QuotaId string `json:"quota_id"`
	//假期有效期开始时间，毫秒级时间戳
	StartTime int64 `json:"start_time"`
	//额度有效期结束时间，毫秒级时间戳
	EndTime int64 `json:"end_time"`
	//以小时计算的额度总数, 假期类型按小时，计算该值不为空且按百分之一小时折算。例如：1000=10小时。
	QuotaNumPerHour int64 `json:"quota_num_per_hour"`
	//以天计算的额度总数, 假期类型按天计算时，该值不为空且按百分之一天折算 例如：1000=10天。
	QuotaNumPerDay int64 `json:"quota_num_per_day"`
	//以天计算的使用额度
	UsedNumPerDay int64 `json:"used_num_per_day"`
	//以小时计算的使用额度
	UsedNumPerHour int64 `json:"used_num_per_hour"`
}
