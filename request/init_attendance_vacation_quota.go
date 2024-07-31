package request

type InitAttendanceVacationQuota struct {
	// 当前企业内拥有OA审批应用权限的管理员的userId
	OpUserid string `json:"op_userid"`
	// 假期额度信息
	LeaveQuotas LeaveQuotas `json:"leave_quotas"`
}

type LeaveQuotas struct {
	// 员工的userId
	Userid string `json:"userid" validate:"required"`
	// 额度有效期结束时间，毫秒级时间戳
	EndTime int64 `json:"end_time" validate:"required"`
	// 额度有效期开始时间，毫秒级时间戳
	StartTime int64 `json:"start_time" validate:"required"`
	// 假期类型唯一标识
	LeaveCode string `json:"leave_code" validate:"required"`
	// 操作原因
	Reason string `json:"reason"`
	//以天计算的额度总数, 假期类型按天计算时，该值不为空且按百分之一天折算 例如：1000=10天。
	QuotaNumPerDay int64 `json:"quota_num_per_day"`
	//以小时计算的额度总数, 假期类型按小时，计算该值不为空且按百分之一小时折算。例如：1000=10小时。
	QuotaNumPerHour int64 `json:"quota_num_per_hour"`
	// 额度所对应的周期,格式必须满足“yyyy”。
	QuotaCycle string `json:"quota_cycle"`
}
