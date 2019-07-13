package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// Bytes generates n random bytes or returns an error if there was one.
// It uses the crypto/rand package to generate the n random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String generates a byte slize of size nBytes and then returns a string that is the base64 URL encoded version of that byte slice
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Minimum bytes used for remember token, recommended is at least 16
const RememberTokenBytes = 32

func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
