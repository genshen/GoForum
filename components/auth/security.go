package auth

import (
	"crypto/sha1"
	"io"
	"encoding/hex"
)

func Hash(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	return hex.EncodeToString(h.Sum(nil))
}
