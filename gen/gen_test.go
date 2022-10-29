/**
 * @Author: yinwei
 * @Description:
 * @File: gen_test
 * @Version: 1.0.0
 * @Date: 2022/10/29 13:59
 */
package gen

import (
	"os"
	"testing"
	"xm-common/easymap/test"
)

func TestRun(t *testing.T) {
	g := NewGenerator("tf_easymap.go")
	g.SetPkg("gen", "gen")
	g.Add(&test.TestStruct{})
	f, err := os.Create("gen_easymap.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g.Run(f)

}
