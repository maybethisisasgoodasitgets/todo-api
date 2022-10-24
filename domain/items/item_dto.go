package items


type Item struct {
	Id int64 `json:"id"`
	Item string `json:"item"`
	Done bool `json:"done"`
}

type Items []Item

