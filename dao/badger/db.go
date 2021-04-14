package badger

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/dao"
)

type DB struct {
	*badger.DB

	path string
}

// Init implements dao.badger.DB.Init()
func (db *DB) Init() (err error) {
	db.DB, err = badger.Open(badger.DefaultOptions(db.path).WithLogger(&Logger{}))
	return errors.New(err)
}

func New(path string) dao.DB {
	return &DB{
		path: path,
	}
}
