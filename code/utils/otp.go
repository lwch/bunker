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

func (otp *TOTP) hash(t time.Time) uint64 {
	ts := t.Unix() / 30
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(ts))
	otp.Lock()
	otp.h.Reset()
	data = otp.h.Sum(data)
	otp.Unlock()
	offset := data[31] & 23
	return binary.BigEndian.Uint64(data[offset:])
}

// Gen generate otp
func (otp *TOTP) Gen() string {
	return fmt.Sprintf("%d", otp.hash(time.Now()))
}

// GenFrom generate otp from time
func (otp *TOTP) GenFrom(t time.Time) string {
	return fmt.Sprintf("%d", otp.hash(t))
}

// Verify verify otp value
func (otp *TOTP) Verify(str string) bool {
	return otp.Gen() == str
}

// Verify verify otp value from time
func (otp *TOTP) VerifyFrom(str string, t time.Time) bool {
	return otp.GenFrom(t) == str
}
