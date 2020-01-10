package platform

// ID is a unique string for a platform.
type ID string

// Service defines methods each platform must provide.
type Service interface {
	ID() ID
	Name() string
}
