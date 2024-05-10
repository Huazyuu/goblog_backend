package random

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Code(4))
	}
}
