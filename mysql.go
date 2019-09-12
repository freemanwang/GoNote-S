package main

import (
	"database/sql"
	"fmt"
	"time"

	//"encoding/base64"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"

	//"github.com/imroc/biu"
	//"github.com/jmoiron/sqlx"
	//"io/ioutil"
	"log"
)

type QRCode struct {
	id int
	build string
	reqID int64
	appID int64
	addTime, title, remark string
	qrCode sql.NullString
}

type Person struct {
	id int
	name string
}

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "go"
)

func main(){
	dbconfig := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db, err := sql.Open("mysql", dbconfig)
	if err != nil {
		log.Println("open databse failed, err: ",err)
	}
	log.Println(db)

	db.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)//设置最大连接数
	db.SetMaxIdleConns(16) //设置闲置连接数

	rows, err := db.Query("select * from test")
	if err != nil {
		log.Fatal("查询失败" ,err)
	}

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id: %d   name: %s\n", id, name)
	}


	// 读取 mysql 中 qrcode表 的数据项
	//qrcode := [10]QRCode{}
	//rows, err = db.Query("select * from qrcode" )
	//i := 0
	//for rows.Next() {
	//	err = rows.Scan(&qrcode[i].id, &qrcode[i].build, &qrcode[i].reqID, &qrcode[i].appID, &qrcode[i].addTime, &qrcode[i].title, &qrcode[i].desc, &qrcode[i].qrCode)
	//	if err != nil {
	//		panic(err)
	//	}
	//	i ++
	//}
	//i --
	//for i >= 0 {
	//	fmt.Println(qrcode[i])
	//	i --
	//}
	//fmt.Printf( "%T", db)

	//往 mysql 中 qrcode表 insert 内容
	insertqrcode := QRCode{
		//id:      0,
		build:   "8.17",
		reqID:   2564783542311,
		appID:   113542,
		//addTime: "",
		title:   "新插入需求4",
		remark:    "新插入需求4的描述字段啊啊啊啊",
		qrCode:  sql.NullString{},
	}

	//_, err = db.Exec(
	//	"insert into qrcode (build, reqID, appID, title, remark, qrcode) values (?, ?, ?, ?, ?, ?)",
	//	insertqrcode.build, insertqrcode.reqID, insertqrcode.appID, insertqrcode.title, insertqrcode.remark, insertqrcode.qrCode,
	//	)

	insertQRCode(db, insertqrcode)

	queryQRCode(db)

}

func insertQRCode(db *sql.DB, insertqrcode QRCode) bool{
	if db == nil {
		return false;
	}

	_, err := db.Exec(
		"insert into qrcode (build, reqID, appID, title, remark, qrcode) values (?, ?, ?, ?, ?, ?)",
		insertqrcode.build, insertqrcode.reqID, insertqrcode.appID, insertqrcode.title, insertqrcode.remark, insertqrcode.qrCode,
	)
	if err != nil {
		return false
	}
	return true
}


func queryQRCode(db *sql.DB){

	// 读取 mysql 中 go表 的数据项
	qrcode := [10]QRCode{}
	rows, err := db.Query("select * from qrcode" )

	i := 0   // 记录多少条数据

	// 查询并写入 qrcode 数组
	for rows.Next() {
		err = rows.Scan(&qrcode[i].id, &qrcode[i].build, &qrcode[i].reqID, &qrcode[i].appID, &qrcode[i].addTime, &qrcode[i].title, &qrcode[i].remark, &qrcode[i].qrCode)
		checkErr(err)
		i ++
	}

	//输出
	for i >= 0 {
		fmt.Println(qrcode[i])
		i --
	}
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}