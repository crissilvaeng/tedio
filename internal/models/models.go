package models

import "time"

type Player struct {
	Username     string `json:"username"`
	Salt         string `json:"-"`
	HashPassword string `json:"-"`
}

type Game struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Players      []Player  `json:"players"`
	WordLenght   int       `json:"word_lenght"`
	TotalGuesses int       `json:"total_guesses"`
}

type InviteCode struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	GameID    string    `json:"game_id"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
