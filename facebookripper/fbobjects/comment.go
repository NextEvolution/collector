package objects

type CommentEnvelope struct {
	Data []Comment `json:"data"`
	Paging Paging `json:"paging"`
}

type Comment struct {
	CreatedTime string `json:"created_time"`
	From User `json:"from"`
	Message string `json:"message"`
	Id string `json:"id"`
}
