package request

type GetAttendanceShiftDetail struct {
	// 操作人userId
	UserId string `json:"op_user_id" validate:"required"`
	// 班次ID
	ShiftId int64 `json:"shift_id" validate:"required"`
}

func NewGetAttendanceShiftDetail(userId string, shiftId int64) *GetAttendanceShiftDetail {
	return &GetAttendanceShiftDetail{
		UserId:  userId,
		ShiftId: shiftId,
	}
}
