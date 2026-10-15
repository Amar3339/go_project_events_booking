package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://postgres:3339@localhost:5432/api?sslmode=disable")
	if err != nil {
		panic("unabale to connect with database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	CreateTable()
}
func CreateTable() {
	
CreateEventsTable:=`
 create table if not exists events(
Id   serial primary key  ,
 Name text not null ,
Address  text not null ,
description text not null ,
DateTime timestamp,
UserId int 
)`
_, err := DB.Exec(CreateEventsTable)

if err != nil {
    panic(err)
}
}
