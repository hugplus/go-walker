package dto

type DemoDto struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
