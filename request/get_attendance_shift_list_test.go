package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAttendanceShiftList(t *testing.T) {
	list := NewGetAttendanceShiftList("manager164", 1)
	assert.NotNil(t, list)
	assert.Equal(t, list.UserId, "manager164")
	assert.Equal(t, list.Cursor, int64(1))
}
