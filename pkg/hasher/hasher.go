package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type hasher struct {
	salt []byte
}

func NewHasher(salt []byte) *hasher {
	return &hasher{
		salt: salt,
	}
}

func (h hasher) GetHash(msg string) (string, error) {
	md := md5.New()
	_, err := md.Write([]byte(msg))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(md.Sum(h.salt)), nil
}

// only for testing
type badHasher struct {
}

func NewBadHasher() *badHasher {
	return &badHasher{}
}

func (b badHasher) GetHash(msg string) (string, error) {
	return "", errors.New("some error")
}
