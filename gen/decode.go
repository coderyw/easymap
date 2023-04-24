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
	"strings"
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

	var (
		field reflect.StructField
		js    string
		err   error
	)

	var buf = new(bytes.Buffer)
	for i := 0; i < r.NumField(); i++ {
		field = r.Field(i)
		if !isExportedOrBuiltinType(field.Type) {
			continue
		}
		if strings.HasPrefix(field.Name, "XXX_") {
			continue
		}
		js = field.Tag.Get(tag)
		if js == "" {
			js = field.Name
		}
		if js != "" {
			if out, err := g.decodeField(field, field.Type, false); err != nil {
				break
			} else {
				if out == nil {
					continue
				}
				fmt.Fprintln(buf, fmt.Sprintf("\tif val,ok=m[\"%v\"];ok{", js))
				fmt.Fprint(buf, out)
				fmt.Fprintln(buf, "\t}")
			}

		}
	}

	if buf.Len() > 0 {
		stBuf := new(bytes.Buffer)
		fmt.Fprintln(stBuf, fmt.Sprintf("\tvar ("))
		fmt.Fprintln(stBuf, fmt.Sprintf("\t\tok bool"))
		fmt.Fprintln(stBuf, fmt.Sprintf("\t\tval string"))
		fmt.Fprintln(stBuf, fmt.Sprintf("\t)"))
		fmt.Fprintln(stBuf, buf)
		fmt.Fprintln(g.out, stBuf)
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

func (g *generator) decodeField(field reflect.StructField, t reflect.Type, isPtr bool) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseInt(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Int:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := int(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int(pv)", field.Name))
			}
		case reflect.Int8:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := int8(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int8(pv)", field.Name))
			}
		case reflect.Int16:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := int16(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int16(pv)", field.Name))
			}
		case reflect.Int32:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := int32(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = int32(pv)", field.Name))
			}
		case reflect.Int64:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", field.Name))
			}
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseUint(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Uint:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := uint(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint(pv)", field.Name))
			}
		case reflect.Uint8:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := uint8(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint8(pv)", field.Name))
			}
		case reflect.Uint16:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := uint16(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint16(pv)", field.Name))
			}
		case reflect.Uint32:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := uint32(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = uint32(pv)", field.Name))
			}
		case reflect.Uint64:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", field.Name))
			}
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Float32, reflect.Float64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseFloat(val, 10); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Float32:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := float32(pv)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = float32(pv)", field.Name))
			}
		case reflect.Float64:
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", field.Name))
			}
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.String:
		if isPtr {
			fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v = &val", field.Name))
		} else {
			fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v = val", field.Name))
		}
	case reflect.Bool:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseBool(val); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		if isPtr {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pv", field.Name))
		} else {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = pv", field.Name))
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t}"))

	case reflect.Ptr:
		return g.decodeField(field, t.Elem(), true)
	//case reflect.Interface:

	default:
		return nil, nil
	}
	return out, nil

}
