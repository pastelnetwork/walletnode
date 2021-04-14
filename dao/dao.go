package dao

type DB interface {
	// Init initializes the database
	Init() error
}
