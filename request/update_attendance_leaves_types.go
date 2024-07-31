package request

type UpdateAttendanceLeavesTypes struct {
	// 假期规则名称
	LeaveName string `json:"leaveName"`

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
	// 接口添加的假期规则标识，leave_code必须是通过接口添加的假期类型。
	LeaveCode string `json:"leaveCode" validate:"required"`
}
