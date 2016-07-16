package objects

type PhotoEnvelope struct {
	Data []Photo `json:"data"`
	Paging Paging `json:"paging"`
}

type Photo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatedTime int `json:"created_time"`
	Images []Image `json:"images"`
}

type Image struct {
	Height int `json:"height"`
	Width int `json:"width"`
	Source string `json:"source"`
}