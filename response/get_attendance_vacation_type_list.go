package response

type GetAttendanceVacationTypeList struct {
	Response
	Result []GetAttendanceVacationType `json:"result"`
}

type GetAttendanceVacationType struct {
	//假期规则唯一标识
	LeaveCode string `json:"leave_code"`
	//假期名称
	LeaveName string `json:"leave_name"`
	//请假单位。
	//day：天
	//halfDay：半天
	//hour：小时
	LeaveViewUnit string `json:"leave_view_unit"`
	//请假证明
	LeaveCertificate *LeaveCertificateVo `json:"leave_certificate"`
	//限时提交规则
	SubmitTimeRule *SubmitTimeRuleVo `json:"submit_time_rule"`
	//假期类型。
	//general_leave：普通假期
	//lieu_leave：加班转调休
	BizType string `json:"biz_type"`
	//是否按照自然日统计请假时长。
	//true：按照自然日统计请假时长
	//false：不按照自然日统计请假时长
	NaturalDayLeave string `json:"natural_day_leave"`
	//有效类型。
	//absolute_time：绝对时间
	//relative_time：相对时间
	ValidityType string `json:"validity_type"`
	//延长日期。
	//当validity_type为absolute_time该值不为空且满足“yy-mm”格式。
	//当validity_type为relative_time该值为大于1的整数。
	ValidityValue string `json:"validity_value"`
	//每天折算的工作时长，百分之一。
	//例如：1天=10小时=1000
	HoursInPerDay int64 `json:"hours_in_per_day"`
	// 假期来源。
	//external：开放接口自定义的
	//inner：oa后台新建的
	Source string `json:"source"`
}

type LeaveCertificateVo struct {
	//请假证明单位。
	//hour：小时
	//day：天
	Unit string `json:"unit"`
	//超过多长时间需提供请假证明
	Duration int64 `json:"duration"`
	//是否开启请假证明。
	//true：开启
	//false：未开启
	Enable bool `json:"enable"`
	//请假提示文案
	PromptInformation string `json:"prompt_information"`
}

type SubmitTimeRuleVo struct {
	//限制值。
	//当timeUnit为day时，有效值范围是0至30天；
	//timeUnit为hour时，有效值范围是0至24小时。
	TimeValue int `json:"time_value"`
	//时间单位
	//day：天
	//hour：小时
	TimeUnit string `json:"time_unit"`
	//限制类型。
	//before：提前
	//after：补交
	TimeType string `json:"time_type"`
	//是否开启限时提交功能。
	//true：开启
	//false：不开启
	EnableTimeLimit bool `json:"enable_time_limit"`
}
