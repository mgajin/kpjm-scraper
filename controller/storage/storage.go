package storage

// StoringService interface is used in every storage implementation.
type StoringService interface {
	Store(interface{}) error
}
