package storage

import (
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/crissilvaeng/tedio/internal/misc"
	"github.com/crissilvaeng/tedio/internal/models"
	"github.com/google/uuid"
)

type InMemoryStorage struct {
	mutex     *sync.RWMutex
	gamesByID map[string]*models.Game
	games     []*models.Game
	logger    *log.Logger
}

func NewInMemoryStorage() GameRepository {
	return &InMemoryStorage{
		mutex:     &sync.RWMutex{},
		gamesByID: make(map[string]*models.Game),
		games:     make([]*models.Game, 0),
		logger:    log.New(os.Stdout, "storage: ", log.LstdFlags),
	}
}

func (s *InMemoryStorage) CreateGame(req *models.Game) (*models.Game, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	game := &models.Game{
		ID:           uuid.New().String(),
		CreatedAt:    time.Now(),
		Rounds:       make([]models.Round, 0),
		Players:      make([]models.Player, 0),
		WordLenght:   misc.GetOrElseInt(req.WordLenght, DefaultWordLenght),
		TotalGuesses: misc.GetOrElseInt(req.TotalGuesses, DefaultTotalGuesses),
	}
	s.gamesByID[game.ID] = game
	s.games = append(s.games, game)
	s.logger.Printf("created game %s: %+v", game.ID, game)
	return game, nil
}

func (s *InMemoryStorage) GetGame(id string) (*models.Game, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.logger.Printf("games in storage: %+v", reflect.ValueOf(s.gamesByID).MapKeys())
	game := s.gamesByID[id]
	if game == nil {
		return nil, &GameNotFoundErr{ID: id}
	}
	return game, nil
}

func (s *InMemoryStorage) GetGames(limit, offset int) ([]*models.Game, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	size := len(s.games)
	left := misc.GetMinValue(size, offset)
	right := misc.GetMinValue(size, offset+limit)

	return s.games[left:right], nil
}
