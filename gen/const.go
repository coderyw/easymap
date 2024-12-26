/**
 * @Author: yinwei
 * @Description:
 * @File: const
 * @Version: 1.0.0
 * @Date: 2022/10/30 16:13
 */
package gen

import (
	"fmt"
	"reflect"
)

const tag = "em"
const jsTag = "json"

func (g *generator) encodeConst(t reflect.Type) error {
	if t.Kind() == reflect.Struct {
		return g.encodeConstStruct(t)
	} else {
		return fmt.Errorf("暂时只支持struct")
	}
}

func (g *generator) encodeConstStruct(t reflect.Type) error {
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, fmt.Sprintf("type %vField string", t.Name()))
	fmt.Fprintln(g.out, fmt.Sprintf("func (v %vField) MarshalBinary() (data []byte, err error) {", t.Name()))
	fmt.Fprintln(g.out, fmt.Sprintf("return %s(string(v)), nil", g.str2BytesName))
	fmt.Fprintln(g.out, fmt.Sprint("}"))

	fmt.Fprintln(g.out, fmt.Sprintf("const("))
	fmt.Fprintln(g.out, fmt.Sprintf(""))
	for i := 0; i < t.NumField(); i++ {
		g.encodeConstField(t, t.Field(i), t.Field(i).Type)
	}
	fmt.Fprintln(g.out, fmt.Sprintf(")"))
	return nil
}

func (g *generator) encodeConstField(pt reflect.Type, field reflect.StructField, fv reflect.Type) {
	if fv.Kind() == reflect.Ptr {
		g.encodeConstField(pt, field, fv.Elem())
	} else {
		fmt.Fprintln(g.out, fmt.Sprintf("\t%v_%v %vField = \"%v\"", pt.Name(), field.Name, pt.Name(), g.getTag(field)))
	}
}

func (g *generator) genOther(t reflect.Type) error {
	if t.Kind() != reflect.Struct {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			return g.genOther(t)
		}
		return fmt.Errorf("暂时只支持struct")
	}
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, fmt.Sprintf("func (v *%v) CheckField(m map[string]string) error {", t.Name()))

	for i := 0; i < t.NumField(); i++ {
		tag := g.getTag(t.Field(i))

		fmt.Fprintln(g.out, fmt.Sprintf(`	if _,ok := m["%v"]; !ok {`, tag))
		fmt.Fprintln(g.out, fmt.Sprintf(`		return fmt.Errorf("field %v is miss")`, tag))
		fmt.Fprintln(g.out, "	}")
	}

	fmt.Fprintln(g.out, fmt.Sprintf("\t return nil"))
	fmt.Fprintln(g.out, fmt.Sprintf("}"))

	return nil
}
