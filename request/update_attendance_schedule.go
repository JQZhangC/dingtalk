package request

type UpdateAttendanceSchedule struct {
	// 操作人userId
	OpUserId string `json:"op_user_id"  validate:"required"`
	// 考勤组ID
	GroupId int64 `json:"group_id"  validate:"required"`
	// 排班详情
	Schedules []AttendanceSchedule `json:"schedules" validate:"required"`
}

type AttendanceSchedule struct {
	// 班次ID，休息班次传1
	ShiftId int64 `json:"shift_id" validate:"required"`
	// 排班日期 (可排班日期不早于180天前，不晚于180天后)
	WorkDate int64 `json:"work_date"`
	// 是否休息： true：休息(当该参数为true时，shift_id传1) false：不休息
	IsRest bool `json:"is_rest"`
	// 用户的userId
	UserId string `json:"userid" validate:"required"`
}

func NewUpdateAttendanceSchedule(opUserId string, groupId int64, schedules []AttendanceSchedule) *UpdateAttendanceSchedule {
	return &UpdateAttendanceSchedule{
		OpUserId:  opUserId,
		GroupId:   groupId,
		Schedules: schedules,
	}
}

func NewAttendanceSchedule(shiftId int64, workDate int64, isRest bool, userId string) AttendanceSchedule {
	return AttendanceSchedule{
		ShiftId:  shiftId,
		WorkDate: workDate,
		IsRest:   isRest,
		UserId:   userId,
	}
}
