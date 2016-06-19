package objects

type PhotoEnvelope struct {
	Data []Photo `json:"data"`
	Paging Paging `json:"paging"`
}

type Photo struct {
	Id string `json:"id"`
	CreatedTime string `json:"created_time"`
}
