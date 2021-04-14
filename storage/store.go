package storage

// KVStorage represents the basic interface of each database
// model that should be key-value based
type KVStorage interface {
	Start(*Config) error
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) error
}
