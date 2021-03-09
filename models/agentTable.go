package Model

import (
	"fmt"
)

type AgentInfo struct {
    Id      int    `xorm:"INT(10)"`
	Name   string `xorm:"VARCHAR(512)"`
    Appid   string `xorm:"VARCHAR(512)"`
    Secret    string `xorm:"VARCHAR(512)"`
    Pubsec string `xorm:"VARCHAR(512)"`
	State    int `xorm:"TINYINT(4)`
	Ctime  string    `xorm:"DATETIME"`
	Stime  string    `xorm:"DATETIME"`
}


func GetOneAgentByAppid(appid string) AgentInfo {
	var info AgentInfo

    engine := GetInstance()
	err := engine.Table("agent_info").Where("appid  =? ",appid).Where("status  =? ",0).Find(&info)
	if err != nil{
		fmt.Println(err)
		return info
	}
    return info
}

