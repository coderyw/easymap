/**
 * @Author: yinwei
 * @Description:
 * @File: encode
 * @Version: 1.0.0
 * @Date: 2022/10/29 14:31
 */
package gen

import (
	"fmt"
	"reflect"
)

func (g *generator) encode(t reflect.Type) error {
	if t.Kind() == reflect.Struct {
		return g.encodeStruct(t)
	} else {
		return fmt.Errorf("暂时只支持struct")
	}
}

func (g *generator) encodeStruct(t reflect.Type) error {
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, fmt.Sprintf("func (v *%v) MarshalMap() (map[string]interface{}, error) {", t.Name()))
	fmt.Fprintln(g.out, fmt.Sprintf("\tm := make(map[string]interface{})"))
	var (
		field reflect.StructField
		fv    reflect.Type
	)
	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		fv = field.Type
		g.encodeField(fv, field, false)
	}
	fmt.Fprintln(g.out, fmt.Sprintf("\treturn m, nil"))
	fmt.Fprintln(g.out, fmt.Sprintf("}"))
	return nil
}

func (g *generator) getTag(s reflect.StructField) string {
	tag := s.Tag.Get(tag)
	if tag == "" {
		tag = s.Name
	}
	return tag
}

func (g *generator) encodeField(fv reflect.Type, field reflect.StructField, isPtr bool) {
	if (fv.Kind() < reflect.Complex64 && fv.Kind() > reflect.Invalid) || fv.Kind() == reflect.String {
		if isPtr {
			fmt.Fprintln(g.out, fmt.Sprintf("\tif v.%v != nil {", field.Name))
			fmt.Fprintln(g.out, fmt.Sprintf("\t\tm[\"%v\"] = *v.%v", g.getTag(field), field.Name))
			fmt.Fprintln(g.out, fmt.Sprintf("\t}"))
		} else {
			fmt.Fprintln(g.out, fmt.Sprintf("\tm[\"%v\"] = v.%v", g.getTag(field), field.Name))
		}

	} else if fv.Kind() == reflect.Ptr {
		fv = fv.Elem()
		g.encodeField(fv, field, true)
	}
}
