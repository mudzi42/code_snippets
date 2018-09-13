package codesnippets

import (
	"crypto/sha1"
	"fmt"
)

func generateSHA1(pw string) string {
	h := sha1.New()
	h.Write([]byte(pw))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
