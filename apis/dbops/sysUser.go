package dbops

import (
	"database/sql"
	"log"
	"time"
	"github.com/faster-snail/goblog/utils"
)
//SysUser 表结构
type SysUser struct {
	ID				int64
	UserName		string
	Passwd			string
	Email			string
	HeadPortrait	string
	Prefession		string
	IsOnline		int8
	CreateTime		string
	UpdateTime		string
	LastLoginTime	string
}

//SelectOneByID 根据id获取一个用户
func (u *SysUser) SelectOneByID () {
	db := GoBlog()
	defer db.Close()
	stmt,err := db.Prepare("SELECT" +
	" id,user_name,passwd,email,head_portrait,"+
	"prefession,is_online,create_time,update_time,last_login_time"+
	" FROM sys_user where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.ID).Scan(
		&u.ID,
		&u.UserName,
		&u.Passwd,
		&u.Email,
		&u.HeadPortrait,
		&u.Prefession,
		&u.IsOnline,
		&u.CreateTime,
		&u.UpdateTime,
		&u.LastLoginTime)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err.Error())
	}
}

//InsertOne 创建一个用户
//returing 最后更改id 和修改的行数
func (u SysUser) InsertOne () (lstID int64,rowCnt int64) {
	db := GoBlog()
	defer db.Close()
	now := 	utils.NowDateTime()
	u.CreateTime = now
	u.UpdateTime = now
	u.LastLoginTime = now
	u.IsOnline = 1
	u.ID = utils.NewID()	//唯一id生成
	stmt,err := db.Prepare("INSERT INTO "+
	"sys_user(id, user_name, passwd, email, head_portrait, "+
	"prefession, is_online, create_time, update_time, last_login_time) "+
	"VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	res,err := stmt.Exec(u.ID,
			u.UserName,
			u.Passwd,
			u.Email,
			u.HeadPortrait,
			u.Prefession,
			u.IsOnline,
			u.CreateTime,
			u.UpdateTime,
			u.LastLoginTime)
	if err != nil {
		log.Fatal(err.Error())
	}
	lstID,err = res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	rowCnt,err = res.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}
	return lstID,rowCnt
}

//UpdateOneByUserName 修改个用户
func (u SysUser) UpdateOneByUserName () {

}
//DeleteOneByID 删除一个用户
func (u SysUser) DeleteOneByID () {

}
//GetUserList 获取用户列表
func (u SysUser) GetUserList () []SysUser {
	var  user []SysUser
	return user
}

//OnLine 在线
func (u SysUser) OnLine () {

}

//OutLine 离线
func (u SysUser) OutLine () {

}