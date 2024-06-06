package response

type GetAttendanceScheduleList struct {
	Response
	Result []AttendanceScheduleList `json:"result"`
}

type AttendanceScheduleList struct {
	// 考勤类型： Onduty：上班打卡 OffDuty：下班打卡
	CheckType string `json:"check_type"`
	// 审批类型
	ApproveType string `json:"approve_type"`
	// 最后更新时间
	GmtModified string `json:"gmt_modified"`
	// 创建时间
	GmtCreate string `json:"gmt_create"`
	// 企业的corpId
	CorpID string `json:"corp_id"`
	// 该员工打卡时间
	CheckDateTime string `json:"check_date_time"`
	// 考勤组ID
	GroupID int `json:"group_id"`
	// 班次名称
	ClassName string `json:"class_name"`
	// 用户的userId
	UserID string `json:"user_id"`
	// 排班绑定的假勤审批类型： 1：加班 2：出差 3：请假
	ApproveBizType int `json:"approve_biz_type"`
	// 排班绑定的审批单ID
	ApproveId int `json:"approve_id"`
	// 排班关联的班次设置ID
	ClassSettingID int `json:"class_setting_id"`
	// 排班绑定的假勤审批单名称
	ApproveTagName string `json:"approve_tag_name"`
	// 扩展字段
	Features string `json:"features"`
	// 班次ID
	ClassID int `json:"class_id"`
	// 打卡状态： Init：未打 Checked：已打卡 Timeout：缺卡
	CheckStatus string `json:"check_status"`
	// 工作日，代表具体哪一天的排班
	WorkDate string `json:"work_date"`
	// 结束打卡时间
	CheckEndTime string `json:"check_end_time"`
	// 是否休息： Y：当天排休 N：不休息
	IsRest string `json:"is_rest"`
	// 开始打卡时间
	CheckBeginTime string `json:"check_begin_time"`
	// 开启弹性工时卡点调整后用户应打卡时间
	RealPlanTime string `json:"real_plan_time"`
	// 排班ID
	ID int `json:"id"`
}
