package request

type GetAttendanceScheduleUsersList struct {
	// 操作人userId
	OpUserId string `json:"op_user_id"`
	// 要查询的人员userId列表，多个userId用逗号分隔，一次最多可传50个
	UserIds string `json:"userids"`
	// 起始日期，Unix时间戳，单位毫秒
	FromDateTime int64 `json:"from_date_time"`
	// 结束日期，Unix时间戳，单位毫秒
	ToDateTime int64 `json:"to_date_time"`
}

func NewGetAttendanceScheduleUsersList(opUserId, userIds string, fromDateTime, toDateTime int64) *GetAttendanceScheduleUsersList {
	return &GetAttendanceScheduleUsersList{
		OpUserId:     opUserId,
		UserIds:      userIds,
		FromDateTime: fromDateTime,
		ToDateTime:   toDateTime,
	}
}
