package request

import "encoding/json"

type DeleteAttendanceShift struct {
	// 操作人userId
	UserId string `json:"op_user_id"`
	// 班次ID
	ShiftId int64 `json:"shift_id"`
}

func NewDeleteAttendanceShift(userId string, shiftId int64) *DeleteAttendanceShift {
	return &DeleteAttendanceShift{
		UserId:  userId,
		ShiftId: shiftId,
	}
}

func (d *DeleteAttendanceShift) String() string {
	str, _ := json.Marshal(d)
	return string(str)
}
