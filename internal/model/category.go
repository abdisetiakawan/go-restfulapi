package model

type Category struct {
    ID   uint   `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"type:varchar(200);not null"`
}