package num66

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3, 4, 5}), []int{1, 2, 3, 4, 6})
	assert.Equal(t, plusOne([]int{1, 2, 3, 4, 9}), []int{1, 2, 3, 5, 0})
	assert.Equal(t, plusOne([]int{9, 9}), []int{1, 0, 0})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
}
