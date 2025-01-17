// Package gen
// @Author: yinwei
// @File: unmarshal_mapinterface
// @Version: 1.0.0
// @Date: 2025/1/17 18:14

package gen

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

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
			if out, err := g.decodeInterField(field.Name, field.Type, false, r.PkgPath()); err != nil {
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

func (g *generator) decodeInterField(fieldName string, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	turnStr := ""
	var tp reflect.Type = t
	if t.Kind() == reflect.Ptr {
		return g.decodeInterField(fieldName, t.Elem(), true, pkgPath)
	} else if t.Kind() == reflect.Struct && !(tp.PkgPath() == "github.com/shopspring/decimal" && tp.Name() == "Decimal") {
		return g.decodeStructField1(fieldName, t, false, pkgPath)
	} else if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		//return nil, nil
		return g.decodeArrayField(fieldName, t, isPtr, pkgPath)
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
		kindInterArr["string"] = struct{}{}
		for v := range kindInterArr {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			switch v {
			case "string":
				g.imports[pkgStrconv] = "strconv"

				fmt.Fprintln(out, fmt.Sprintf("\t\t\t if len(acval)==0{ v.%v = 0", fieldName))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t }else { pvv, err := strconv.ParseInt(acval, 10, 64)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tif err != nil {return err}"))
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\td := %v(pvv)", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &d", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(pvv)}", fieldName, turnStr))
				}
			default:
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", fieldName, turnStr))
				}
			}
		}

	case reflect.Float32, reflect.Float64:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindFloatArr[turnStr] = struct{}{}
		kindFloatArr["float32"] = struct{}{}
		kindFloatArr["float64"] = struct{}{}
		kindFloatArr["int"] = struct{}{}
		kindFloatArr["int8"] = struct{}{}
		kindFloatArr["int16"] = struct{}{}
		kindFloatArr["int32"] = struct{}{}
		kindFloatArr["int64"] = struct{}{}
		kindFloatArr["uint"] = struct{}{}
		kindFloatArr["uint8"] = struct{}{}
		kindFloatArr["uint16"] = struct{}{}
		kindFloatArr["uint32"] = struct{}{}
		kindFloatArr["uint64"] = struct{}{}
		kindFloatArr["string"] = struct{}{}
		for v := range kindFloatArr {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			switch v {
			case "float32", "float64", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := %v(acval)", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", fieldName, turnStr))
				}
			case "string":
				g.imports[pkgStrconv] = "strconv"
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t if len(acval)==0{ v.%v = 0", fieldName))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t }else {pvv, err := strconv.ParseFloat(acval, 10)"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t if err != nil {return err}"))
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\ttmp:= %v(pvv)", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(pvv)}", fieldName, turnStr))
				}
			}

		}

	case reflect.String:
		if turnStr == "" {
			turnStr = t.Kind().String()
		}
		kindString[turnStr] = struct{}{}
		kindString["int"] = struct{}{}
		kindString["int8"] = struct{}{}
		kindString["int16"] = struct{}{}
		kindString["int32"] = struct{}{}
		kindString["int64"] = struct{}{}
		kindString["uint"] = struct{}{}
		kindString["uint8"] = struct{}{}
		kindString["uint16"] = struct{}{}
		kindString["uint32"] = struct{}{}
		kindString["uint64"] = struct{}{}
		kindString["float64"] = struct{}{}
		kindString["float32"] = struct{}{}
		for v := range kindString {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			switch v {
			case "string":
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := acval"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = acval", fieldName))
				}
			case "int", "int8", "int16", "int32", "int64":
				g.imports[pkgStrconv] = "strconv"
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := strconv.FormatInt(int64(acval), 10)"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = strconv.FormatInt(int64(acval), 10)", fieldName))
				}
			case "uint", "uint8", "uint16", "uint32", "uint64":
				g.imports[pkgStrconv] = "strconv"
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := strconv.FormatUint(uint64(acval), 10)"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = strconv.FormatUint(uint64(acval), 10)", fieldName))
				}
			case "float64", "float32":
				g.imports[pkgStrconv] = "strconv"
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t pvv := strconv.FormatFloat(float64(acval), 'f',-1,64)"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = strconv.FormatFloat(float64(acval),  'f',-1,64)", fieldName))
				}
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
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\tv.%v = &pvv", fieldName))
			} else {
				fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = %v(acval)", fieldName, turnStr))
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
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", fieldName))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = decimal.NewFromInt(int64(acval))", fieldName))
						}
					case 1:
						if isPtr {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar decc decimal.Decimal"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tdecc = decimal.NewFromFloat(float64(acval))"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", fieldName))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = decimal.NewFromFloat(float64(acval))", fieldName))
						}
					case 2:
						if isPtr {
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar err error"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tvar decc decimal.Decimal"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tdecc,err = decimal.NewFromString(acval)"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tif err!= nil {"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\t\treturn err"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\t}"))
							fmt.Fprintln(out, fmt.Sprintf("\t\t\tv.%v = &decc", fieldName))
						} else {
							fmt.Fprintln(out, fmt.Sprintf("\t\tvar err error"))
							fmt.Fprintln(out, fmt.Sprintf("\t\tv.%v,err = decimal.NewFromString(acval)", fieldName))
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

func (g *generator) decodeArrayField(fieldName string, t reflect.Type, isPtr bool, pkgPath string) (*bytes.Buffer, error) {
	///注意 t必须是数组，且没有elem
	out := new(bytes.Buffer)
	var arrFieldIsPtr bool
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		arrFieldIsPtr = true
		t = t.Elem()
	}
	kindString := make(map[string]struct{})
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String, reflect.Struct:

	default:
		return out, nil

	}
	turnStr := t.Kind().String()
	if t.Kind() == reflect.Struct && t.Name() == "" {
		return out, nil
	}
	trueKind := "[]" + t.Kind().String()
	fmt.Fprintln(out, fmt.Sprintf("\t\t switch acval := val.(type) {"))
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		kindString["[]int"] = struct{}{}
		kindString["[]int8"] = struct{}{}
		kindString["[]int16"] = struct{}{}
		kindString["[]int32"] = struct{}{}
		kindString["[]int64"] = struct{}{}
		kindString["[]uint"] = struct{}{}
		kindString["[]uint8"] = struct{}{}
		kindString["[]uint16"] = struct{}{}
		kindString["[]uint32"] = struct{}{}
		kindString["[]uint64"] = struct{}{}
		kindString["[]string"] = struct{}{}
		kindString["[]interface{}"] = struct{}{}
		for v := range kindString {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if trueKind == v {
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &acval", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = acval", fieldName))
				}
			} else {

				switch v {
				case "[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmpArr[i] = %v(k)}", turnStr))
					if isPtr {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
					} else {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
					}
				case "[]string":
					g.imports[pkgStrconv] = "strconv"
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
					switch t.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmp,err := strconv.ParseInt(k, 10, 64)"))
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmp,err := strconv.ParseUint(k, 10, 64)"))
					}
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if err!= nil{return err}"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmpArr[i] = %v(tmp)}", turnStr))
					if isPtr {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
					} else {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
					}
				case "[]interface{}":

				}
			}
		}
	case reflect.Float32, reflect.Float64:
		kindString["[]float32"] = struct{}{}
		kindString["[]float64"] = struct{}{}
		kindString["[]int"] = struct{}{}
		kindString["[]int8"] = struct{}{}
		kindString["[]int16"] = struct{}{}
		kindString["[]int32"] = struct{}{}
		kindString["[]int64"] = struct{}{}
		kindString["[]uint"] = struct{}{}
		kindString["[]uint8"] = struct{}{}
		kindString["[]uint16"] = struct{}{}
		kindString["[]uint32"] = struct{}{}
		kindString["[]uint64"] = struct{}{}
		kindString["[]string"] = struct{}{}
		for v := range kindString {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))
			if trueKind == v {
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &acval", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = acval", fieldName))
				}
			} else {
				switch v {
				case "[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64", "[]float32", "[]float64":
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmpArr[i] = %v(k)}", turnStr))
					if isPtr {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
					} else {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
					}
				case "[]string":
					g.imports[pkgStrconv] = "strconv"
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))

					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmp,err := strconv.ParseFloat(k, 10)"))

					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if err!= nil{return err}"))
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t tmpArr[i] = %v(tmp)}", turnStr))
					if isPtr {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
					} else {
						fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
					}
				}
			}
		}
	case reflect.String:
		kindString["[]string"] = struct{}{}
		kindString["[]float64"] = struct{}{}
		kindString["[]int"] = struct{}{}
		kindString["[]int8"] = struct{}{}
		kindString["[]int16"] = struct{}{}
		kindString["[]int32"] = struct{}{}
		kindString["[]int64"] = struct{}{}
		kindString["[]uint"] = struct{}{}
		kindString["[]uint8"] = struct{}{}
		kindString["[]uint16"] = struct{}{}
		kindString["[]uint32"] = struct{}{}
		kindString["[]uint64"] = struct{}{}
		kindString["[]string"] = struct{}{}
		for v := range kindString {
			fmt.Fprintln(out, fmt.Sprintf("\t\tcase %v:", v))

			switch v {
			case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
				g.imports[pkgStrconv] = "strconv"
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = strconv.FormatInt(int64(k), 10)}", fieldName))
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
				}
			case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
				g.imports[pkgStrconv] = "strconv"
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = strconv.FormatUint(uint64(k), 10 )}", fieldName))
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
				}
			case "[]float32", "[]float64":
				g.imports[pkgStrconv] = "strconv"
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t tmpArr := make([]%v, len(acval))", turnStr))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
				fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = strconv.FormatFloat(float64(k), 'f', -1,64 )}", fieldName))
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &tmpArr", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = tmpArr", fieldName))
				}
			case "[]string":
				if isPtr {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = &acval", fieldName))
				} else {
					fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = acval", fieldName))
				}
			}
		}
	case reflect.Struct:
		kindString["[]map[string]string"] = struct{}{}
		kindString["[]map[string]interface{}"] = struct{}{}
		g.imports[pkgFacade] = "github.com/coderyw/easymap/gen/facade"
		fmt.Fprintln(out, fmt.Sprintf("\t\tcase []map[string]string:"))
		if !arrFieldIsPtr {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = make([]%v, len(acval))", fieldName, t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t var kjj interface{} = %v{}", t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if b, ok := kjj.(easy_facade.EasyMapString); ok {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\t if err := b.UnMarshalMap(k); err != nil { return err }}"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = kjj.(%v)}", fieldName, t.Name()))
		} else {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = make([]*%v, len(acval))", fieldName, t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t var kjj interface{} = &%v{}", t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if b, ok := kjj.(easy_facade.EasyMapString); ok {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\t if err := b.UnMarshalMap(k); err != nil { return err }}"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = kjj.(*%v)}", fieldName, t.Name()))
		}
		fmt.Fprintln(out, fmt.Sprintf("\t\tcase []map[string]interface{}:"))
		if !arrFieldIsPtr {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = make([]%v, len(acval))", fieldName, t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t var kjj interface{} = %v{}", t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if b, ok := kjj.(easy_facade.EasyMapInter); ok {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\t if err := b.UnMarshalMapInterface(k); err != nil { return err }}"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = kjj.(%v)}", fieldName, t.Name()))
		} else {
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t v.%v = make([]*%v, len(acval))", fieldName, t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t for i, k := range acval {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t var kjj interface{} = &%v{}", t.Name()))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t if b, ok := kjj.(easy_facade.EasyMapInter); ok {"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t\t if err := b.UnMarshalMapInterface(k); err != nil { return err }}"))
			fmt.Fprintln(out, fmt.Sprintf("\t\t\t\t v.%v[i] = kjj.(*%v)}", fieldName, t.Name()))
		}
	}
	fmt.Fprintln(out, fmt.Sprintf("\t\t }"))
	return out, nil
}
