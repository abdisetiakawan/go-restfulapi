package config

import (
	"database/sql"

	"github.com/abdisetiakawan/go-restfulapi/internal/helper"
)

func NewDB() *sql.DB {
    db, err := sql.Open("mysql", "root:@/restapi_go?charset=utf8mb4&parseTime=True&loc=Local")
    helper.PanicIfError(err)

    db.SetConnMaxIdleTime(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(5 * 60)

    return db
}