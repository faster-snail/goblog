package dbops

import (
	"database/sql"
	"log"
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

//SelectOneByUserName 根据id获取一个用户
func (u *SysUser) SelectOneByUserName () {
	db := GoBlog()
	defer db.Close()
	stmt,err := db.Prepare("SELECT" +
	" id,user_name,passwd,email,head_portrait,"+
	"prefession,is_online,create_time,update_time,last_login_time"+
	" FROM sys_user where user_name = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.UserName).Scan(
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

//UpdateOneByUserName 修改个用户
func (u SysUser) UpdateOneByUserName () (lstID int64,rowCnt int64){
	db :=  GoBlog()
	defer db.Close()
	var user SysUser
	user.UserName = u.UserName
	user.SelectOneByUserName()
	// 设定值的时候需要做判断，如果当前的结构体数据为空，则查询数据，并传入当前结构体中
	if u.Passwd == "" {
		u.Passwd = user.Passwd
	}
	if u.Prefession == "" {
		u.Prefession = user.Prefession
	}
	if u.Email == "" {
		u.Email = user.Email
	}
	if u.HeadPortrait == "" {
		u.HeadPortrait = user.HeadPortrait
	}
	u.UpdateTime = utils.NowDateTime()
	stmt,err := db.Prepare("update sys_user set passwd = ?,"+
	"email = ?,"+
	"head_portrait = ?,"+
	"prefession = ?,"+
	"update_time = ? "+
	"where user_name = ?")

	if err != nil {
		log.Fatal(err.Error())
	}
	res,err := stmt.Exec(u.Passwd,u.Email,u.HeadPortrait,u.Prefession,u.UpdateTime,u.UserName)
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

//DeleteOneByID 删除一个用户
func (u SysUser) DeleteOneByID () (lstID int64,rowCnt int64) {
	db := GoBlog()
	defer db.Close()
	stmt,err := db.Prepare("DELETE FROM sys_user where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	res,err := stmt.Exec(u.ID)
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
//SelectUserList 获取用户列表
func (u SysUser) SelectUserList () []SysUser {
	var  user []SysUser
	db := GoBlog()
	defer db.Close()
	stmt,err := db.Prepare("SELECT" +
	" id,user_name,passwd,email,head_portrait,"+
	"prefession,is_online,create_time,update_time,last_login_time"+
	" FROM sys_user")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.ID,
			&u.UserName,
			&u.Passwd,
			&u.Email,
			&u.HeadPortrait,
			&u.Prefession,
			&u.IsOnline,
			&u.CreateTime,
			&u.UpdateTime,
			&u.LastLoginTime)
			user = append(user,u)
	}
	return user
}

//OnLine 在线
func (u SysUser) OnLine () {
	u.IsOnline = 1
	lstID,rowCnt := u.UpdateOneByUserName()
	if lstID != 0 && rowCnt != 1 {
		panic("设置在线失败！")
	}
}

//OutLine 离线
func (u SysUser) OutLine () {
	u.IsOnline = 0
	lstID,rowCnt := u.UpdateOneByUserName()
	if lstID != 0 && rowCnt != 1 {
		panic("设置离线失败！")
	}
}