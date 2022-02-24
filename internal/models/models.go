package models

import "time"

type Round struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Word      string    `json:"word"`
}

type Score struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Score     int       `json:"score"`
}

type Player struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Scores   []Score `json:"scores"`
}

type Game struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Rounds       []Round   `json:"rounds"`
	Players      []Player  `json:"players"`
	WordLenght   int       `json:"word_lenght"`
	TotalGuesses int       `json:"total_guesses"`
}
