package objects

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After string `json:"after"`
}