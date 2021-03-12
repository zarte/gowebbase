package Model

import (
	"fmt"
	"time"
)

type Loginlog struct {
    Id      int    `xorm:"INT(10)"`
	Uid   string `xorm:"VARCHAR(512)"`
    Ip   string `xorm:"VARCHAR(512)"`
	Ctime  string    `xorm:"DATETIME"`
}

func AddLoginLog(uid string,ip string) (int,error) {
	engine := GetInstance()
	info := Loginlog{
		Uid:  uid,
		Ip:   ip,
		Ctime: time.Now().Format("2006-01-02 15:04:05"),
	}
	pos,err := engine.Table("user").InsertOne(info)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}
	return int(pos),nil
}