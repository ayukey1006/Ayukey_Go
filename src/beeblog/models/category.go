package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//数据结构
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

//添加分类
func AddCateGory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return fmt.Errorf("该分类已经存在")
	}

	_, err = o.Insert(cate)
	return err
}

//删除分类
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := Category{Id: cid}
	_, err = o.Delete(&cate)
	return err
}

func ModifyCategory(name string, count int64) error {
	o := orm.NewOrm()
	cate := &Category{
		Title: name,
	}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err != nil {
		return err
	}

	cate.TopicCount += count

	_, err = o.Update(cate)
	return err
}

//获取分类列表
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}
