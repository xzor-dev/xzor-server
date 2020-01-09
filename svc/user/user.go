package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/xzor-dev/xzor-server/lib"
)

const (
	// PrivateKeyLength defines the length of a user's private key.
	PrivateKeyLength = 2048
)

// Hash is a string representation of a user's public key.
type Hash string

// User holds data for the individuals using xzor.
type User struct {
	Key  *rsa.PrivateKey
	hash *Hash
}

// New creates a new user with a new private key.
func New() (*User, error) {
	key, err := rsa.GenerateKey(rand.Reader, PrivateKeyLength)
	if err != nil {
		return nil, err
	}
	return &User{
		Key: key,
	}, nil
}

// Hash returns the user's unique hash based on their private key.
func (u *User) Hash() (Hash, error) {
	if u.hash == nil {
		pk, err := x509.MarshalPKIXPublicKey(u.Key.Public())
		if err != nil {
			return "", err
		}
		hash, err := lib.NewHash(pk)
		if err != nil {
			return "", err
		}
		ch := Hash(hash)
		u.hash = &ch
	}

	return *u.hash, nil
}
