package objects

type Envelope struct {
	Data []interface{} `json:"data"`
	Paging Paging `json:"paging"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
	Next string `json:"next"`
}

type Cursors struct {
	Before string `json:"before"`
	After string `json:"after"`
}