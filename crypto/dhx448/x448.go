package dhx448

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/cloudflare/circl/dh/x448"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/crypto"
	"golang.org/x/crypto/sha3"
)

// DHX448 model represents public and private key based on x448 algorithm
type DHX448 struct {
	shared, prv, pub x448.Key
}

// New function returns DHX448 with already generated key pair
func New() crypto.DHKey {
	var pub, prv x448.Key
	_, _ = io.ReadFull(rand.Reader, prv[:])
	x448.KeyGen(&pub, &prv)
	return &DHX448{
		pub: pub,
		prv: prv,
	}
}

// PubKey method returns x448 public key
func (dhx448 *DHX448) PubKey() []byte {
	return toBytes(dhx448.pub)
}

// Shared method generates a shared x448 key based on own private
// and shared public key
func (dhx448 *DHX448) Shared(key []byte) error {
	pubKey, err := fromBytes(key)
	if err != nil {
		return errors.Errorf("can't make shared key | %v", err)
	}
	x448.Shared(&dhx448.shared, &(dhx448.prv), &pubKey)
	return nil
}

// Encrypt method encrypts message using  shared key
func (dhx448 *DHX448) Encrypt(plainText []byte) ([]byte, error) {
	h := sha3.New256()
	h.Write(dhx448.shared[:])
	hash := h.Sum(nil)

	block, err := aes.NewCipher(hash)
	if err != nil {
		return nil, errors.Errorf("error caused when trying to encrypt message | %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.Errorf("error caused when trying to encrypt message | %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	encText := base64.URLEncoding.EncodeToString(cipherText)
	return []byte(encText), nil
}

// Decrypt method decrypts message using private key
func (dhx448 *DHX448) Decrypt(secureMsg []byte) ([]byte, error) {
	h := sha3.New256()
	h.Write(dhx448.shared[:])
	hash := h.Sum(nil)

	cipherText, err := base64.URLEncoding.DecodeString(string(secureMsg))
	if err != nil {
		return nil, errors.Errorf("error caused when trying to decrypt message | %v", err)
	}

	block, err := aes.NewCipher(hash)
	if err != nil {
		return nil, errors.Errorf("error caused when trying to decrypt message | %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.Errorf("ciphertext block size is too short!")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}

// fromBytes function converts slice of bytes into x448.Key
func fromBytes(key []byte) (x448.Key, error) {
	if len(key) != x448.Size {
		return x448.Key{}, errors.Errorf("invalid x448 key")
	}
	var x448Key = x448.Key{}
	copy(x448Key[:], key)
	return x448Key, nil
}

// toBytes function convert x448.Key to slice of bytes
func toBytes(x448Key x448.Key) []byte {
	var key = make([]byte, x448.Size)
	copy(key, x448Key[:])
	return key
}
