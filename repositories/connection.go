package repositories

import (
	/*
		github.com/go-sql-driver/mysql is not used directly by application
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db Store the Dataase connection
var Db *sqlx.DB

//OpenConnectionMySQL Function that Open Connection MySQL
func OpenConnectionMySQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:helloworld@tcp(127.0.0.1:3308)/testapp")
	//docker-compose up -d 
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
