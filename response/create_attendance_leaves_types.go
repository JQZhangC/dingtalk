package response

import "github.com/zhaoyunxing92/dingtalk/v2/request"

type CreateAttendanceLeavesTypes struct {
	Response

	Result CreateAttendanceLeavesType `json:"result"`
}

type CreateAttendanceLeavesType struct {
	//假期名称
	LeaveName string `json:"leaveName"`
	//假期规则唯一标识
	LeaveCode string `json:"leaveCode"`
	//请假单位
	//day：天
	//halfDay：半天
	//hour：小时
	LeaveViewUnit string `json:"leaveViewUnit"`
	//假期类型。general_leave：普通假期;lieu_leave：加班转调休
	BizType string `json:"bizType"`
	// 是否按照自然日统计请假时长 true：是 false：否
	NaturalDayLeave bool `json:"naturalDayLeave"`
	// 每天折算的工作时长，为参数值的百分之一
	HoursInPerDay int64 `json:"hoursInPerDay"`
	// 适用范围规则列表，不传默认为全公司
	VisibilityRules []request.AttendanceLeavesTypesRules `json:"visibilityRules,omitempty"`

	// 限时提交规则
	SubmitTimeRule request.AttendanceLeavesTypesSubmitTimeRule `json:"submitTimeRule"`

	// 请假证明
	LeaveCertificate request.AttendanceLeavesTypesLeaveCertificate `json:"leaveCertificate"`
}
