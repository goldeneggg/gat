package client

// Client is a common interface for target services
type Client interface {
	// Check config to prepare to cat file
	CheckConf() error

	// Concatnate file
	Cat(catInf *CatInfo) (string, error)
}
