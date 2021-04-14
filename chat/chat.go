package chat

import (
	"github.com/pastelnetwork/walletnode/dao"
	"github.com/pastelnetwork/walletnode/nats"
)

type Chat struct {
	client *nats.Client
	db     dao.DB
}

func New(client *nats.Client, db dao.DB) *Chat {
	return &Chat{
		client: client,
		db:     db,
	}
}
