package items

import "encoding/json"

type PublicItem struct {
	Id   int64  `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

type PrivateItem struct {
	Id   int64  `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

func (items Items) Marshall(isPublic bool) []interface{} {

	result := make([]interface{}, len(items))

	for index, item := range items {
		result[index] = item.Marshall(isPublic)

	}
	return result
}

func (item *Item) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicItem{
			Id:   item.Id,
			Item: item.Item,
			Done: item.Done,
		}
	}

	itemJson, _ := json.Marshal(item)
	var privateItem PrivateItem
	json.Unmarshal(itemJson, &privateItem)
	return privateItem

}
