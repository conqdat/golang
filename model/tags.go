package model

type Tags struct {
	Id   int `gorm:"type:int; primary_key"`
	Name int `gorm:"type:varchar(255)"`
}
