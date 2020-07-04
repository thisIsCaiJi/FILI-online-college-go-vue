package model

import (
	"errors"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
	"reflect"
	"time"
)

// 公共创建方法，当id为空时会生成id
func Create(m interface{}) error {
	id,err := util.GetId()
	if err!=nil {
		return err
	}
	t := reflect.ValueOf(time.Now())
	v := reflect.ValueOf(m)
	v = v.Elem()
	gmtCreate := v.FieldByName("GmtCreate")
	gmtModified := v.FieldByName("GmtModified")
	idv := v.FieldByName("Id")
	if gmtCreate.CanSet() {
		gmtCreate.Set(t)
	}
	if gmtModified.CanSet() {
		gmtModified.Set(t)
	}
	if idv.String() == ""{
		if idv.CanSet(){
			idv.SetString(id)
		}
	}
	return db.Create(m).Error
}

func ListPage(m interface{},current,limit int,list interface{}) (total uint,err error) {
	err = db.Model(m).Where(m).Count(&total).Limit(limit).Offset((current-1) * limit).Find(list).Error
	return total,err
}

func List(m interface{}, list interface{}) (total uint,err error) {
	err = db.Model(m).Where(m).Count(&total).Find(list).Find(list).Error
	return total,err
}

func Update(m interface{}) error {
	t := reflect.ValueOf(time.Now())
	v := reflect.ValueOf(m)
	v = v.Elem()
	gmtModified := v.FieldByName("GmtModified")
	if gmtModified.CanSet() {
		gmtModified.Set(t)
	}
	return db.Save(m).Error
}

func Remove(m interface{}) error {
	return db.Delete(m).Error
}

func One(m interface{},one interface{}) error {
	if db.Where(m).First(one).RecordNotFound() {
		return errors.New("resource is not found")
	}
	return nil
}

