package response

type CreateAttendanceShift struct {
	Response

	AttendanceShift struct {
		// 班次id
		Id int `json:"id"`
		// 班次名称
		Name string `json:"name"`
	} `json:"result"`
}
