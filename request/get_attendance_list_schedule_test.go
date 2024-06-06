package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAttendanceListSchedule(t *testing.T) {
	schedule := NewGetAttendanceListSchedule("2024-06-04", 1, 10)

	assert.Equal(t, schedule.WorkDate, "2024-06-04")
	assert.Equal(t, schedule.Offset, 1)
	assert.Equal(t, schedule.Size, 10)
}
