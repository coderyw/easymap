/**
 * @Author: yinwei
 * @Description:
 * @File: generator
 * @Version: 1.0.0
 * @Date: 2022/10/29 12:23
 */
package gen

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

const pkgStrconv = "strconv"

type generator struct {
	pkgName     string
	pkgPath     string
	fileName    string
	typesUnseen []reflect.Type
	out         *bytes.Buffer
	imports     map[string]string
}

func NewGenerator(filename string) *generator {
	ret := &generator{
		fileName: filename,
		imports: map[string]string{
			pkgStrconv: "strconv",
		},
	}
	return ret
}

func (g *generator) SetPkg(name, path string) {
	g.pkgName = name
	g.pkgPath = path
}

func (g *generator) Add(obj interface{}) {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	g.typesUnseen = append(g.typesUnseen, t)
}

func (g *generator) Run(out io.Writer) error {
	g.out = new(bytes.Buffer)
	var err error
	for len(g.typesUnseen) > 0 {
		t := g.typesUnseen[len(g.typesUnseen)-1]
		g.typesUnseen = g.typesUnseen[:len(g.typesUnseen)-1]
		if err = g.decode(t); err != nil {
			return err
		}
		if err = g.encode(t); err != nil {
			return err
		}
	}
	//fmt.Println(g.out.String())
	g.writeImports(out)
	out.Write(g.out.Bytes())

	return nil
}

func (g *generator) writeImports(out io.Writer) {
	out.Write([]byte(fmt.Sprintf("package %v\n", g.pkgName)))
	out.Write([]byte("import(\n"))
	for k, v := range g.imports {
		out.Write([]byte(fmt.Sprintf("\t%v \"%v\"\n", k, v)))
	}
	out.Write([]byte(")\n"))
}
