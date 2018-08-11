package num28

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	assert.Equal(t, strStr("hello", "ll"), 2)
	assert.Equal(t, strStr("aaaaa", "bba"), -1)
	assert.Equal(t, strStr("a", "a"), 0)
	assert.Equal(t, strStr("mississippi", "issip"), 4)
}

func Test_strStr2(t *testing.T) {
	assert.Equal(t, strStr2("hello", "ll"), 2)
	assert.Equal(t, strStr2("aaaaa", "bba"), -1)
	assert.Equal(t, strStr2("a", "a"), 0)
	assert.Equal(t, strStr2("mississippi", "issip"), 4)
}
