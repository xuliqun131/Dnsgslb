package handlers

import (
	"database/sql"
	"os/exec"
	"fmt"
	"time"

	"strings"
	"strconv"
)

const (
	USERNAME	= "root"
	PASSWORD	= "root"
	NETWORK		= "tcp"
	SERVER		= "localhost"
	PORT		=  3306
	DATABASE	= "powerdns"
)

func InitPdnsDB() *sql.DB {

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

func Online(domain string, name string, types string, ttl int, content string) error {
	ttlstr := strconv.Itoa(ttl)
	names := strings.Split(name, ".")
	cmd := exec.Command("/bin/bash", "-c", "pdnsutil add-record "+ domain + " "+ names[0] +" "+ types +" "+ ttlstr +" "+ content)
	fmt.Println("pdnsutil add-record "+ domain + " "+ names[0] +" "+ types +" "+ ttlstr +" "+ content)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Erro:The command is err", err)
	}

	return nil
}

func Offline(content string) error {
	DB := InitPdnsDB()
	_, err := DB.Exec("delete from records where content = ?", content)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}
