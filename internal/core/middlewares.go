package core

import "net/http"

func (s *Server) Secure(next http.HandlerFunc) http.HandlerFunc {
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
