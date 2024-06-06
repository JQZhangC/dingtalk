package request

type IdToKeyAttendanceGroup struct {
	// 操作人userId
	UserId string `json:"op_user_id"  validate:"required"`
	// 考勤组ID
	GroupId int64 `json:"group_id"  validate:"required"`
}

func NewIdToKeyAttendanceGroup(userId string, groupId int64) *IdToKeyAttendanceGroup {
	return &IdToKeyAttendanceGroup{
		UserId:  userId,
		GroupId: groupId,
	}
}
