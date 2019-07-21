package main

import (
	"fmt"
	"github.com/faster-snail/goblog/apis/dbops"
)

func main() {
	// var user dbops.SysUser
	// user.ID = "FG23SDTASDYYY4123"
	// user.SelectOneByID()
	// user.ID = "FG23SDTASDXDN6533"
	// user.UserName = "yuyu"
	// lstID,rowCnt := user.InsertOne()
	// fmt.Println(user)
	var  u dbops.SysUser
	u.UserName = "songxiao"
	u.Email = "opsbible@foxmial.com"
	u.Passwd = "123456"
	u.HeadPortrait = "http://www.opsbible/images/head/1.png"
	u.Prefession = "it"
	lstid,rowCnt := u.InsertOne()
	if rowCnt < 1 {
		fmt.Println("is false!")	
	}
	fmt.Println(lstid,rowCnt)

}