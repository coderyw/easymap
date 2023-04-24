/**
 * @Author: yinwei
 * @Description:
 * @File: generator_test.go
 * @Version: 1.0.0
 * @Date: 2023/4/24 11:05
 */
package gen

import (
	"github.com/coderyw/easymap/gen/test/model"
	"os"
	"testing"
)

func Test_generator_Run(t *testing.T) {
	type fields struct {
		generator string
		pkg       string
		obj       interface{}
		outFile   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "OnlyInterface", fields: fields{
			generator: "OnlyInterface.go",
			pkg:       "github.com/coderyw/easymap/test/model",
			obj:       model.EasyMAP_exporter_OnlyInterface(nil),
			outFile:   "test/model/OnlyInterface_easymap.go",
		}, wantErr: false},
		{name: "TestStruct", fields: fields{
			generator: "TestStruct.go",
			pkg:       "github.com/coderyw/easymap/test/model",
			obj:       model.EasyMAP_exporter_TestStruct(nil),
			outFile:   "test/model/TestStruct_easymap.go",
		}, wantErr: false},
		{name: "Struct2", fields: fields{
			generator: "Struct2.go",
			pkg:       "github.com/coderyw/easymap/test/model",
			obj:       model.EasyMAP_exporter_Struct2(nil),
			outFile:   "test/model/Struct2_easymap.go",
		}, wantErr: false},
		{name: "Resp360", fields: fields{
			generator: "Resp360.go",
			pkg:       "github.com/coderyw/easymap/test/model",
			obj:       model.EasyMAP_exporter_Resp360(nil),
			outFile:   "test/model/Resp360_easymap.go",
		}, wantErr: false},
		{name: "ConfigureAliCdnDomainReq", fields: fields{
			generator: "ConfigureAliCdnDomainReq.go",
			pkg:       "github.com/coderyw/easymap/test/model",
			obj:       model.EasyMAP_exporter_ConfigureAliCdnDomainReq(nil),
			outFile:   "test/model/ConfigureAliCdnDomainReq_easymap.go",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator(tt.fields.generator)
			g.SetPkg("model", tt.fields.pkg)
			g.Add(tt.fields.obj)
			f, err := os.Create(tt.fields.outFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{
				defer f.Close()
				g.Run(f)
			}
		})
	}
}
