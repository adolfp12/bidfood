package model

type Product struct {
	Id   int
	Name string `json:"name"`
	Desc string `json:"desc"`
}
