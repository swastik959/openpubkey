package gq

import (
	"crypto"
	"crypto/sha256"
)

// Hardcoded padding prefix for SHA-256 from https://github.com/golang/go/blob/eca5a97340e6b475268a522012f30e8e25bb8b8f/src/crypto/rsa/pkcs1v15.go#L268
var prefix = []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}

// encodePKCS1v15 is taken from the go stdlib, see [crypto/rsa.SignPKCS1v15].
//
// https://github.com/golang/go/blob/eca5a97340e6b475268a522012f30e8e25bb8b8f/src/crypto/rsa/pkcs1v15.go#L287-L317
func encodePKCS1v15(k int, data []byte) []byte {
	hashLen := crypto.SHA256.Size()
	tLen := len(prefix) + hashLen

	// EM = 0x00 || 0x01 || PS || 0x00 || T
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k-hashLen], prefix)

	hashed := sha256.Sum256(data)
	copy(em[k-hashLen:k], hashed[:])
	return em
}
