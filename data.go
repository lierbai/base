package base

import (
	"encoding/xml"
)

// Data map[string]interface{}的快捷方式
type Data map[string]interface{}

// MarshalXML 允许类型D用户xml.Marshal.
func (d Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range d {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Set set
func (d Data) set(key string, Obj interface{}) {
	d[key] = Obj
}

// Err Data的另一种封装
func Err(code int, msg string, jsonObj interface{}) Data {
	data := Data{
		"code": code,
		"msg":  msg,
	}
	if jsonObj != nil {
		data.set("data", jsonObj)
	}
	return data
}
