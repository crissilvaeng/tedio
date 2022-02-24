package storage

import "fmt"

type GameNotFoundErr struct {
	ID string
}

func (e *GameNotFoundErr) Error() string {
	return fmt.Sprintf("Game with ID %s not found", e.ID)
}
