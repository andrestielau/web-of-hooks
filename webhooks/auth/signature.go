package auth

import (
	"crypto/hmac"
	"crypto/sha256"
)

type HMAC []byte

func (v HMAC) Sign(data []byte) []byte {
	mac := hmac.New(sha256.New, v)
	mac.Write(data)
	return mac.Sum(nil)
}

func (v HMAC) Validate(toSign, expected []byte) bool {
	return hmac.Equal(v.Sign(toSign), expected)
}
