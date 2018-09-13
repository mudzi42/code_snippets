package codesnippets

import "encoding/hex"

func convertIPtoHex(ip string) string {
	src := []byte(ip)

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	hexString := string(dst[:])
	return hexString
}
