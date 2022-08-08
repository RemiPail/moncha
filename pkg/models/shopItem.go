package models

type ShopItem struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
	Desc    string `json:"desc"`
}
