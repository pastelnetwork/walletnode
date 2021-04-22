package crypto

// DHKey interface provides all necessary methods to implement DH key
type DHKey interface {
	// PubKey returns public key of generated key pair
	PubKey() []byte
	// Shared makes a shared key using generated public key and passed public key
	Shared([]byte) error
	// Encrypt encrypts message using shared key
	Encrypt([]byte) ([]byte, error)
	// Decrypt decrypts message using generated private key
	Decrypt([]byte) ([]byte, error)
}
