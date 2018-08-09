package num27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	assert.Equal(t, removeElement([]int{3, 2, 2, 3}, 3), 2)
	assert.Equal(t, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), 5)
	assert.Equal(t, removeElement([]int{1, 2, 3, 4, 5, 6}, 7), 6)
	assert.Equal(t, removeElement([]int{3, 3, 3, 3, 3, 3}, 3), 0)
}

func Test_removeElement1(t *testing.T) {
	assert.Equal(t, removeElement1([]int{3, 2, 2, 3}, 3), 2)
	assert.Equal(t, removeElement1([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), 5)
	assert.Equal(t, removeElement1([]int{1, 2, 3, 4, 5, 6}, 7), 6)
	assert.Equal(t, removeElement1([]int{3, 3, 3, 3, 3, 3}, 3), 0)
}

func Test_removeElement2(t *testing.T) {
	assert.Equal(t, removeElement2([]int{3, 2, 2, 3}, 3), 2)
	assert.Equal(t, removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), 5)
	assert.Equal(t, removeElement2([]int{1, 2, 3, 4, 5, 6}, 7), 6)
	assert.Equal(t, removeElement2([]int{3, 3, 3, 3, 3, 3}, 3), 0)
}
