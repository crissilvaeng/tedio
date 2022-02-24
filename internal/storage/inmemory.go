package storage

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/crissilvaeng/tedio/internal/misc"
	"github.com/crissilvaeng/tedio/internal/models"
	"github.com/google/uuid"
)

type InMemoryStorage struct {
	mutex     *sync.RWMutex
	games     []*models.Game
	gamesByID map[string]*models.Game
	invites   map[string]*models.InviteCode
	users     map[string]*models.Player
	logger    *log.Logger
}

func NewInMemoryStorage() GameRepository {
	return &InMemoryStorage{
		mutex:     &sync.RWMutex{},
		games:     make([]*models.Game, 0),
		gamesByID: make(map[string]*models.Game),
		invites:   make(map[string]*models.InviteCode),
		users:     make(map[string]*models.Player),
		logger:    log.New(os.Stdout, "storage: ", log.LstdFlags),
	}
}

func (s *InMemoryStorage) CreateGame(req *models.Game) (*models.Game, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	game := &models.Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		// Rounds:       make([]models.Round, 0),
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

func (s *InMemoryStorage) GetInviteCode(gameID string) (*models.InviteCode, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	game := s.gamesByID[gameID]
	if game == nil {
		return nil, &GameNotFoundErr{ID: gameID}
	}

	invite := &models.InviteCode{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		GameID:    game.ID,
	}

	s.invites[invite.ID] = invite
	s.logger.Printf("created invite %s: %+v", invite.ID, invite)
	return invite, nil
}

func (s *InMemoryStorage) RedeemInviteCode(code string, cred models.Credentials) (*models.Player, error) {
	s.logger.Printf("redeeming invite code %s: credentials %+v", code, cred)
	if code == "" {
		return nil, &InviteCodeNotFoundErr{}
	}

	s.mutex.RLock()
	invite, ok := s.invites[code]
	if !ok {
		s.mutex.RUnlock()
		return nil, &InviteCodeNotFoundErr{ID: code}
	}
	s.mutex.RUnlock()

	s.mutex.RLock()
	if user := s.users[cred.Username]; user != nil {
		s.mutex.RUnlock()
		return nil, &UsernameAlreadyInUseErr{Username: cred.Username}
	}
	s.mutex.RUnlock()

	salt := misc.GenerateSalt()
	user := &models.Player{
		Username:     cred.Username,
		Salt:         salt,
		HashPassword: misc.HashPassword(cred.Password, salt),
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.invites, code)

	s.users[user.Username] = user
	s.gamesByID[invite.GameID].Players = append(s.gamesByID[invite.GameID].Players, *user)

	s.logger.Printf("created user %s: %+v", user.Username, user)
	return user, nil
}
