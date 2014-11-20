package client

// Information of input
type CatInfo struct {
	// Input file name and data
	Files map[string][]byte
}

// Create a new CatInfo
func NewCatInfo(files map[string][]byte) *CatInfo {
	return &CatInfo{
		Files: files,
	}
}
