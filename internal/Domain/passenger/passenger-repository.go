package passenger

type Repository interface {
	GetAll() ([]Passenger, error)
}
