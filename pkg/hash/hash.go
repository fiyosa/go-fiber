package hash

import (
	"go-fiber/pkg/secret"

	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/bcrypt"
)

func Create(data string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func Verify(check string, original string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(original), []byte(check)); err != nil {
		return false
	}
	return true
}

func Encode(data int) (string, error) {
	h := setupHD()
	encode, err := h.Encode([]int{data})
	if err != nil {
		return "", err
	}
	return encode, err
}

func Decode(data string) (int, error) {
	h := setupHD()
	decode, err := h.DecodeWithError(data)
	if err != nil {
		return -1, err
	}
	return decode[0], err
}

func setupHD() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = secret.APP_SECRET
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	return h
}
