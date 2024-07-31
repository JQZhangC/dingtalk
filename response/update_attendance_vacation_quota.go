package response

type UpdateAttendanceVacationQuota struct {
	Response
	Result []UpdateAttendanceVacationQuotaResult `json:"result"`
}

type UpdateAttendanceVacationQuotaResult struct {
	//失败原因
	Reason string `json:"reason"`
	//失败记录
	Quota []AttendanceVacationQuotaInitQuota `json:"quota"`
}
