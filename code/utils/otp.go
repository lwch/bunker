package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"sync"
	"time"
)

// TOTP time based one time password
type TOTP struct {
	sync.Mutex
	h hash.Hash
}

// NewTOTP create time based one time password hasher
func NewTOTP(secret string) *TOTP {
	return &TOTP{h: hmac.New(sha256.New, []byte(secret))}
}

// Gen generate otp
func (otp *TOTP) Gen() string {
	return otp.GenFrom(time.Now())
}

// GenFrom generate otp from time
func (otp *TOTP) GenFrom(t time.Time) string {
	ts := t.Unix() / 30
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(ts))
	otp.Lock()
	otp.h.Reset()
	otp.h.Write(data)
	data = otp.h.Sum(nil)
	otp.Unlock()
	return fmt.Sprintf("%x", data)
}

// Verify verify otp value
func (otp *TOTP) Verify(str string) bool {
	return otp.Gen() == str
}

// Verify verify otp value from time
func (otp *TOTP) VerifyFrom(str string, t time.Time) bool {
	return otp.GenFrom(t) == str
}
