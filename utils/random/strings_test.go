package random

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Strings(6))
	}
}
