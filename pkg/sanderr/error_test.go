package sanderr

import (
	"testing"
)

func TestErr(test *testing.T) {
	err := &Err{
		Message: "test message",
		Details: nil,
	}
	errmsg := err.Error()
	test.Log(errmsg)
}
