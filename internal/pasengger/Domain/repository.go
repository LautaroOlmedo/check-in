package Domain

import (
	"errors"
)

var (
	ErrPassengerNotFound    = errors.New("cannot found the passenger")
	ErrFailedToAddPassenger = errors.New("cannot created the passenger")
	ErrUpdatePassenger      = errors.New("cannot update the passenger")
)

type PassengerRepository interface {
}
