package num7

import (
	"fmt"
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println(math.MaxInt32)
	ret := reverse(2147483647)
	fmt.Println(ret)
}
