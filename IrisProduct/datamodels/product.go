package datamodels


type Product struct {
	ID int64 `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
	Image string `json:"image" sql:"image"`
	Num int64 `json:"num" sql:"num"`
	Url string `json:"url" sql:"url"`
}
