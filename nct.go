package main

type NCT struct {
	Url string
}

type NCTResponse struct {
	DownloadUrl string
}

func NewNCTHandler(url string) *NCT {
	return &NCT{url}
}

func (nct *NCT) GetBest() (string, error) {
	response, err := nct.Get()
	if err != nil {
		return "", err
	}
	return response.DownloadUrl, nil
}

//TODO
func (nct *NCT) Get() (*NCTResponse, error) {
	return nil, nil
}
