package request

type UpdateAttendanceVacationQuota struct {
	// 当前企业内拥有OA审批应用权限的管理员的userId
	OpUserid string `json:"op_userid"`
	// 假期额度信息
	LeaveQuotas []LeaveQuotas `json:"leave_quotas"`
}
