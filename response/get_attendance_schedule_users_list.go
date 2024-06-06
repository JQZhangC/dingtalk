package response

type GetAttendanceScheduleUsersList struct {
	Response

	Result []AttendanceScheduleUsersList `json:"result"`
}

type AttendanceScheduleUsersList struct {

	// 考勤类型： Onduty：上班打卡 OffDuty：下班打卡
	CheckType string `json:"check_type"`
	// 计划打卡时间
	PlanDateTime string `json:"plan_check_time"`
	// 考勤组ID
	GroupID int `json:"group_id"`
	// 用户的userId
	UserID string `json:"userid"`
	// 排班绑定的审批单ID
	ApproveId int `json:"approve_id"`
	// 工作日，代表具体哪一天的排班
	WorkDate string `json:"work_date"`
	// 排班ID
	ID int `json:"id"`
	// 班次版本
	ShiftVersion int `json:"shift_version"`
	// 班次ID，该值为0，表明当天休息
	ShiftId int `json:"shift_id"`
	// 是否休息： Y：当天排休 N：不休息
	IsRest string `json:"is_rest"`
}
