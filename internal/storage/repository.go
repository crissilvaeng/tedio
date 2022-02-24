package storage

import "github.com/crissilvaeng/tedio/internal/models"

const (
	DefaultTotalGuesses = 6
	DefaultWordLenght   = 5
)

type GameRepository interface {
	CreateGame(game *models.Game) (*models.Game, error)
	GetGame(id string) (*models.Game, error)
	GetGames(limit, offset int) ([]*models.Game, error)
	GetInviteCode(gameID string) (*models.InviteCode, error)
}
