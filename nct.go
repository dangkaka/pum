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

func (nct *NCT) GetDownloadObject() (*DownloadObject, error) {
	response, err := nct.Parse()
	if err != nil {
		return nil, err
	}
	return &DownloadObject{
		Url:         nct.Url,
		DownloadUrl: response.DownloadUrl,
	}, nil
}

//TODO
func (nct *NCT) Parse() (*NCTResponse, error) {
	return nil, nil
}
