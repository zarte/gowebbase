package Model

import (
	"fmt"
	"time"
)

type Urole struct {
    Id      int    `xorm:"INT(10)"`
	Uid   int `xorm:"VARCHAR(512)"`
    Roleid   int `xorm:"VARCHAR(512)"`
	Status    int `xorm:"TINYINT(4)`
	Ctime  string    `xorm:"DATETIME"`
	Stime  string    `xorm:"DATETIME"`
}

func UpdateUrole(uid int,roleid []int) (int,error) {
	if len(roleid) <=0 {
		return 0,nil
	}
	engine := GetInstance()
	//删除原权限
	engine.Where("uid  =? ",uid).Update(&Urole{
		Status: 1,
	})
	//循环添加
	var list []Urole
	curtime := time.Now().Format("2006-01-02 15:04:05")
	for _,rid := range roleid{
		var item Urole
		item.Uid = uid
		item.Roleid = rid
		item.Ctime = curtime
		item.Stime = curtime
		list = append(list,item)
	}
	pos,err := engine.Insert(list)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}
	return int(pos),nil
}


func GetUroles(uid int) ([]int,error) {
	if uid<=0 {
		return nil,nil
	}
	engine := GetInstance()
	list := make([]*Urole, 0)
	err := engine.Where("uid  =? ",uid).Find(&list)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	res := make([]int,0)
	for _,v := range list{
		res = append(res,v.Roleid)
	}
	return res,nil
}