package objects

type PhotoEnvelope struct {
	Data []Photo `json:"data"`
	Paging Paging `json:"paging"`
}

type Photo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	//CreatedTime string `json:"created_time"`
	//Images []Image `json:"images"`
}

type Image struct {
	Height int `json:"height"`
	Width int `json:"width"`
	Source string `json:"source"`
}

//func (p *Photo) UnmarshalJSON(b []byte) error {
//	type aux struct {
//		Id string `json:"id"`
//		Name string `json:"name"`
//		CreatedTime string `json:"created_time"`
//		Images []Image `json:"images"`
//	}
//
//	t := aux{}
//	err := json.Unmarshal(b, t)
//	if err != nil {
//		return err
//	}
//	p.Id = t.Id
//	p.Name = t.Name
//	i, err := strconv.ParseInt(t.CreatedTime, 10, 64)
//	if err != nil {
//		log.Fatalf("unable to parse time: %d",err)
//	}
//	p.CreatedTime = time.Unix(i, 0)
//	p.Images = t.Images
//return nil
//}