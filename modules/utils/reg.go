package utils

import (
	"fmt"
	"regexp"
)

func RegPhone(str []byte) bool  {
	Reg,err := regexp.Compile(`^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`)
	if err != nil{
		fmt.Println("reg err ...",err)
		return false
	}
	return Reg.Match(str)
}


func RegPasswd(str []byte) bool {
	Reg, err := regexp.Compile(`^[A-Za-z0-9_\.\+\*\-]{6,16}$`)
	if err != nil {
		fmt.Println("reg err ...", err)
		return false
	}
	return Reg.Match(str)
}

func RegUsername(str []byte) bool {
	Reg, err := regexp.Compile(`^[\x{4e00}-\x{9fa5}A-Za-z0-9_]{5,16}$`)
	if err != nil {
		fmt.Println("reg err ...", err)
		return false
	}
	return Reg.Match(str)
}

func RegEmail(str []byte) bool {
	Reg, err := regexp.Compile(`^([0-9A-Za-z_\-]+)@[0-9a-z]+\.[a-z]{2,3}$`)
	if err != nil {
		fmt.Println("reg err ...", err)
		return false
	}
	return Reg.Match(str)
}
