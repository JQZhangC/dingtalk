package response

type GetAttendanceShiftList struct {
	Response
	Result AttendanceShiftList `json:"result"`
}

type AttendanceShiftList struct {
	// 是否还有下一页
	HasMore bool `json:"has_more"`
	// 下一页的游标位置
	Cursor int64 `json:"cursor"`
	// 班次信息
	Result []AttendanceShiftListResult `json:"result"`
}

type AttendanceShiftListResult struct {
	// 班次名称
	Name string `json:"name"`
	// 班次ID
	Id int64 `json:"id"`
}
