package apgs

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"reflect"
)

// 通用MAP
type Map = map[string]interface{}

// Map 格式返回接口信息
func ApiReturn(code int, msg string, data interface{}) *iris.Map {
	return &iris.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

// Map 格式返回状态 消息 跳转连接
func ApiRedirect(code int, msg string, redirectUrl string) *iris.Map {
	return &iris.Map{
		"code":    code,
		"message": msg,
		"url":     redirectUrl,
	}
}

// ToMap 结构体转为Map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
