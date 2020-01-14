package lib

import (
	"math/rand"
	"time"
)

// Rander is an instance of rand.Rand.
var Rander *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewRandomString generates a new string at the supplied length.
func NewRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[Rander.Intn(len(chars))]
	}
	return string(b)
}
