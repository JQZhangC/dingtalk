package request

type DeleteAttendanceGroup struct {
	// 操作人userId
	UserId string `json:"op_user_id"`
	// 考勤组ID
	GroupKey string `json:"group_key"`
}

func NewDeleteAttendanceGroup(userId, groupKey string) *DeleteAttendanceGroup {
	return &DeleteAttendanceGroup{
		UserId:   userId,
		GroupKey: groupKey,
	}
}
