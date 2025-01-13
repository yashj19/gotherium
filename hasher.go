package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash(data fmt.Stringer) string {
	hash := sha256.New()
	hash.Write([]byte(data.String()))
	return hex.EncodeToString(hash.Sum(nil))
}

func HashString(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

