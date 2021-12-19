package data

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.EngineGroup

func init() {
	engine, err := xorm.NewEngineGroup("mysql", []string{"root:@tcp(127.0.0.1:3306)/dms?charset=utf8mb4"})
	if err != nil {
		panic(err)
	}
	err = engine.Sync(&Poet{})
	fmt.Println(err)
	Engine = engine
}

type Poet struct {
	Id      int    `xorm:"BIGINT UNSIGNED pk autoincr 'id'"`
	Author  string `xorm:"varchar(255) not null 'author' index" description:"作者"`
	Title   string `xorm:"varchar(255) not null 'title' index" description:"标题"`
	Content string `xorm:"TEXT not null 'content'" description:"内容"`
}

func (p *Poet) TableName() string {
	return "poet"
}

func InsertPoet(p *Poet) error {
	_, err := Engine.InsertOne(p)
	return err
}

func UpdatePoet(p *Poet) error {
	_, err := Engine.Where("id = ?", p.Id).Update(p)
	return err
}

func DeletePoet(id int) error {
	_, err := Engine.Where("id = ?", id).Delete(&Poet{})
	return err
}

func GetPoet(id int) (*Poet, error) {
	var ret Poet
	exist, err := Engine.Id(id).Get(&ret)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.New("不存在该id")
	}
	return &ret, nil
}

func GetPoetListByAuthor(author string, limit, offset int) ([]*Poet, int64, error) {
	var ret []*Poet
	count, err := Engine.Where("author = ?", author).Limit(limit, offset).FindAndCount(&ret)
	return ret, count, err
}

func GetPoetListByTitle(title string, limit, offset int) ([]*Poet, int64, error) {
	var ret []*Poet
	count, err := Engine.Where("title = ?", title).Limit(limit, offset).FindAndCount(&ret)
	return ret, count, err
}

func GetPoetListByContent(content string, limit, offset int) ([]*Poet, int64, error) {
	var ret []*Poet
	count, err := Engine.Where("content like ?", "%"+content+"%").Limit(limit, offset).FindAndCount(&ret)
	return ret, count, err
}
