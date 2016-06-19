package objects

type GroupEnvelope struct {
	Data []Group `json:"data"`
	Paging Paging `json:"paging"`
}

type Group struct {
	Name string `json:"name"`
	Privacy string `json:"privacy"`
	Id string `json:"id"`
}