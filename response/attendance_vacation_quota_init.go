package response

type AttendanceVacationQuotaInit struct {
	Response
	Result []AttendanceVacationQuotaInitResult `json:"result"`
}

type AttendanceVacationQuotaInitResult struct {
	//失败原因
	Reason string `json:"reason"`
	//失败记录
	Quota AttendanceVacationQuotaInitQuota `json:"quota"`
}

type AttendanceVacationQuotaInitQuota struct {
	//假期类型唯一标识
	LeaveCode string `json:"leave_code"`
	//员工的userId
	UserId string `json:"userid"`
	//额度所对应的周期
	QuotaCycle string `json:"quota_cycle"`
}
