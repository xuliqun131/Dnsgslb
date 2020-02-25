package models

import (
	"database/sql"
	"fmt"
	"time"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/orm"
)

const (
	USERNAME	= "powerdns"
	PASSWORD	= "root"
	NETWORK		= "tcp"
	SERVER		= "localhost"
	PORT		=  3306
	DATABASE	= "powerdns"
)

type Pdnsrecords struct {
	Id				int
	Domain_id		int
	Type 			string
	Name 			string
	Content 		string
}

var DB *sql.DB


func InitDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect fail err is :", err.Error())
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	fmt.Println("connnect success")

	return DB
}


func InsertRecords(domain string, name string, types string, content string, ttl int, disable string, weight string, monitors string, view string) error {
	DB := InitDB()
	result, err := DB.Exec("insert INTO records(name, type, content) values(?,?,?)", name, types, content)
	if err != nil {
		fmt.Println("insert err is :", err.Error())
	}
	lastInsertID, err := result.LastInsertId()

	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return err
	}

	fmt.Println("LastInsertID:", lastInsertID)
	return nil
}









//func main() {
//	err := Init()
//	if err != nil {
//		fmt.Println("Init err :", err)
//	}
//}
//
//func Init() error {
//	err := orm.RegisterDataBase("powerdns", "mysql", "powerdns:root@tcp(localhost:3306)/powerdns?charset=utf8")
//	if err != nil {
//		fmt.Println("the err is :", err.Error())
//		return err
//	}
//	o := orm.NewOrm()
//	pdns := Pdnsrecords{
//		Id:                     90,
//		Domain_id:              4,
//		Type:                   "A",
//		Name:                   "dio.xu.com",
//		Content:                "192.0.2.102",
//	}
//	id, err := o.Insert(&pdns)
//	if err != nil {
//		fmt.Println("insert err :", err)
//	}
//	fmt.Printf("插入了数据,id是%d\n",id)
//
//	return nil
//}
