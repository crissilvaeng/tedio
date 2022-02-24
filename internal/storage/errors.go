package storage

import "fmt"

type GameNotFoundErr struct {
	ID string
}

func (e *GameNotFoundErr) Error() string {
	return fmt.Sprintf("Game with ID %s not found", e.ID)
}

type InviteCodeNotFoundErr struct {
	ID string
}

func (e *InviteCodeNotFoundErr) Error() string {
	return fmt.Sprintf("Invite code with ID %s not found", e.ID)
}

type UsernameAlreadyInUseErr struct {
	Username string
}

func (e *UsernameAlreadyInUseErr) Error() string {
	return fmt.Sprintf("Username %s already in use", e.Username)
}

type PlayerNotFoundErr struct {
	Username string
}

func (e *PlayerNotFoundErr) Error() string {
	return fmt.Sprintf("Player with username %s not found", e.Username)
}
