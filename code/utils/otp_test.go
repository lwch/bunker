package utils

import (
	"testing"
)

func TestTOTP(t *testing.T) {
	otp := NewTOTP("0123456789")
	h := otp.Gen()
	if !otp.Verify(h) {
		t.Fatal("invalid otp")
	}
}
