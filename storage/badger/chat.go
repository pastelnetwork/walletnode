package badger

import (
	v3Badger "github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/storage"

	"github.com/pastelnetwork/go-commons/log"
)

const (
	chatLogPrefix = "[chatdb]"
)

type ChatDB struct {
	db *v3Badger.DB
}

// NewChatDB creates an instance of ChatDB and starts fetching data
func NewChatDB(cfg *storage.Config) storage.KeyValue {
	ch := &ChatDB{db: &v3Badger.DB{}}
	if err := ch.start(cfg); err != nil {
		log.WithField("dir", cfg.ChatDBDir).Debugln(chatLogPrefix, "start")
		return nil
	}
	return ch
}

func (db *ChatDB) Init() error {
	return nil
}

// Init method opens the v3Badger db instance using specified configuration data and logger instance
func (db *ChatDB) start(cfg *storage.Config) (err error) {
	db.db, err = v3Badger.Open(v3Badger.DefaultOptions(cfg.ChatDBDir).WithLogger(log.NewLogger()))
	if err != nil {
		return errors.Errorf("can't fetch database from %v | %v", cfg.ChatDBDir, err)
	}
	return nil
}

// Get method fetches data by key
func (db *ChatDB) Get(key string) (result []byte, err error) {
	err = db.db.View(func(txn *v3Badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		if result, err = item.ValueCopy(result); err != nil {
			return errors.Errorf("can not fetch data by key %s | %v", key, err)
		}
		return nil
	})
	return
}

// Set method inserts data using key and value
func (db *ChatDB) Set(key string, value []byte) error {
	return db.db.Update(func(txn *v3Badger.Txn) error {
		if err := txn.Set([]byte(key), value); err != nil {
			return errors.Errorf("can not set data key %s value %s | %v", key, value, err)
		}
		return nil
	})
}

// Delete method deletes data using key
func (db *ChatDB) Delete(key string) (err error) {
	return db.db.Update(
		func(txn *v3Badger.Txn) error {
			if err := txn.Delete([]byte(key)); err != nil {
				return errors.Errorf("can not delete data by key %s | %v", key, err)
			}
			return nil
		})
}

// Close method closes v3Badger db database
func (db *ChatDB) Close() error {
	if err := db.db.Close(); err != nil {
		return errors.Errorf("can't close database | %v", err)
	}
	return nil
}
