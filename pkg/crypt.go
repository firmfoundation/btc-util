package crypt

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func RandByte() string {
	key := make([]byte, 8)

	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}

	str := hex.EncodeToString(key)

	return str
}
