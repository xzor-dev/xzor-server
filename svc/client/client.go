package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/xzor-dev/xzor-server/lib"
)

const (
	PrivateKeyLength = 2048
)

// Hash is a string representation of a client's public key.
type Hash string

// Client holds data for the individuals using xzor.
type Client struct {
	Key  *rsa.PrivateKey
	hash *Hash
}

// New creates a new client with a new private key.
func New() (*Client, error) {
	key, err := rsa.GenerateKey(rand.Reader, PrivateKeyLength)
	if err != nil {
		return nil, err
	}
	return &Client{
		Key: key,
	}, nil
}

// Hash returns the client's unique hash based on their private key.
func (c *Client) Hash() (Hash, error) {
	if c.hash == nil {
		pk, err := x509.MarshalPKIXPublicKey(c.Key.Public())
		if err != nil {
			return "", err
		}
		hash, err := lib.NewHash(pk)
		if err != nil {
			return "", err
		}
		ch := Hash(hash)
		c.hash = &ch
	}

	return *c.hash, nil
}
