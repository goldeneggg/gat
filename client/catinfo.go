package client

// CatInfo show information of input
type CatInfo struct {
	// Input file name and data
	Files map[string][]byte
}

// NewCatInfo returns a new CatInfo given a map of files information
// files argument should be structured by file name key and file content value
func NewCatInfo(files map[string][]byte) *CatInfo {
	return &CatInfo{
		Files: files,
	}
}
