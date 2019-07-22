package dbops

import (
	"database/sql"
	"log"
	"github.com/faster-snail/goblog/utils"
)
//Article 表结构
type Article struct {
	ID				int64
	CreateTime		string
	UpdateTime		string	
}

//WriteOne 创建一篇文章
func (a Article) WriteOne () (lstID int64,rowCnt int64) {
}

//DeleteOneByID 删除一篇文章
func (a Article) DeleteOneByID () (lstID int64,rowCnt int64) {

}

//UpdateOneByID 更新一篇文章
func (a Article) UpdateOneByID () (lstID int64,rowCnt int64){

}

//ListArticleOrderByCreateTime 根据文章创建时间查询列表
func (a *Article) ListArticleOrderByCreateTime () {

}


//Show 显示
func (a Article) Show () {

}

//hiden 离线
func (a Article) hiden () {

}