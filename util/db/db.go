package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Action struct {
	TableName string        // 表名
	Query     string        // 查询条件
	Value     []interface{} // 查询值
	Sql       string        // 原生SQL
	Limit     int64         // 分页
	Offset    int64
	Total     interface{} // 总数
}

/*
	条件查询
*/
func (a Action) QueryAndFind(out interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	d := conn.Where(a.Query, a.Value...).Find(out)
	if a.Total != nil {
		d.Count(a.Total)
	}
}

/*
	普通查询
*/
func (a Action) Find(out interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	d := conn.Find(out)
	if a.Total != nil {
		d.Count(a.Total)
	}
}

/*
	条件查询加分页
*/
func (a Action) QueryAndPagination(out interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	d := conn.Where(a.Query, a.Value...).Count(a.Total).Limit(a.Limit).Offset(a.Offset).Find(out)
	if a.Total != nil {
		d.Count(a.Total)
	}
}

/*
	插入一条记录
*/
func (a Action) InsertOne(data interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	conn.Create(data)
}

/*
	执行SQL
*/
func (a Action) Exce() {
	conn := newConnection(a.TableName)
	defer conn.Close()
	conn.Exec(a.Sql)
}

/*
	SQL批量查询
*/
func (a Action) QueryBySQL(out interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	conn.Raw(a.Sql).Scan(out)
}

/*
	删除记录
*/
func (a Action) DeleteOne(model interface{}) {
	conn := newConnection(a.TableName)
	defer conn.Close()
	conn.Where(a.Query, a.Value...).Delete(model)
}

/*
	新建数据库连接
*/
func newConnection(tbName string) *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=flipped dbname=DMS sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	return db.Table(tbName)
}
