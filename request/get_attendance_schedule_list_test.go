package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAttendanceScheduleList(t *testing.T) {
	list := NewGetAttendanceScheduleList("manager164", "manager164", 1200)
	assert.NotNil(t, list)
	assert.Equal(t, list.OpUserId, "manager164")
	assert.Equal(t, list.UserId, "manager164")
	assert.Equal(t, list.DateTime, int64(1200))
}
