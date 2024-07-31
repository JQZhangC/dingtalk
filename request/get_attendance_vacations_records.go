package request

type GetAttendanceVacationsRecords struct {
	//当前企业内拥有OA审批应用权限的管理员的userId，建议传入企业主管理员userId。
	OpUserId string `json:"opUserId"`
	//假期类型唯一标识
	LeaveCode string `json:"leaveCode"`
	//待查询员工userId列表，每次调用最多传50个userId
	UserIds []string `json:"userIds"`
	//当前页码，从0开始
	PageNumber int64 `json:"pageNumber"`
	//每页条目数，最大值200
	PageSize int `json:"pageSize"`
}
