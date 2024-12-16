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

func (g *generator) decodeInterStruct(r reflect.Type) error {
	fmt.Fprintln(g.out, fmt.Sprintf("func (v *%v) UnMarshalMapInterface(m map[string]interface{}) error{", r.Name()))

	var (
		field reflect.StructField
		js    string
		err   error
	)

	var buf = new(bytes.Buffer)
	for i := 0; i < r.NumField(); i++ {
		field = r.Field(i)
		//if !isExportedOrBuiltinType(field.Type) {
		//	continue
		//}
		if strings.HasPrefix(field.Name, "XXX_") {
			continue
		}
		js = g.getTag(field)
		if js != "" {
			if out, err := g.decodeInterField(field, field.Type, false, r.PkgPath()); err != nil {
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
		fmt.Fprintln(stBuf, fmt.Sprintf("\t\tval interface{}"))
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

func (g *generator) decodeInterField(field reflect.StructField, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	turnStr := ""
	var tp reflect.Type = t
	if t.Kind() == reflect.Ptr {
		return g.decodeInterField(field, t.Elem(), true, pkgPath)
	} else if t.Kind() == reflect.Struct && !(tp.PkgPath() == "github.com/shopspring/decimal" && tp.Name() == "Decimal") {
		return g.decodeStructField1(field, t, false, pkgPath)
	} else if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		return nil, nil
		return g.decodeArrayField(field, t, isPtr, pkgPath)
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
	kindInterArr := map[string]struct{}{
		reflect.Int.String(): {}, reflect.Int8.String(): {}, reflect.Int16.String(): {}, reflect.Int32.String(): {}, reflect.Int64.String(): {},
		reflect.Uint.String(): {}, reflect.Uint8.String(): {}, reflect.Uint16.String(): {}, reflect.Uint32.String(): {}, reflect.Uint64.String(): {},
	}
	kindFloatArr := map[string]struct{}{
		reflect.Float32.String(): {}, reflect.Float64.String(): {},
	}
	kindString := map[string]struct{}{
		reflect.String.String(): {},
	}
	kindBool := map[string]struct{}{
		reflect.Bool.String(): {},
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\tswitch acval:=val.(type){"))

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindInterArr[turnStr] = struct{}{}
		for v := range kindInterArr {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", field.Name, turnStr))
			}
		}
	case reflect.Float32, reflect.Float64:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindFloatArr[turnStr] = struct{}{}
		for v := range kindFloatArr {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", field.Name, turnStr))
			}
		}
	case reflect.String:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindString[turnStr] = struct{}{}
		for v := range kindString {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", field.Name, turnStr))
			}
		}
	case reflect.Bool:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindString[turnStr] = struct{}{}
		for v := range kindBool {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if isPtr {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", field.Name))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", field.Name, turnStr))
			}
		}
	case reflect.Struct:
		if tp.PkgPath() == "github.com/shopspring/decimal" && tp.Name() == "Decimal" {
			g.imports[pkgDecimal] = "github.com/shopspring/decimal"
			ab := []map[string]struct{}{
				kindInterArr,
				kindFloatArr,
				kindString,
			}
			for i, v := range ab {
				for k := range v {
					fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", k))
					switch i {
					case 0:
						if isPtr {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar decc decimal.Decimal"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tdecc = decimal.NewFromInt(int64(acval))"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", field.Name))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = decimal.NewFromInt(int64(acval))", field.Name))
						}
					case 1:
						if isPtr {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar decc decimal.Decimal"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tdecc = decimal.NewFromFloat(float64(acval))"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", field.Name))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = decimal.NewFromFloat(float64(acval))", field.Name))
						}
					case 2:
						if isPtr {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar err error"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar decc decimal.Decimal"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tdecc,err = decimal.NewFromString(acval)"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tif err!= nil {"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\t\treturn err"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", field.Name))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\tvar err error"))
							fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v,err = decimal.NewFromString(acval)", field.Name))
							fmt.Fprintln(out, fmt.Sprintf("\t\tif err!= nil {"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\treturn err"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t}"))
						}
					}
				}
			}

		}
		//type:.eq.github.com/shopspring/decimal.Decimal

	default:
		delete(g.imports, add)
		return nil, nil
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

func (g *generator) decodeStructField1(field reflect.StructField, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}
	g.imports[pkgFacade] = "github.com/coderyw/easymap/gen/facade"

	fmt.Fprintln(out, fmt.Sprintf("\t\tvar i interface{}=v.%v", field.Name))
	fmt.Fprintln(out, fmt.Sprintf("\t\tif m1,ok:= val.(map[string]interface{}); ok {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\tif b,ok :=i.(%v.EasyMapInter); ok  {", pkgFacade))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tif err := b.UnMarshalMapInterface(m1); err != nil {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t}else if m2, ok := val.(map[string]string); ok {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\tif b, ok := i.(%v.EasyMapString); ok {", pkgFacade))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tif err := b.UnMarshalMap(m2); err != nil {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t}"))

	return out, nil
}

func (g *generator) decodeArrayField(field reflect.StructField, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	///注意 t必须是数组，且没有elem
	out := new(bytes.Buffer)
	var arrFieldIsPtr bool
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		arrFieldIsPtr = true
		t = t.Elem()
	}
	name := t.Name()
	fmt.Fprintln(out, fmt.Sprintf("\t\tif m1,ok:= val.([]map[string]interface{}); ok {"))
	if arrFieldIsPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tvv := make([]*%v, 0)", name))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tvv := make([]%v, 0)", name))
	}

	fmt.Fprintln(out, fmt.Sprintf("\t\t\tfor _, v1 := range m1 {"))
	if arrFieldIsPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tab := new(%v)", t.Name()))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tab := %v{}", t.Name()))
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tif err := ab.UnMarshalMapInterface(v1); err != nil {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tvv = append(vv, ab)"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	if isPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &vv", field.Name))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = vv", field.Name))
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\t} else if m1,ok:= val.([]map[string]string); ok {"))

	if arrFieldIsPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tvv := make([]*%v, 0)", t.Name()))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tvv := make([]%v, 0)", t.Name()))
	}

	fmt.Fprintln(out, fmt.Sprintf("\t\t\tfor _, v1 := range m1 {"))
	if arrFieldIsPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tab := new(%v)", t.Name()))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tab := %v{}", t.Name()))
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tif err := ab.UnMarshalMap(v1); err != nil {"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\treturn err"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t}"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tvv = append(vv, ab)"))
	fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
	if isPtr {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &vv", field.Name))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = vv", field.Name))
	}

	fmt.Fprintln(out, fmt.Sprintf("\t\t}"))

	return out, nil
}
