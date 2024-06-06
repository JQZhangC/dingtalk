package response

type GetAttendanceScheduleShift struct {
	Response

	Result []AttendanceScheduleShift `json:"result"`
}

type AttendanceScheduleShift struct {
	// 工作日，代表具体哪一天的排班
	WorkDate string `json:"work_date"`
	// 班次名称
	ShiftNames []string `json:"shift_names"`
	// 用户的userId
	UserID string `json:"userid"`
	// 班次版本ID
	ShiftVersions []int64 `json:"shift_versions"`
	// 班次ID
	ShiftIds []int64 `json:"shift_ids"`
	// 考勤组ID
	GroupID int `json:"group_id"`
	// 企业的corpId
	CorpId string `json:"corp_id"`
}
