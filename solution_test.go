package solution

import "testing"

func TestHello(t *testing.T) {
	result := hello()
	expected := "Hello \U0001f5fa\ufe0f !"
	if result != expected {
		t.Error(result, "!=", expected)
	}
}
