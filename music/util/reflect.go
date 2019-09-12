package util

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

//结构体转XML
func StructToMapXML(s interface{}) map[string]interface{} {
	return StructToMap(s, "xml")
}

//结构体转JSON
func StructToMapJSON(s interface{}) map[string]interface{} {
	return StructToMap(s, "json")
}

func StructToMap(s interface{}, tag string) map[string]interface{} {
	params := make(map[string]interface{})
	v := reflect.Indirect(reflect.ValueOf(s))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		isUnexported := ft.PkgPath != ""
		if isUnexported {
			continue
		}
		fv := v.Field(i)
		switch ft.Type.Kind() {
		case reflect.Struct:
			for k, v := range StructToMapXML(fv.Interface()) {
				params[k] = v
			}
		case reflect.String:
			tags := strings.Split(ft.Tag.Get(tag), ",")
			if len(tags) > 0 && tags[0] != "-" {
				params[tags[0]] = v.Field(i).String()
			}
		case reflect.Int:
			tags := strings.Split(ft.Tag.Get(tag), ",")
			if len(tags) > 0 && tags[0] != "-" {
				params[tags[0]] = int(v.Field(i).Int())
			}
		default:
			panic(fmt.Sprintf("invalid type \"%s\" of field \"%s\" in struct \"%s\".", ft.Type.Kind(), ft.Name, t.Kind()))
		}
	}
	return params
}

func Copy(dst interface{}, src interface{}) (err error) {
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		err = errors.New("dst isn't a pointer to struct")
		return
	}
	dstElem := dstValue.Elem()
	if dstElem.Kind() != reflect.Struct {
		err = errors.New("pointer doesn't point to struct")
		return
	}

	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		err = errors.New("src isn't struct")
		return
	}

	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		sv := srcValue.FieldByName(sf.Name)
		if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
			dv.Set(sv)
		}
	}
	return
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
