package app

import (
	"database/sql"
	"fxlik/golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/golang_restful_api?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
