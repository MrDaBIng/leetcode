package num67

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, addBinary("11", "1"), "100")
	assert.Equal(t, addBinary("1010", "1011"), "10101")
	assert.Equal(t, addBinary("101111", "10"), "110001")
}
