package request

type GetAttendanceVacationQuotaList struct {
	// 假期类型唯一标识
	LeaveCode string `json:"leave_code" validate:"required"`
	// 当前企业内拥有OA审批应用权限的管理员的userId
	OpUserId string `json:"op_userid"`
	//待查询的员工ID列表,英文逗号分割
	UserIds string `json:"userids"`
	//分页偏移，从0开始的非负整数
	Offset int `json:"offset"`
	//分页偏移，最大50
	Size int `json:"size"`
}
