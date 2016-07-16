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

//func (c *Comment) UnmarshalJSON(b []byte) error {
//	type aux struct {
//		CreatedTime string `json:"created_time"`
//		From User `json:"from"`
//		Message string `json:"message"`
//		Id string `json:"id"`
//	}
//
//	t := aux{}
//	err := json.Unmarshal(b, t)
//	if err != nil {
//		return err
//	}
//	c.Id = t.Id
//	c.Message = t.Message
//	i, err := strconv.ParseInt(t.CreatedTime, 10, 64)
//	if err != nil {
//		log.Fatalf("unable to parse time: %d",err)
//	}
//	c.CreatedTime = time.Unix(i, 0)
//	c.From = t.From
//	return nil
//}