package objects

type AlbumEnvelope struct {
	Data []Album `json:"data"`
	Paging Paging `json:"paging"`
}

type Album struct {
	Name string `json:"name"`
	Id string `json:"id"`
	CreatedTime int `json:"created_time"`
}