package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
    dsn := "root:@tcp(localhost:3306)/restapi_go?parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    return db, nil
}