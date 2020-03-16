package models

import (
	"database/sql"
	"fmt"
	"time"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/orm"
	"Dnsgslb/pkg/api/types"
)

const (
	USERNAME	= "root"
	PASSWORD	= "root"
	NETWORK		= "tcp"
	SERVER		= "localhost"
	PORT		=  3306
	DATABASE	= "dnsgslb"
)

var DB *sql.DB

func InitDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect fail err is :", err.Error())
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(200)
	DB.SetMaxIdleConns(16)
	return DB
}

func Selectcontent(recordid int) (string,error) {
	DB := InitDB()
	res := types.Records{}
	result, err := DB.Query("select domains.domain, names.name, nametypes.ttl, nametypes.persistence, types.type, records.disable, records.weight, contents.content, monitors.monitor, views.view, records.status " +
		"from domains " +
		"join names on domains.id = names.domain_id " +
		"join nametypes on names.id = nametypes.name_id " +
		"join types on nametypes.type_value = types.value " +
		"join records on nametypes.id = records.name_type_id " +
		"join content_monitors on records.content_monitors_id = content_monitors.id " +
		"join contents on content_monitors.content_id = contents.id " +
		"join monitors on content_monitors.monitor_id = monitors.id " +
		"join views on records.view_id = views.id where records.id = ?", recordid)
	if err != nil {
		return res.Content, err
	}
	defer DB.Close()
	if result.Next() {
		result.Scan(&res.Domain, &res.Name, &res.TTL, &res.Persistence, &res.Type, &res.Disable, &res.Weight, &res.Content, &res.Monitors, res.View, res.Status)
	}
	return res.Content, nil
}

func ListRecords(res *types.Records, args int) (*types.Records, error) {
	DB := InitDB()
	result, err := DB.Query("select domains.domain, names.name, nametypes.ttl, nametypes.persistence, types.type, records.disable, records.weight, contents.content, monitors.monitor, views.view, records.status " +
		"from domains " +
		"join names on domains.id = names.domain_id " +
		"join nametypes on names.id = nametypes.name_id " +
		"join types on nametypes.type_value = types.value " +
		"join records on nametypes.id = records.name_type_id " +
		"join content_monitors on records.content_monitors_id = content_monitors.id " +
		"join contents on content_monitors.content_id = contents.id " +
		"join monitors on content_monitors.monitor_id = monitors.id " +
		"join views on records.view_id = views.id where records.id = ?", args)
	if err != nil {
		return res, err
	}
	defer DB.Close()
	if result.Next() {
		result.Scan(&res.Domain, &res.Name, &res.TTL, &res.Persistence, &res.Type, &res.Disable, &res.Weight, &res.Content, &res.Monitors, res.View, res.Status)
	}
	return res, nil
}

func SelectRecords(res *types.Records) (map[int]types.Records, error) {
	DB := InitDB()
	result, err := DB.Query("select records.id, domains.domain, names.name, nametypes.ttl, nametypes.persistence, types.type, records.disable, records.weight, contents.content, monitors.monitor, views.view, records.status " +
		"from domains " +
		"join names on domains.id = names.domain_id " +
		"join nametypes on names.id = nametypes.name_id " +
		"join types on nametypes.type_value = types.value " +
		"join records on nametypes.id = records.name_type_id " +
		"join content_monitors on records.content_monitors_id = content_monitors.id " +
		"join contents on content_monitors.content_id = contents.id " +
		"join monitors on content_monitors.monitor_id = monitors.id " +
		"join views on records.view_id = views.id")
	listrecords := make(map[int]types.Records)
	defer func(){
		if result != nil {
			result.Close()
		}
	}()
	if err != nil {
		return listrecords, err
	}
	var count int
	count = 0
	for result.Next() {
		err = result.Scan(&res.Id, &res.Domain, &res.Name, &res.TTL, &res.Persistence, &res.Type, &res.Disable, &res.Weight, &res.Content, &res.Monitors, &res.View, &res.Status)
		if err != nil {
			return listrecords, err
		}
		listrecords[count] = *res
		count ++
	}

	return listrecords, nil
}

func InsertRecords(name string, types string, content string, ttl int, disable int, weight int, monitors string, view string) (int64,error) {
	DB := InitDB()
	var id int64
	result, err := DB.Exec("insert into records (name_type_id, content_monitors_id, view_id, disable, ttl, weight)" +
		"select(select nametypes.id from nametypes join names on nametypes.name_id = names.id join types on nametypes.type_value = types.value where names.name = ? and types.type = ?) as name_type_id," +
		"(select content_monitors.id from content_monitors join contents on content_monitors.content_id = contents.id join monitors on content_monitors.monitor_id = monitors.id where contents.content = ? and monitors.monitor = ?) as content_monitors_id," +
		"(select views.id from views where view = ? ) as view_id, (?) as disable, (?) as ttl, (?) weight", name, types, content, monitors, view, disable, ttl, weight)
	if err != nil {
		return id, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}
	defer DB.Close()
	return id, nil
}

func DelRecords(id int64) error {
	DB := InitDB()
	// 解除外键约束
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0");
	_, err := DB.Exec("delete records, content_monitors, contents from records " +
		"join nametypes on records.name_type_id = nametypes.id " +
		"join names on nametypes.name_id = names.id " +
		"join content_monitors on records.content_monitors_id = content_monitors.id " +
		"join contents on content_monitors.content_id = contents.id where records.id = ? ", id)
	if err != nil {
		return err
	}
	// 重新设置外键删除约束
	DB.Exec("SET FOREIGN_KEY_CHECKS = 1");
	defer DB.Close()
	return nil
}

func ChangeRecords(name string, domain string, types string, content string, ttl int, disable int, weight int, monitors string, view string, recordsId int64) error {
	DB := InitDB()
	_, err := DB.Exec("update records set name_type_id = " +
		"(select nametypes.id from nametypes join names on nametypes.name_id = names.id " +
		"join domains on names.domain_id = domains.id " +
		"join types on nametypes.type_value = types.value where names.name = ? and domains.domain = ? and types.type = ?), " +
		"content_monitors_id = " +
		"(select content_monitors.id from content_monitors join contents on content_monitors.content_id = contents.id " +
		"join monitors on content_monitors.monitor_id = monitors.id where contents.content = ? and monitors.monitor = ?), " +
		"view_id = (select views.id from views where views.view = ?), disable = ?, ttl = ?, weight = ? where records.id = ? ", name, domain, types, content, monitors, view, disable, ttl, weight, recordsId)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeStatus(status bool, recordsId int64) error {
	DB := InitDB()
	_, err := DB.Exec("update records set status = ? where id = ?", status, recordsId)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}