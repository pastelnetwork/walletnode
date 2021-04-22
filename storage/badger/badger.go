package badger

import (
	v3Badger "github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/storage"

	"github.com/pastelnetwork/go-commons/log"
)

// DB structure represents wrapper over the badger.DB
type DB struct {
	db     *v3Badger.DB
	logger *Logger
}

// NewDB creates an instance of BadgerDB and starts fetching data
func NewDB(cfg *Config) storage.KeyValue {
	ch := &DB{
		db:     &v3Badger.DB{},
		logger: NewLogger(),
	}
	if err := ch.start(cfg); err != nil {
		ch.logger.Errorf("error caused when trying to start badger db from %v", cfg.Dir)
		return nil
	}
	return ch
}

func (db *DB) Init() error {
	return nil
}

// Init method opens the v3Badger db instance using specified configuration data and logger instance
func (db *DB) start(cfg *Config) (err error) {
	db.db, err = v3Badger.Open(v3Badger.DefaultOptions(cfg.Dir).WithLogger(log.DefaultLogger))
	if err != nil {
		return errors.Errorf("can't fetch database from %v | %v", cfg.Dir, err)
	}
	return nil
}

// Get method fetches data by key
func (db *DB) Get(key string) (result []byte, err error) {
	err = db.db.View(func(txn *v3Badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return errors.Errorf("can not fetch data by key %s | %v", key, err)
		}
		if result, err = item.ValueCopy(result); err != nil {
			return errors.Errorf("can not fetch data by key %s | %v", key, err)
		}
		return nil
	})
	return
}

// Set method inserts data using key and value
func (db *DB) Set(key string, value []byte) error {
	return db.db.Update(func(txn *v3Badger.Txn) error {
		if err := txn.Set([]byte(key), value); err != nil {
			return errors.Errorf("can not set data key %s value %s | %v", key, value, err)
		}
		return nil
	})
}

// Delete method deletes data using key
func (db *DB) Delete(key string) (err error) {
	return db.db.Update(
		func(txn *v3Badger.Txn) error {
			if err := txn.Delete([]byte(key)); err != nil {
				return errors.Errorf("can not delete data by key %s | %v", key, err)
			}
			return nil
		})
}

// Close method closes v3Badger db database
func (db *DB) Close() error {
	if err := db.db.Close(); err != nil {
		return errors.Errorf("can't close database | %v", err)
	}
	return nil
}
