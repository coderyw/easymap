/**
 * @Author: yinwei
 * @Description:
 * @File: gen_test
 * @Version: 1.0.0
 * @Date: 2022/10/29 13:59
 */
package gen

import (
	"github.com/coderyw/easymap/gen/test/model"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	g := NewGenerator("model_easymap.go")
	g.SetPkg("model", "github.com/coderyw/easymap/test/model")
	//g.Add(model.EasyMAP_exporter_Resp360(nil))
	//g.Add(model.EasyMAP_exporter_Struct2(nil))
	//g.Add(model.EasyMAP_exporter_TestStruct(nil))
	g.Add(model.EasyMAP_exporter_ConfigureAliCdnDomainReq(nil))
	f, err := os.Create("test/model/gen_easymap.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g.Run(f)

}
