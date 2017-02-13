package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

//数据结构
type Device struct {
	Id        int64
	DeviceID  string `orm:"index"`
	Name      string
	IPAddress string
	LastTime  time.Time
}

//更新设备信息
func UpdateDevice(id, name, ip string) error {
	o := orm.NewOrm()
	dev := &Device{
		DeviceID:  id,
		Name:      name,
		IPAddress: ip,
		LastTime:  time.Now(),
	}

	qs := o.QueryTable("device")
	err := qs.Filter("DeviceID", id).One(dev)
	if err == nil {
		log.Println("该设备已经存在,进行更新操作")
		_, err = o.Update(dev)
		return err
	}
	log.Println("该设备不存在,进行新增操作")
	_, err = o.Insert(dev)
	return err
}

//查找文章列表
func GetAllDevice() ([]*Device, error) {
	o := orm.NewOrm()
	devs := make([]*Device, 0)

	qs := o.QueryTable("device")

	_, err := qs.OrderBy("-lastTime").All(&devs)

	return devs, err
}
