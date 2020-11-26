package binding

import (
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
)

type multipartRequest http.Request

var _ setter = (*multipartRequest)(nil)

// TrySet 尝试通过multipart请求来设置值,并绑定Form文件
func (r *multipartRequest) TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (isSetted bool, err error) {
	if files := r.MultipartForm.File[key]; len(files) != 0 {
		return setByMultipartFormFile(value, field, files)
	}

	return setByForm(value, field, r.MultipartForm.Value, key, opt)
}

func setByMultipartFormFile(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSetted bool, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		switch value.Interface().(type) {
		case *multipart.FileHeader:
			value.Set(reflect.ValueOf(files[0]))
			return true, nil
		}
	case reflect.Struct:
		switch value.Interface().(type) {
		case multipart.FileHeader:
			value.Set(reflect.ValueOf(*files[0]))
			return true, nil
		}
	case reflect.Slice:
		slice := reflect.MakeSlice(value.Type(), len(files), len(files))
		isSetted, err = setArrayOfMultipartFormFiles(slice, field, files)
		if err != nil || !isSetted {
			return isSetted, err
		}
		value.Set(slice)
		return true, nil
	case reflect.Array:
		return setArrayOfMultipartFormFiles(value, field, files)
	}
	// 不支持的字段类型(异常抛出)
	return false, errors.New("unsupported field type for multipart.FileHeader")
}

func setArrayOfMultipartFormFiles(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSetted bool, err error) {
	if value.Len() != len(files) {
		// 不支持的数组长度(异常抛出)
		return false, errors.New("unsupported len of array for []*multipart.FileHeader")
	}
	for i := range files {
		setted, err := setByMultipartFormFile(value.Index(i), field, files[i:i+1])
		if err != nil || !setted {
			return setted, err
		}
	}
	return true, nil
}
