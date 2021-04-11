package store

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/log"
)

type ChatDB struct {
	db *badger.DB
}

// NewChatDB creates an instance of ChatDB
func NewChatDB() *ChatDB {
	return &ChatDB{db: &badger.DB{}}
}

// Start method opens the badger db instance using specified configuration data and logger instance
func (db *ChatDB) Start(cfg *Config, logger *log.Logger) (err error) {
	db.db, err = badger.Open(badger.DefaultOptions(cfg.path).WithLogger(logger))
	return err
}

// Get method fetches data by key
func (db *ChatDB) Get(key []byte) (result []byte, err error) {
	err = db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		result, err = item.ValueCopy(result)
		return err
	})
	return
}

// Set method inserts data using key and value
func (db *ChatDB) Set(key, value []byte) (err error) {
	err = db.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return
}

// Delete method deletes data using key
func (db *ChatDB) Delete(key []byte) (err error) {
	err = db.db.Update(
		func(txn *badger.Txn) error {
			err := txn.Delete(key)
			if err != nil {
				return err
			}
			return nil
		})
	return
}

// Close method closes badger db database
func (db *ChatDB) Close() error {
	return db.db.Close()
}
