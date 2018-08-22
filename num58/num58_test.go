package num58

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	assert.Equal(t, lengthOfLastWord("Hello World"), 5)
	assert.Equal(t, lengthOfLastWord("HelloWorld"), 10)
	assert.Equal(t, lengthOfLastWord("HelloWorld "), 10)
	assert.Equal(t, lengthOfLastWord(" HelloWorld "), 10)
}

func Test_maxSubArray2(t *testing.T) {
	assert.Equal(t, lengthOfLastWord2("Hello World"), 5)
	assert.Equal(t, lengthOfLastWord2("HelloWorld"), 10)
	assert.Equal(t, lengthOfLastWord2("HelloWorld "), 10)
	assert.Equal(t, lengthOfLastWord2(" HelloWorld "), 10)
}
