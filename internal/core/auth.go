package core

import (
	"fmt"

	"github.com/crissilvaeng/tedio/internal/misc"
)

type Auth struct {
	apikey string
	secret string
}

func (a *Auth) authenticate(apikey, apisecret string) error {
	if len(apikey) == 0 || len(apisecret) == 0 {
		return fmt.Errorf("invalid arguments")
	}
	if !misc.Compare(a.apikey, apikey) || !misc.Compare(a.secret, apisecret) {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
