package response

type GetAttendanceShiftDetail struct {
	Response
	// 班次详情
	Result AttendanceShiftDetail `json:"result"`
}

type AttendanceShiftDetail struct {
	// 班次组名称
	ShiftGroupName string `json:"shift_group_name"`
	// 企业的corpId
	CorpID string `json:"corp_id"`
	// 班次ID
	ID int `json:"id"`
	// 班次名称
	Name string `json:"name"`
	// 卡段
	Sections []AttendanceShiftDetailSection `json:"sections"`
	// 班次组ID
	ShiftGroupId string `json:"shift_group_id"`
	// 班次负责人的userId
	Owner string `json:"owner"`
}

type AttendanceShiftDetailSection struct {
	// 卡段ID。一次上下班成为一个卡段，一个班次可能会有多个卡段
	ID int `json:"id"`
	// 卡点
	Punches []AttendanceShiftDetailPunch `json:"punches"`
	// 休息时段信息
	Rests []AttendanceShiftDetailRest `json:"rests"`
}

type AttendanceShiftDetailPunch struct {
	// 打卡类型：OnDuty：上班 OffDuty：下班
	CheckType string `json:"check_type"`
	// 允许的最晚的打卡时间。 单位是分钟，用打卡时间加上分钟数可以计算出最晚打卡时间。
	EndMin int `json:"end_min"`
	// 是否跨天，跨天是指打卡时间是第二天：0：不跨天 1：跨天
	Across int `json:"across"`
	// 打卡时间，Unix时间戳，仅有时分秒信息。
	CheckTime string `json:"check_time"`
	//
	AbsenteeismLateMinutes string `json:"absenteeism_late_minutes"`
	// 是否免打卡： true：免打卡 false：需打卡
	FreeCheck bool `json:"free_check"`
	// 卡点ID
	ID int `json:"id"`
	// 允许早退/迟到的时长，单位分钟
	PermitMinutes int `json:"permit_minutes"`
	// 严重早退/迟到的时长，单位分钟
	SeriousLateMinutes string `json:"serious_late_minutes"`
}

type AttendanceShiftDetailRest struct {
	// 休息时间
	CheckTime string `json:"check_time"`
	// 是否跨天，跨天是指休息时间是第二天： /0：不跨天 1：跨天
	Across int `json:"across"`
	// 休息类型： OnDuty：休息开始 OffDuty：休息结束
	CheckType string `json:"check_type"`
	// 休息点ID
	ID int `json:"id"`
}
