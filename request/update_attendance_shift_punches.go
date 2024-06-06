package request

type UpdateAttendanceShiftPunches struct {
	// 操作者userId
	UserId string `json:"op_user_id"  validate:"required"`
	// 卡点信息
	Punches []AttendanceShiftPunches `json:"punches"`
	// 班次ID
	ShiftId int64 `json:"shift_id"  validate:"required"`
}

type AttendanceShiftPunches struct {
	// 卡点ID
	Id int64 `json:"id"  validate:"required"`
	// 是否无需打卡。true：开启无需打卡。 false：关闭无需打卡
	FreeCheck bool `json:"free_check"  validate:"required"`
}

func NewUpdateAttendanceShiftPunches(userId string, shiftId int64, punches []AttendanceShiftPunches) *UpdateAttendanceShiftPunches {
	return &UpdateAttendanceShiftPunches{
		UserId:  userId,
		ShiftId: shiftId,
		Punches: punches,
	}
}
