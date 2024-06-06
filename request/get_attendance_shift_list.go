package request

type GetAttendanceShiftList struct {
	// 操作人userId
	UserId string `json:"op_user_id"  validate:"required"`
	// 游标ID，起始值为0
	Cursor int64 `json:"cursor"`
}

func NewGetAttendanceShiftList(userId string, cursor int64) *GetAttendanceShiftList {
	return &GetAttendanceShiftList{
		UserId: userId,
		Cursor: cursor,
	}
}
