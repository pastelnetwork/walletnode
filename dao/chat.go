package dao

import (
	"github.com/pastelnetwork/walletnode/storage"
	"github.com/pastelnetwork/walletnode/storage/badger"
)

// ChatDB is a wrapper over badgerDB where should be stored all chat metadata
type ChatDB struct {
	db     storage.KeyValue
	prefix string
}

const (
	keyPrefix = "chat/"
)

func NewChatDB(cfg *badger.Config) storage.KeyValue {
	return &ChatDB{
		db:     badger.NewBadgerDB(cfg),
		prefix: keyPrefix,
	}
}

// Init initializes the database.
func (chatDB *ChatDB) Init() error {
	return chatDB.db.Init()
}

// Get looks for key and returns corresponding Item.
// If key is not found, ErrKeyNotFound is returned.
func (chatDB *ChatDB) Get(key string) (value []byte, err error) {
	return chatDB.db.Get(chatDB.prefix + key)
}

// Set adds a key-value pair to the database.
func (chatDB *ChatDB) Set(key string, value []byte) (err error) {
	return chatDB.db.Set(chatDB.prefix+key, value)
}

// Delete deletes a key.
func (chatDB *ChatDB) Delete(key string) error {
	return chatDB.db.Delete(chatDB.prefix + key)
}
