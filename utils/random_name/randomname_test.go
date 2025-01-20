package randomname

import "testing"

func TestGenerateName(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GenerateName())
	}
}
