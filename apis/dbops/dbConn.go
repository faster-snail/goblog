package dbops

import (
	"database/sql"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)
//DBcon is a config struct
type DBcon struct {
	Host	string
	Port	string
	User	string
	Passwd	string
	DB		string
}

//后续操作要加上 dbc.Close()
func (c DBcon) dbc () *sql.DB {
	constr := c.User + ":" + c.Passwd + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DB
	db,err := sql.Open("mysql",constr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//GoBlog 是goblog库的链接
func GoBlog () *sql.DB {
	var db = DBcon{Host: "localhost",Port: "3306",User: "root",Passwd: "123456",DB: "goblog"}
	return db.dbc()
}


