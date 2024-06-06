package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAttendanceShiftDetail(t *testing.T) {
	list := NewGetAttendanceShiftDetail("manager164", 1200)
	assert.NotNil(t, list)
	assert.Equal(t, list.UserId, "manager164")
	assert.Equal(t, list.ShiftId, int64(1200))
}
