package functions

import (
	"strings"
	"testing"
)

func Test_GetListOfStrings_Returns_OK(t *testing.T) {
	list := GetRandomStringList()

	if len(list) == 0 {
		t.Fail()
	}

	for _, msg := range list {
		if len(strings.TrimSpace(msg)) == 0 {
			t.Fail()
		}
	}
}
