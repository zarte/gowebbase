package utils

import (
	"fmt"
	"github.com/zarte/comutil/Goconfig"
	Model "gowebbase/models"
	"strconv"
)

var cfg *Goconfig.ConfigFile

func GetConfigIni(filepath string) (err error) {
	config, err := Goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config

	var pagesize string
	if pagesize, err = cfg.GetValue("self", "PageSize"); err != nil {
		fmt.Println("配置文件中不存在PageSize", err)
		pagesize = "15"
	}
	Model.PageSize,err = strconv.Atoi(pagesize)
	if err !=nil {
		fmt.Println("pagesize err:"+err.Error())
	}
	return nil
}

func GetDatabase() (types, local, online string, err error) {
	if types, err = cfg.GetValue("database", "types"); err != nil {
		fmt.Println("配置文件中不存在types", err)
		return types, local, online, err
	}
	if local, err = cfg.GetValue("database", "local"); err != nil {
		fmt.Println("配置文件中不存在local", err)
		return types, local, online, err
	}
	if online, err = cfg.GetValue("database", "online"); err != nil {
		fmt.Println("配置文件中不存在online", err)
		return types, local, online, err
	}
	return types, local, online, nil
}
func GetIniVal(key string, sec string)string{
	if sec==""{
		sec = "self"
	}
	if val, err := cfg.GetValue(sec, key); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return val
	}else{
		return val
	}
}
func GetIniValInt(key string, sec string)int{
	if sec==""{
		sec = "self"
	}
	if val, err := cfg.GetValue(sec, key); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return 0
	}else{
		newval,err := strconv.Atoi(val)
		if err!=nil{
			fmt.Println("配置Atoifail", err)
			return 0
		}
		return newval
	}
}
func GetSelf() (port string, flag, tag int, err error) {
	if port, err = cfg.GetValue("self", "port"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return port, flag, tag, err
	}

	flag_temp, err := cfg.GetValue("self", "flag")
	if err != nil {
		fmt.Println("配置文件中不存在flag", err)
		return port, flag, tag, err
	}
	flag, err = strconv.Atoi(flag_temp)
	if err != nil {
		fmt.Println("配置文件中flag类型有误", err)
		return port, flag, tag, err
	}

	tag_temp, err := cfg.GetValue("self", "tag")
	if err != nil {
		fmt.Println("配置文件中不存在tag", err)
		return port, flag, tag, err
	}
	tag, err = strconv.Atoi(tag_temp)
	if err != nil {
		fmt.Println("配置文件中tag类型有误", err)
		return port, flag, tag, err
	}

	return port, flag, tag, nil
}