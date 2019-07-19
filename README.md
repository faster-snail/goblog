## 个人博客系统，采用前后端分离技术，前端使用vue，后端用go
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
> 结构体定义---------------------------------
```go
/*
Users mysql user表
*/
type Users struct {
	ID          int     `json:"id"`
	User		string  `json:"user"`
	Password	string  `json:"password"`
	Email		string  `json:"email"`
	Name		string  `json:"name"`
	CreateTime	string  `json:"createTime"`
	UpdateTime	string  `json:"updateTime"`
}
```
> 声明结构体变量-----------------------------
```go
var myuser Users
var myusers []Users
```
> 函数或方法中进行调用-----------------------
```go
func getUser (w http.ResponseWriter, r *http.Request) {
    db,err := sql.Open("mysql","root:opsbible@tcp(localhost:3306)/opsbible")
    if err != nil {
    	panic(err.Error())
    }
    defer db.Close()
    if db.Ping() != nil {
    	panic(err.Error())
    }
    q,err := db.Query("select id,user,name,password,email,create_time,update_time from user where id = 0")
    if err != nil {
    	panic(err.Error())
    }
    defer q.Close()
    for q.Next() {
    	err = q.Scan(&myuser.ID, &myuser.User, &myuser.Name, &myuser.Password,  &myuser.Email,& myuser.CreateTime,&myuser.UpdateTime)
    	if err != nil {
    		panic(err.Error())
        }
        myusers = append(myusers,myuser)
    }

    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myusers)
}

```
### mux 操作
```go
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user", getUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",r))
}
```
