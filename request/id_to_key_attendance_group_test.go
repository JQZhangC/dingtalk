package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdToKeyAttendanceGroup(t *testing.T) {
	idToKey := NewIdToKeyAttendanceGroup("manager164", 18357154439)
	assert.NotNil(t, idToKey)
	assert.Equal(t, idToKey.UserId, "manager164")
	assert.Equal(t, idToKey.GroupId, int64(18357154439))
}
