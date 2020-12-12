package functions

import (
	"testing"
)

func Test_GetListOfStrings_Returns_OK(t *testing.T) {
	if list := GetRandomStringList(); len(list) == 0 {
		t.Fail()
	}
}
