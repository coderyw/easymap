/**
 * @Author: yinwei
 * @Description:
 * @File: decode
 * @Version: 1.0.0
 * @Date: 2022/10/29 14:30
 */
package gen

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode"
	"unicode/utf8"
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

	stBuf := new(bytes.Buffer)
	fmt.Fprintln(stBuf, fmt.Sprintf("\tvar ("))
	fmt.Fprintln(stBuf, fmt.Sprintf("\t\tok bool"))
	fmt.Fprintln(stBuf, fmt.Sprintf("\t\tval string"))
	fmt.Fprintln(stBuf, fmt.Sprintf("\t)"))

	var (
		field reflect.StructField
		js    string
		err   error
		got   bool = false
	)

	for i := 0; i < r.NumField(); i++ {
		field = r.Field(i)
		if !isExportedOrBuiltinType(field.Type) {
			continue
		}
		js = field.Tag.Get("json")
		if js != "" {
			fmt.Fprintln(stBuf, fmt.Sprintf("\tif val,ok=m[\"%v\"];ok{", js))
			if err = g.decodeField(stBuf, field); err != nil {
				break
			}
			got = true
			fmt.Fprintln(stBuf, "\t}")
		}
	}
	if err == nil && got {
		fmt.Fprintln(g.out, stBuf.String())
	}

	fmt.Fprintln(g.out, "\treturn nil")
	fmt.Fprintln(g.out, "}")
	return err
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

func (g *generator) decodeField(out *bytes.Buffer, r reflect.StructField) error {
	switch r.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseInt(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Int:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int(pv)", r.Name))
		case reflect.Int8:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int8(pv)", r.Name))
		case reflect.Int16:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int16(pv)", r.Name))
		case reflect.Int32:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int32(pv)", r.Name))
		case reflect.Int64:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseUint(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Uint:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint(pv)", r.Name))
		case reflect.Uint8:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint8(pv)", r.Name))
		case reflect.Uint16:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint16(pv)", r.Name))
		case reflect.Uint32:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint32(pv)", r.Name))
		case reflect.Uint64:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Float32, reflect.Float64:
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseFloat(val, 10); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch r.Type.Kind() {
		case reflect.Float32:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = float32(pv)", r.Name))
		case reflect.Float64:
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.String:
		fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v = val", r.Name))
	case reflect.Bool:
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseBool(val); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", r.Name))
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))

	default:
		return fmt.Errorf("不支持的field,%v", r.Name)
	}
	return nil

}
