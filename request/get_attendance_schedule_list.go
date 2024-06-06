package request

type GetAttendanceScheduleList struct {
	// 操作人userId
	OpUserId string `json:"op_user_id"`
	// 要查询的人员userId
	UserId string `json:"user_id"`
	// 查询的时间，Unix时间戳，单位毫秒
	DateTime int64 `json:"date_time"`
}

func NewGetAttendanceScheduleList(opUserId, userId string, dateTime int64) *GetAttendanceScheduleList {
	return &GetAttendanceScheduleList{
		OpUserId: opUserId,
		UserId:   userId,
		DateTime: dateTime,
	}
}
