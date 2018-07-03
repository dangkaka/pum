package main

type NCT struct {
	Url string
}

type NCTResponse struct {
}

func NewNCTHandler(url string) *NCT {
	return &NCT{url}
}

//TODO
func (nct *NCT) Get(link string) (*NCTResponse, error) {
	return nil, nil
}
