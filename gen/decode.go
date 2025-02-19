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
		if err := g.decodeStruct(t); err != nil {
			return err
		}
		return g.decodeInterStruct(t)
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
		js = g.getTag(field)
		if js != "" {
			if out, err := g.decodeField(field, field.Type, false, r.PkgPath()); err != nil {
				continue
			} else {
				if out == nil || out.Len() == 0 {
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
	//fmt.Println(t.PkgPath())
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

func (g *generator) decodeField(field reflect.StructField, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	turnStr := ""
	var tp reflect.Type = t
	if t.Kind() == reflect.Ptr {
		return g.decodeField(field, t.Elem(), true, pkgPath)
	}
	add := ""
	if tp.String() != tp.Kind().String() {
		path := tp.PkgPath()
		arr := strings.Split(tp.String(), ".")
		if len(arr) < 2 {
			return nil, fmt.Errorf("错误的名称 %v", arr)
		}
		if path != pkgPath { //需要import
			add = arr[0]
			g.imports[arr[0]] = tp.PkgPath()
			turnStr = tp.String()
		} else {
			turnStr = arr[1]
		}

	}

	switch t.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseInt(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Int:
			if turnStr == "" {
				turnStr = "int"
			}
		case reflect.Int8:
			if turnStr == "" {
				turnStr = "int8"
			}
		case reflect.Int16:
			if turnStr == "" {
				turnStr = "int16"
			}
		case reflect.Int32:
			if turnStr == "" {
				turnStr = "int32"
			}
		case reflect.Int64:
			if turnStr == "" {
				turnStr = "int64"
			}
		}
		//fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseUint(val, 10, 64); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Uint:
			if turnStr == "" {
				turnStr = "uint"
			}
		case reflect.Uint8:
			if turnStr == "" {
				turnStr = "uint8"
			}
		case reflect.Uint16:
			if turnStr == "" {
				turnStr = "uint16"
			}
		case reflect.Uint32:
			if turnStr == "" {
				turnStr = "uint32"
			}
		case reflect.Uint64:
			if turnStr == "" {
				turnStr = "uint64"
			}
		}
		//fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	case reflect.Float32, reflect.Float64:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseFloat(val, 10); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		switch t.Kind() {
		case reflect.Float32:
			if turnStr == "" {
				turnStr = "float32"
			}
		case reflect.Float64:
			if turnStr == "" {
				turnStr = "float64"
			}
		}
	case reflect.String:
		if turnStr == "" {
			turnStr = "string"
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\t{pv := val"))
	case reflect.Bool:
		g.imports[pkgStrconv] = "strconv"
		fmt.Fprintln(out, fmt.Sprintf("\t\tif pv, err := strconv.ParseBool(val); err != nil {"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
		fmt.Fprintln(out, fmt.Sprintf("\t\t} else {"))
		if turnStr == "" {
			turnStr = "bool"
		}
	case reflect.Ptr:
		return g.decodeField(field, t.Elem(), true, pkgPath)
	case reflect.Struct:

		if tp.PkgPath() == "github.com/shopspring/decimal" && tp.Name() == "Decimal" {
			g.imports[pkgDecimal] = "github.com/shopspring/decimal"
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\tvar err error"))
				fmt.Fprintln(out, fmt.Sprintf("\t\tvar decc decimal.Decimal"))
				fmt.Fprintln(out, fmt.Sprintf("\t\tdecc,err = decimal.NewFromString(val)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\tif err!= nil {"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
				fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v = &decc", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\tvar err error"))
				fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v,err = decimal.NewFromString(val)", field.Name))
				fmt.Fprintln(out, fmt.Sprintf("\t\tif err!= nil {"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
			}
		} else {
			delete(g.imports, add)
		}
		//type:.eq.github.com/shopspring/decimal.Decimal

		return out, nil
	default:
		delete(g.imports, add)
		return nil, nil
	}
	if isPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %s(pv)", turnStr))
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &pvv", field.Name))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %s(pv)", field.Name, turnStr))
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
	return out, nil

}

func (g *generator) decodeStructField(field reflect.StructField, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\tif m1,ok:= val.(map[string]interface{}); ok {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\tif err := v.%v.UnMarshalMapInterface(m1); err != nil {", field.Name))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t}else if m2, ok := val.(map[string]string); ok {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\tif err := v.%v.UnMarshalMap(m2); err != nil {", field.Name))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t}"))

	return out, nil
}

func (g *generator) decodeStructField1(fieldName string, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}
	g.imports[pkgFacade] = "github.com/coderyw/easymap/gen/facade"

	var name string
	if pkgPath == t.PkgPath() {
		name = t.Name()
	} else {
		modelName := strings.ReplaceAll(strings.ReplaceAll(t.PkgPath(), "/", "_"), ".", "_")
		name = modelName + "." + t.Name()
		g.imports[modelName] = t.PkgPath()
	}
	if !isPtr {
		fmt.Fprintln(out,
			fmt.Sprintf(`
		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &%v{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.%v = *i.(*%v)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &%v{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.%v = *i.(*%v)
				}
			}
		case %v:
			v.%v = acval
		case *%v:
			if acval == nil {
				break
			}
			v.%v = *acval
		}

`, name, fieldName, name, name, fieldName, name, name, fieldName, name, fieldName))
	} else {
		fmt.Fprintln(out,
			fmt.Sprintf(`
		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &%v{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.%v = i.(*%v)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &%v{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.%v = i.(*%v)
				}
			}
		case %v:
			v.%v = &acval
		case *%v:
			v.%v = acval
		}

`, name, fieldName, name, name, fieldName, name, name, fieldName, name, fieldName))
	}
	return out, nil
}
