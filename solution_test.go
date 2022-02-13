package solution

import "testing"

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	expected := "Hello \U0001f5fa\ufe0f !"
	if result != expected {
		t.Error(result, "!=", expected)
	}
}
