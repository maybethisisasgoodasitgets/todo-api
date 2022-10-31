package domain
import()


type PublicItem  struct {

	Id int64 `json:"id"`
	Item string `json:"item"`
	Done bool `json:"done"`
	
}

type PrivateItem struct{
	Id int64 `json:"id"`
	Item string `json:"item"`
	Done bool `json:"done"`


}

func (items Items) Marshall(isPublic bool) []interface{}{
	
	
	result:= make([]interface{},len(items))

	for index, item:= range items{
		result[index]= item.Marshall(isPublic)

	}
	return result 
}
