package Model

import (
	"errors"
	"fmt"
	"time"
	"xorm.io/xorm"
)

type User struct {
    Id      int    `xorm:"INT(10)"`
	Username   string `xorm:"VARCHAR(512)"`
    Passwd   string `xorm:"VARCHAR(512)"`
	Salt   string `xorm:"VARCHAR(512)"`
	Email    string `xorm:"VARCHAR(512)"`
	Roles    string `xorm:"VARCHAR(512)"`
    Point int `xorm:"INT(11)"`
	Money int `xorm:"INT(11)"`
	Phone    string `xorm:"VARCHAR(512)"`
	Status    int `xorm:"TINYINT(4)`
	Ctime  string    `xorm:"DATETIME"`
	Stime  string    `xorm:"DATETIME"`
	BaseModel        `json:"-" xorm:"-"`
}


func (task *User) parseWhere(session *xorm.Engine, params CommonMap) {
	if len(params) == 0 {
		return
	}
	username, ok := params["username"]
	if ok && username.(string) !="" {
		session.Where("username like ?", "%"+username.(string)+"%")
	}
	roleid, ok := params["roleid"]
	if ok && roleid.(string) !="" {
		session.Where("roles like ?", "%,"+roleid.(string)+",%")
	}
}

func (task *User) Total(params CommonMap) (int64, error) {
	engine := GetInstance()
	task.parseWhere(engine, params)
	total, err := engine.Where("status  =? ",0).Count(task)
	return total, err
}
func (task *User) GetUserPage(params CommonMap) ([]User,error) {
	engine := GetInstance()
	task.parseWhere(engine, params)
	var list []User
	err := engine.Table("user").Where("status  =? ",0).OrderBy("id desc").Limit(task.PageSize, task.pageLimitOffset()).Find(&list)
	if err != nil{
		fmt.Println(err)
		return list,err
	}
	return list,nil
}

func CheckUserNameExist(username string) bool {
	var info User
    engine := GetInstance()
	res,err := engine.Where("username  =? ",username).Get(&info)
	if err != nil{
		fmt.Println(err)
		return false
	}
	return res
}

func GetUserByName(username string) (User,error) {
	var info User
	engine := GetInstance()
	res,err := engine.Where("username  =? ",username).Where("status  =? ",0).Get(&info)
	if err != nil{
		fmt.Println(err)
		return info,err
	}
	if !res {
		return info,errors.New("不存在")
	}else {
		return info,nil
	}

}

func (task *User)GetUserById(id int) (User,error) {
	var info User
	engine := GetInstance()
	res,err := engine.Where("id  =? ",id).Where("status  =? ",0).Get(&info)
	if err != nil{
		fmt.Println(err)
		return info,err
	}
	if !res {
		return info,errors.New("不存在")
	}else {
		return info,nil
	}
}


func AddUserInfo(info User) (int,error) {
	engine := GetInstance()
	curtime :=time.Now().Format("2006-01-02 15:04:05")
	info.Ctime = curtime
	info.Stime = curtime
	pos,err := engine.InsertOne(info)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}
	return int(pos),nil
}
func UpdateUserInfoById(id int,info User) (int,error) {
	engine := GetInstance()
	curtime :=time.Now().Format("2006-01-02 15:04:05")
	info.Stime = curtime
	pos,err := engine.Table("user").ID(id).Cols(`passwd,salt,email,point,money,phone,stime,roles`).Update(info)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}
	return int(pos),nil
}
func DeleteUserById(id int) (int,error) {
	engine := GetInstance()
	user := new(User)
	user.Status = 1
	pos,err := engine.Table("user").Where("id  =? ",id).Limit(1).Update(user)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}
	return int(pos),nil
}
