package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAttendanceScheduleUsersList(t *testing.T) {
	list := NewGetAttendanceScheduleUsersList("manager164", "manager164,manager165", 1200, 1300)
	assert.NotNil(t, list)
	assert.Equal(t, list.OpUserId, "manager164")
	assert.Equal(t, list.UserIds, "manager164,manager165")
	assert.Equal(t, list.FromDateTime, int64(1200))
	assert.Equal(t, list.ToDateTime, int64(1300))
}
