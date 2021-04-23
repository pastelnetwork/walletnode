package dhx448

import (
	"testing"

	"github.com/cloudflare/circl/dh/x448"
	"github.com/stretchr/testify/assert"
)

func TestX448DH(t *testing.T) {
	x448KeyA, err := New()
	assert.Nil(t, err)
	x448KeyB, err := New()
	assert.Nil(t, err)
	pubKeyA := x448KeyA.PubKey()
	pubKeyB := x448KeyB.PubKey()
	// check the length of the generated public keys
	assert.Equal(t, len(pubKeyA), x448.Size)
	assert.Equal(t, len(pubKeyB), x448.Size)

	err = x448KeyA.Shared(x448KeyB.PubKey())
	assert.Nil(t, err)

	err = x448KeyB.Shared(x448KeyA.PubKey())
	assert.Nil(t, err)

	plainMessage := []byte("DHX448 message")

	encryptedMsgA, err := x448KeyA.Encrypt(plainMessage)
	assert.Nil(t, err)

	encryptedMsgB, err := x448KeyB.Encrypt(plainMessage)
	assert.Nil(t, err)

	originalMsgA, err := x448KeyA.Decrypt(encryptedMsgA)
	assert.Nil(t, err)
	assert.Equal(t, originalMsgA, plainMessage)

	originalMsgB, err := x448KeyB.Decrypt(encryptedMsgB)
	assert.Nil(t, err)
	assert.Equal(t, originalMsgB, plainMessage)
}
