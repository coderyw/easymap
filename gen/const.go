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
	} else if (fv.Kind() < reflect.Complex64 && fv.Kind() > reflect.Invalid) || fv.Kind() == reflect.String {
		fmt.Fprintln(g.out, fmt.Sprintf("\t%v_%v = \"%v\"", pt.Name(), field.Name, g.getTag(field)))
	}
}
