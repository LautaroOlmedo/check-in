package passenger

import (
	"checkin/entity"
	"github.com/google/uuid"
)

type Passenger struct {
	person *entity.Person
}

func NewPassenger(name string) (Passenger, error) {

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Passenger{
		person: person,
	}, nil

}
