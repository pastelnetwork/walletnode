package badger

import (
	v3Badger "github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/storage"
)

// badgerDB structure represents wrapper over the badger.badgerDB
type badgerDB struct {
	db     *v3Badger.DB
	logger *Logger
	config *Config
}

// NewBadgerDB creates an instance of badgerDB
func NewBadgerDB(cfg *Config) storage.KeyValue {
	return &badgerDB{
		db:     &v3Badger.DB{},
		logger: NewLogger(),
		config: cfg,
	}
}

// Init method opens the v3Badger db instance using specified configuration data and logger instance
func (db *badgerDB) Init() (err error) {
	db.db, err = v3Badger.Open(v3Badger.DefaultOptions(db.config.Dir).WithLogger(db.logger))
	if err != nil {
		return errors.Errorf("can't fetch database from %v | %v", db.config.Dir, err)
	}
	return
}

// Get method fetches data by key
func (db *badgerDB) Get(key string) (result []byte, err error) {
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
func (db *badgerDB) Set(key string, value []byte) error {
	return db.db.Update(func(txn *v3Badger.Txn) error {
		if err := txn.Set([]byte(key), value); err != nil {
			return errors.Errorf("can not set data key %s value %s | %v", key, value, err)
		}
		return nil
	})
}

// Delete method deletes data using key
func (db *badgerDB) Delete(key string) (err error) {
	return db.db.Update(
		func(txn *v3Badger.Txn) error {
			if err := txn.Delete([]byte(key)); err != nil {
				return errors.Errorf("can not delete data by key %s | %v", key, err)
			}
			return nil
		})
}

// Close method closes v3Badger db database
func (db *badgerDB) Close() error {
	if err := db.db.Close(); err != nil {
		return errors.Errorf("can't close database | %v", err)
	}
	return nil
}
