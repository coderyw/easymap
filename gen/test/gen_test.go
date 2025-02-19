/**
 * @Author: yinwei
 * @Description:
 * @File: gen_test
 * @Version: 1.0.0
 * @Date: 2022/10/29 13:59
 */
package test

import (
	"fmt"
	"github.com/coderyw/easymap/gen"
	"github.com/coderyw/easymap/gen/test/model"
	"os"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	g := gen.NewGenerator("model_easymap.go")
	g.SetPkg("model", "github.com/coderyw/gen/easymap/test/model")
	g.Add(model.EasyMAP_exporter_Resp360(nil))
	g.Add(model.EasyMAP_exporter_InModel(nil))
	g.Add(model.EasyMAP_exporter_UserCache(nil))
	g.Add(model.EasyMAP_exporter_PayAccount(nil))
	g.Add(model.EasyMAP_exporter_Connection(nil))
	g.Add(model.EasyMAP_exporter_UserTinyInfo(nil))
	g.Add(model.EasyMAP_exporter_OutModel(nil))
	g.Add(model.EasyMAP_exporter_caMsgServiceClient(nil))
	//g.Add(model.EasyMAP_exporter_ConfigureAliCdnDomainReq(nil))
	//g.Add(model.EasyMAP_exporter_StreamCtlServiceServer(nil))
	f, err := os.Create("model/ErrorStruct_easymap.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g.Run(f)

}

func TestBuf(t *testing.T) {
	ttt := new(model.ErrorStruct)
	typeValue := reflect.TypeOf(ttt).Elem()
	fmt.Printf("typeValue.PkgPath = %v\n", typeValue.PkgPath())
	for i := 0; i < typeValue.NumField(); i++ {
		fmt.Println(typeValue.Field(i).Name + "-------------")
		fmt.Printf("Type.PkgPath() = %v\n", typeValue.Field(i).Type.PkgPath())
		fmt.Printf("Type.Name() = %v\n", typeValue.Field(i).Type.Name())
		fmt.Printf("Type.String() = %v\n", typeValue.Field(i).Type.String())

		fmt.Printf("Type.Kind() = %v\n", typeValue.Field(i).Type.Kind())
	}
}
