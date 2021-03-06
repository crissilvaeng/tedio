package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/crissilvaeng/tedio/internal/misc"
	"github.com/crissilvaeng/tedio/internal/models"
	"github.com/crissilvaeng/tedio/internal/storage"
	"github.com/gorilla/mux"
)

type Routes struct {
	repository storage.GameRepository
}

func (r *Routes) PostGame(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var in models.Game
	if err := decoder.Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := r.repository.CreateGame(&in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(game); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/games/"+game.ID)
}

func (r *Routes) GetGame(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	game, err := r.repository.GetGame(id)
	if err != nil {
		if ok := err.(*storage.GameNotFoundErr); ok != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Routes) GetGames(w http.ResponseWriter, req *http.Request) {
	limit, err := strconv.Atoi(misc.GetOrElseStr(req.URL.Query().Get("limit"), "10"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(misc.GetOrElseStr(req.URL.Query().Get("offset"), "0"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	games, err := r.repository.GetGames(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Routes) GetInviteCode(w http.ResponseWriter, req *http.Request) {
	game := mux.Vars(req)["id"]
	invite, err := r.repository.GetInviteCode(game)
	if err != nil {
		if ok := err.(*storage.GameNotFoundErr); ok != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(invite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Routes) RedeemInviteCode(w http.ResponseWriter, req *http.Request) {
	invite := mux.Vars(req)["invite"]
	decoder := json.NewDecoder(req.Body)
	var cred models.Credentials
	if err := decoder.Decode(&cred); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := r.repository.RedeemInviteCode(invite, cred)
	if err != nil {
		if ok := err.(*storage.InviteCodeNotFoundErr); ok != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if ok := err.(*storage.UsernameAlreadyInUseErr); ok != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(player); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Routes) WhoAmI(w http.ResponseWriter, req *http.Request) {
	username, _, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	player, err := r.repository.GetPlayerByUsername(username)
	if err != nil {
		if ok := err.(*storage.PlayerNotFoundErr); ok != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", player.Username)
}
