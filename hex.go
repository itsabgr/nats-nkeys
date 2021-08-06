package nkeys

import "encoding/hex"

type hexEncoding struct{}

func (_ hexEncoding) Decode(d, s []byte) (int, error) {
	return hex.Decode(d, s)
}
func (_ hexEncoding) DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func (_ *hexEncoding) DecodedLen(x int) int {
	return hex.DecodedLen(x)
}
func (_ hexEncoding) Encode(d, s []byte) int {
	return hex.Encode(d, s)
}
func (_ hexEncoding) EncodeToString(s []byte) string {
	return hex.EncodeToString(s)
}

func (_ hexEncoding) EncodedLen(x int) int {
	return hex.EncodedLen(x)
}
