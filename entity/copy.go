package entity

import (
	"reflect"
	"time"
)

func Copy(src interface{}, dstType interface{}) interface{} {
	if src == nil {
		return nil
	}
	cpy := reflect.New(reflect.TypeOf(dstType)).Elem()
	copyRecursive(reflect.ValueOf(src), cpy)
	return cpy.Interface()
}

func copyRecursive(src, dst reflect.Value) {
	switch src.Kind() {
	case reflect.Ptr:
		originValue := src.Elem()
		if !originValue.IsValid() {
			return
		}
		// 允许 src 为 ptr 而 dst 为非 ptr
		if dst.Kind() == reflect.Ptr {
			dst.Set(reflect.New(dst.Type().Elem()))
			copyRecursive(originValue, dst.Elem())
		} else {
			dst.Set(reflect.New(dst.Type()).Elem())
			copyRecursive(originValue, dst)
		}
	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originValue := src.Elem()
		copyValue := reflect.New(dst.Type().Elem()).Elem()
		copyRecursive(originValue, copyValue)
		dst.Set(copyValue)
	case reflect.Struct:
		// time.Time 需要特殊处理
		t, ok := src.Interface().(time.Time)
		if ok {
			dst.Set(reflect.ValueOf(t))
			return
		}
		if dst.Kind() == reflect.Ptr {
			// 目标类型是指针而源类型不是指针
			copyValue := reflect.New(dst.Type().Elem()).Elem()
			copyRecursive(src, copyValue)
			dst.Set(copyValue.Addr())
			return
		}
		for i := 0; i < dst.NumField(); i++ {
			if dst.Type().Field(i).PkgPath != "" {
				// 不可导出的字段不拷贝
				continue
			}
			field := src.FieldByName(dst.Type().Field(i).Name)
			if !field.IsValid() {
				// 源字段不存在，忽略（目标自动零值）
				continue
			}
			copyRecursive(field, dst.Field(i))
		}
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dst.Index(i))
		}
	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(dst.Type()))
		for _, key := range src.MapKeys() {
			value := src.MapIndex(key)
			copyValue := reflect.New(dst.Type().Elem()).Elem()
			copyRecursive(value, copyValue)
			copyKey := Copy(key.Interface(), reflect.New(dst.Type().Key()).Elem().Interface())
			dst.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}
	default:
		// 源类型是基础类型
		// 类型不一致但底层类型一致的基本类型，需要强转
		if dst.Kind() == reflect.Ptr {
			// 目标类型是指针而源类型不是指针
			copyValue := reflect.New(dst.Type().Elem()).Elem()
			copyRecursive(src, copyValue)
			dst.Set(copyValue.Addr())
			return
		}
		dst.Set(src.Convert(dst.Type()))
	}
}
