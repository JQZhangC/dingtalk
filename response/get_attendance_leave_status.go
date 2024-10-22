package response

type GetAttendanceLeaveStatus struct {
	Response
	Result GetAttendanceLeaveStatusResult `json:"result"`
}

type GetAttendanceLeaveStatusResult struct {
	HasMore     bool          `json:"has_more"`
	LeaveStatus []LeaveStatus `json:"leave_status"`
}

type LeaveStatus struct {
	DurationUnit    string `json:"duration_unit"`    //请假单位 percent_day：天,percent_hour：小时
	DurationPercent int64  `json:"duration_percent"` //假期时长*100，例如用户请假时长为1天，该值就等于100
	EndTime         int64  `json:"end_time"`         //请假结束时间，Unix时间戳
	StartTime       int64  `json:"start_time"`       //请假开始时间，Unix时间戳
	UserId          string `json:"userid"`           //用户ID
}
