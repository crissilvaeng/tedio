package core

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
)

type Auth struct {
	apikey string
	secret string
}

func (a *Auth) compare(left, right string) bool {
	hashleft := sha256.Sum256([]byte(left))
	hashright := sha256.Sum256([]byte(right))
	return subtle.ConstantTimeCompare(hashleft[:], hashright[:]) == 1
}

func (a *Auth) authenticate(apikey, apisecret string) error {
	if len(apikey) == 0 || len(apisecret) == 0 {
		return fmt.Errorf("invalid arguments")
	}
	if !a.compare(a.apikey, apikey) || !a.compare(a.secret, apisecret) {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
