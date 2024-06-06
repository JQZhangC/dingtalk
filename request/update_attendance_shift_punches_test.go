package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUpdateAttendanceShiftPunches(t *testing.T) {
	punches := []AttendanceShiftPunches{
		{
			Id:        937375145,
			FreeCheck: false,
		},
		{
			Id:        937375146,
			FreeCheck: false,
		},
	}

	shiftPunches := NewUpdateAttendanceShiftPunches("userId", 1, punches)

	assert.Equal(t, shiftPunches.UserId, "userId")
	assert.Equal(t, shiftPunches.ShiftId, int64(1))
	assert.Equal(t, shiftPunches.Punches[0].Id, int64(937375145))
	assert.Equal(t, shiftPunches.Punches[1].Id, int64(937375146))
}
