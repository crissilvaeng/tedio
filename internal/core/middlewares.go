package core

import (
	"net/http"

	"github.com/crissilvaeng/tedio/internal/misc"
)

func (s *Server) Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, secret, ok := r.BasicAuth()
		if ok {
			if s.Auth.authenticate(apikey, secret) == nil {
				next(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func (s *Server) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			player, err := s.storage.GetPlayerByUsername(username)
			if err == nil {
				hashed := misc.HashPassword(password, player.Salt)
				if misc.Compare(player.HashPassword, hashed) {
					next(w, r)
					return
				}
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
