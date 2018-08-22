package num35

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 5), 2)
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 2), 1)
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 7), 4)
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 0), 0)
}
