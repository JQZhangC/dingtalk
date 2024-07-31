package request

type AttendanceLeavesTypesRules struct {
	// 适用范围内数据列表
	Visible []string `json:"visible,omitempty"`
	// 适用范围规则类型
	Type string `json:"type"`
}

type AttendanceLeavesTypesSubmitTimeRule struct {
	// 限制值。
	// 当timeUnit为day时，有效值范围是0至30天。
	// 当timeUnit为hour时，有效值范围是0至24小时
	TimeValue int64 `json:"timeValue"`

	// 时间单位。
	// day：天
	// hour：小时
	TimeUnit string `json:"timeUnit"`

	// 限制类型
	// before：提前
	// after：补交
	TimeType string `json:"timeType"`

	// 是否开启限时提交功能
	// true：开启
	// false：不开启
	EnableTimeLimit bool `json:"enableTimeLimit"`
}

type AttendanceLeavesTypesLeaveCertificate struct {
	// 需提供请假证明时长单位。hour：小时, day：天
	Unit string `json:"unit"`
	// 超过多长时间需提供请假证明
	// 如果unit值为day，表示请假超过一天，需要提供请假证明。
	// 如果unit值为hour，表示请假超过一小时，需要提供请假证明。
	Duration float64 `json:"duration"`
	// 是否开启请假证明, true：开启, false：不开启
	Enable bool `json:"enable"`
	// 请假提示文案
	PromptInformation string `json:"promptInformation"`
}

type CreateAttendanceLeavesTypes struct {
	// 假期规则名称
	LeaveName string `json:"leaveName" validate:"required"`

	// 请假时长单位。 day：天 halfDay：半天 hour：小时
	LeaveViewUnit string `json:"leaveViewUnit" validate:"required"`

	// 假期类型 general_leave：普通假期 lieu_leave：加班转调休
	// 一个企业只能存在一个加班转调休的假期规则。
	BizType string `json:"bizType" validate:"required"`

	// 是否按照自然日统计请假时长
	// true：是 false：否
	NaturalDayLeave bool `json:"naturalDayLeave"`

	// 每天折算的工作时长，为参数值的百分之一
	HoursInPerDay int64 `json:"hoursInPerDay" validate:"required"`

	// 调休假有效期规则
	Extras string `json:"extras,omitempty"`

	// 适用范围规则列表，不传默认为全公司
	VisibilityRules []AttendanceLeavesTypesRules `json:"visibilityRules,omitempty"`

	// 限时提交规则
	SubmitTimeRule *AttendanceLeavesTypesSubmitTimeRule `json:"submitTimeRule,omitempty"`

	// 请假证明
	LeaveCertificate *AttendanceLeavesTypesLeaveCertificate `json:"leaveCertificate,omitempty"`
}
