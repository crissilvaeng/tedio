package misc

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"

	"github.com/google/uuid"
)

func GetOrElseStr(value, fallback string) string {
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetOrElseInt(value, fallback int) int {
	if value == 0 {
		return fallback
	}
	return value
}

func GetMinValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GenerateSalt() string {
	return uuid.New().String()
}

func HashPassword(password, salt string) string {
	hash := sha256.Sum256([]byte(password + salt))
	return fmt.Sprintf("%x", hash)
}

func Compare(left, right string) bool {
	hashleft := sha256.Sum256([]byte(left))
	hashright := sha256.Sum256([]byte(right))
	return subtle.ConstantTimeCompare(hashleft[:], hashright[:]) == 1
}
