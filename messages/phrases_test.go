package messages

import (
	"testing"
)

func Test_GetRandomPhrase_Returns_OK(t *testing.T) {
	if _, err := GetRandomPhrase(); err != nil {
		t.Fail()
	}
}
