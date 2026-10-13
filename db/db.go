package db
import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)
var DB *sql.DB

func InitDB(){
	var err error
	DB,err=sql.Open("pgx","postgres://postgres:YOUR_PASSWORD@localhost:5432/api?sslmode=disable")
	if err!=nil{
panic("unabale to connect with database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
}
func CreateTable(){
	 
}