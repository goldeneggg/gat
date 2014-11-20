package client

// Client is a common interface for target services
type Client interface {
	// Check config to prepare to cat file
	CheckConf() error

	// Concatnate file
	Cat(catInf *CatInfo) (string, error)

	// Concatnate file executing asynchronous parallel proccessing
	CatP(catInf *CatInfo, chOut chan string, chErr chan error)
}
