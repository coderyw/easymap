/**
 * @Author: yinwei
 * @Description:
 * @File: decode
 * @Version: 1.0.0
 * @Date: 2022/10/29 14:30
 */
package gen

import (
	"fmt"
	"reflect"
)

func (g *generator) decode(t reflect.Type) error {
	if t.Kind() == reflect.Struct {
		return g.decodeStruct(t)
	} else {
		return fmt.Errorf("暂时只支持struct")
	}
}

func (g *generator) decodeStruct(r reflect.Type) error {
	fmt.Fprintln(g.out, fmt.Sprintf("func (v *%v) UnMarshalMap(m map[string]string) error{", r.Name()))
	fmt.Fprintln(g.out, fmt.Sprintf("\tvar ("))
	fmt.Fprintln(g.out, fmt.Sprintf("\t\tok bool"))
	fmt.Fprintln(g.out, fmt.Sprintf("\t\tval string"))
	fmt.Fprintln(g.out, fmt.Sprintf("\t)"))

	var (
		field reflect.StructField
		js    string
		err   error
	)
	for i := 0; i < r.NumField(); i++ {
		field = r.Field(i)
		js = field.Tag.Get("json")
		if js != "" {
			fmt.Fprintln(g.out, fmt.Sprintf("\tif val,ok=m[\"%v\"];ok{", js))
			if err = g.decodeField(field); err != nil {
				return err
			}
			fmt.Fprintln(g.out, "\t}")
		}
	}

	fmt.Fprintln(g.out, "\treturn nil")
	fmt.Fprintln(g.out, "}")
	return nil
}

func (g *generator) decodeField(r reflect.StructField) error {
	switch r.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintln(g.out, fmt.Sprintf("\t\tif pv, err := strconv.ParseInt(val, 10, 64); err != nil {"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Int:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = int(pv)", r.Name))
		case reflect.Int8:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = int8(pv)", r.Name))
		case reflect.Int16:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = int16(pv)", r.Name))
		case reflect.Int32:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = int32(pv)", r.Name))
		case reflect.Int64:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t}"))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintln(g.out, fmt.Sprintf("\t\tif pv, err := strconv.ParseUint(val, 10, 64); err != nil {"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Uint:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = uint(pv)", r.Name))
		case reflect.Uint8:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = uint8(pv)", r.Name))
		case reflect.Uint16:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = uint16(pv)", r.Name))
		case reflect.Uint32:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = uint32(pv)", r.Name))
		case reflect.Uint64:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t}"))
	case reflect.Float32, reflect.Float64:
		fmt.Fprintln(g.out, fmt.Sprintf("\t\tif pv, err := strconv.ParseFloat(val, 10); err != nil {"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Float32:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = float32(pv)", r.Name))
		case reflect.Float64:
			fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t}"))
	case reflect.String:
		fmt.Fprintln(g.out, fmt.Sprintf("\t\tv.%v = val", r.Name))
	case reflect.Bool:
		fmt.Fprintln(g.out, fmt.Sprintf("\t\tif pv, err := strconv.ParseBool(val); err != nil {"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t} else {"))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		fmt.Fprintln(g.out, fmt.Sprintf("\t\t}"))

	default:
		return fmt.Errorf("不支持的field,%v", r.Name)
	}
	return nil

}
