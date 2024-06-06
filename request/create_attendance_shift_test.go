package request

import (
	"testing"
)

func TestNewCreateAttendanceShift(t *testing.T) {
	shift := NewCreateAttendanceShift("manager164",
		"创建测试班次", []AttendanceShiftSection{
			AttendanceShiftSection{
				Times: []AttendanceShiftTime{
					AttendanceShiftTime{
						FreeCheck: false,
						BeginMin:  30,
						CheckType: "OnDuty",
						Across:    0,
						EndMin:    -1,
						CheckTime: "2024-06-06 09:00:00",
					},
				},
			},
		}).
		Build()

	t.Log(shift.String())
}
