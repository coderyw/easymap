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
	"strings"
)

const pkgStrconv = "strconv"
const pkgUnsafe = "unsafe"

type generator struct {
	pkgName       string
	pkgPath       string
	fileName      string
	typesUnseen   []reflect.Type
	out           *bytes.Buffer
	imports       map[string]string
	str2BytesName string
}

func NewGenerator(filename string) *generator {
	ret := &generator{
		fileName: filename,
		imports:  map[string]string{
			//pkgStrconv: "strconv",
			//pkgUnsafe:  "unsafe",
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
	if g.pkgName == "" {
		return nil
	}
	g.writeUtil(g.out)
	var err error
	for len(g.typesUnseen) > 0 {
		t := g.typesUnseen[len(g.typesUnseen)-1]
		g.typesUnseen = g.typesUnseen[:len(g.typesUnseen)-1]
		if err = g.decode(t); err != nil {
			continue
		}
		if err = g.encode(t); err != nil {
			continue
		}
		if err = g.encodeConst(t); err != nil {
			continue
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

/*
*

	func Str2Bytes(s string) []byte {
		x := (*[2]uintptr)(unsafe.Pointer(&s))
		h := [3]uintptr{x[0], x[1], x[1]}
		return *(*[]byte)(unsafe.Pointer(&h))
	}
*/
func (g *generator) writeUtil(out io.Writer) {
	arr := strings.Split(g.fileName, ".")
	g.str2BytesName = "str2Bytes_" + strings.ReplaceAll(arr[0], "-", "_")
	fmt.Fprintln(out)
	g.imports[pkgUnsafe] = "unsafe"
	fmt.Fprintln(out, fmt.Sprintf("func %s(s string) []byte {", g.str2BytesName))
	fmt.Fprintln(out, fmt.Sprintf("\tx := (*[2]uintptr)(unsafe.Pointer(&s))"))
	fmt.Fprintln(out, fmt.Sprintf("\th := [3]uintptr{x[0], x[1], x[1]}"))
	fmt.Fprintln(out, fmt.Sprintf("\treturn *(*[]byte)(unsafe.Pointer(&h))"))
	fmt.Fprintln(out, fmt.Sprint("}"))
}
