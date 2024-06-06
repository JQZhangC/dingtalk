package request

import "testing"

func TestDeleteAttendanceShift(t *testing.T) {
	str := NewDeleteAttendanceShift("manager164", 123).String()
	t.Log(str)
}
