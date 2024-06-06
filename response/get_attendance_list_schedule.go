package response

type GetAttendanceListSchedule struct {
	Response
	AttendanceListSchedule struct {
		// 是否还有下一页
		HasMore bool `json:"has_more"`
		// 排班表
		Schedules []AttendanceSchedule `json:"schedules"`
	} `json:"result"`
}

type AttendanceSchedule struct {
	// 排班ID
	PlanID int64 `json:"plan_id"`
	// 打卡类型：Onduty：上班打卡 OffDuty：下班打卡
	CheckType string `json:"check_type"`
	// 考勤的用户userId
	UserID string `json:"userid"`
	// 考勤班次ID
	ClassID int64 `json:"class_id"`
	// 班次配置ID
	ClassSettingID int64 `json:"class_setting_id"`
	// 打卡时间
	PlanCheckTime string `json:"plan_check_time"`
	// 考勤组ID
	GroupID int64 `json:"group_id"`
	// 调整后的卡点时间
	ChangedCheckTime string `json:"changed_check_time"`
}
