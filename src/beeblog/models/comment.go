package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//数据结构
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

//添加评论
func AddReply(tid, nickname, content string) (*Comment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return reply, err
}

//删除评论
func DeleteReply(id string) error {
	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Id: rid,
	}

	o := orm.NewOrm()
	_, err = o.Delete(reply)
	return err
}

//获取评论列表
func GetAllReplies(id string, isDesc bool) ([]*Comment, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	replies := make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	if isDesc {
		_, err = qs.Filter("Tid", tid).OrderBy("-created").All(&replies)
	} else {
		_, err = qs.Filter("Tid", tid).All(&replies)
	}

	return replies, err
}
