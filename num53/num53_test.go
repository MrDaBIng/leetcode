package num53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	assert.Equal(t, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	assert.Equal(t, maxSubArray([]int{-1}), -1)
}
