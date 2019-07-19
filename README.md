# 使用go语言开发一个个人博客系统，采用前后端分离技术，前端使用vue
### 初始包依赖
```go
import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
```
### mysql 连接配置
```go
//结构体定义---------------------------------
/*
Users mysql user表
*/
type Users struct {
	ID			int  	`json:"id"`
	User		string  `json:"user"`
	Password	string  `json:"password"`
	Email		string  `json:"email"`
	Name		string  `json:"name"`
	CreateTime	string  `json:"createTime"`
	UpdateTime	string  `json:"updateTime"`
}
//声明结构体变量-----------------------------
var myuser Users

//函数或方法中进行调用
db,err := sql.Open("mysql","root:opsbible@tcp(localhost:3306)/opsbible")
if err != nil {
	panic(err.Error())
}
defer db.Close()
err = db.Ping()
if err != nil {
	panic(err.Error())
}
fmt.Println("db is ok")
q,err := db.Query("select id,user,name,password,email,create_time,update_time from userwhere id = 0")
if err != nil {
	panic(err.Error())
}
defer q.Close()
for q.Next() {
	err = q.Scan(&myuser.ID, &myuser.User, &myuser.Name, &myuser.Password,  &myuser.Email,&myuser.CreateTime,&myuser.UpdateTime)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println()
}
```