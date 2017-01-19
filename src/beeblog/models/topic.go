package models

import (
	"strconv"
	"time"

	"strings"

	"os"
	"path"

	"github.com/astaxie/beego/orm"
)

//数据结构
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Tags            string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

//添加文章
func AddTopic(title, content, category, tag, attachment string) error {
	tag = "$" + strings.Join(strings.Split(tag, " "), "#$") + "#"

	o := orm.NewOrm()

	topic := Topic{
		Title:      title,
		Content:    content,
		Tags:       tag,
		Category:   category,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
		Attachment: attachment,
	}

	_, err := o.Insert(&topic)
	return err
}

//删除文章
func DeleteTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}

	_, err = o.Delete(topic)
	return err
}

//修改文章
func ModifyTopic(id, title, content, category, tag, attachment string) error {
	tag = "$" + strings.Join(strings.Split(tag, " "), "#$") + "#"

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}

	var oldAttach string

	if o.Read(topic) == nil {
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Tags = tag
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		topic.Attachment = attachment
		o.Update(topic)
	}

	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	return nil
}

//修改文章回复信息
func ModifyTopicReply(id string, add bool, time time.Time) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}

	if o.Read(topic) == nil {
		if add {
			topic.ReplyCount++
			topic.ReplyTime = time
		} else {
			topic.ReplyCount--
			categories, err := GetAllReplies(id, true)
			if err != nil {
				return err
			} else {
				if len(categories) > 0 {
					topic.ReplyTime = categories[0].Created
				}
			}
		}

		o.Update(topic)
	}

	return nil
}

//查找文章
func GetTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}

	qs := o.QueryTable("topic")
	err = qs.Filter("Id", tid).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++

	_, err = o.Update(topic)

	topic.Tags = strings.Replace(strings.Replace(topic.Tags, "#", " ", -1), "$", "", -1)
	return topic, err
}

//查找文章列表
func GetAllTopics(isDesc bool, cate, tag string) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("Category", cate)
		}

		if len(tag) > 0 {
			qs = qs.Filter("tags__contains", "$"+tag+"#")
		}

		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}
