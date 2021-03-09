package Model

import (
	"fmt"
	"time"
)

type Tradeorder struct {
    Id      int    `xorm:"INT(10)"`
	Orderno   string `xorm:"VARCHAR(512)"`
    Tradeno   string `xorm:"VARCHAR(512)"`
	Productname    string `xorm:"VARCHAR(512)"`
    Price int `xorm:"INT(11)"`
	State    int `xorm:"TINYINT(4)`
	Overtime    string    `xorm:"DATETIME"`
	Ctime  string    `xorm:"DATETIME"`
	Stime  string    `xorm:"DATETIME"`
}


func CheckOrderNo(orderNo string) Tradeorder {
	var info Tradeorder
	overtime := time.Now().Format("2006-01-02 15:04:05")
    engine := GetInstance()
	err := engine.Table("agent_info").Where("orderno  =? ",orderNo).Where("status  =? ",0).Where("overtime  >? ",overtime).Find(&info)
	if err != nil{
		fmt.Println(err)
		return info
	}
    return info
}

