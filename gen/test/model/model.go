package model

import (
	"github.com/coderyw/easymap/gen/test/model/model_path"
	"github.com/coderyw/easymap/gen/test/model1"
	"github.com/shopspring/decimal"
	"sync"
)

type TestStruct struct {
	A         int              `json:"a"`
	B         *string          `json:"b"`
	C         float64          `json:"c"`
	D         bool             `json:"d"`
	E         uint8            `json:"e"`
	F         fs               `json:"f"`
	G         *Fs              `json:"g"`
	HH        []Fs             `json:"hh"`
	HHS       []*Fs            `json:"hhs"`
	UnEasyMap model1.UnEasyMap `json:"unEasyMap"`
}

type fs struct {
	Bs int `json:"bs"`
}
type Fs struct {
	DDD string `json:"ddd"`
}
type OutModel struct {
	UnEasyMap model1.UnEasyMap  `json:"unEasyMap"`
	PtrMap    *model1.UnEasyMap `json:"ptrMap"`
	un        *model1.UnEasyMap `json:"un"`
}

type EasyMAP_exporter_TestStruct *TestStruct
type EasyMAP_exporter_Fs *Fs
type EasyMAP_exporter_fs *fs
type EasyMAP_exporter_OutModel *OutModel
type EasyMAP_exporter_Struct2 *Struct2
type EasyMAP_exporter_Resp360 *Resp360
type EasyMAP_exporter_ConfigureAliCdnDomainReq *ConfigureAliCdnDomainReq
type EasyMAP_exporter_OnlyInterface *OnlyInterface
type EasyMAP_exporter_ErrorStruct *ErrorStruct
type EasyMAP_exporter_StreamCtlServiceServer *StreamCtlServiceServer

type Struct2 struct {
	DD string `json:"dd"`
}

type Resp360 struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	B       int64            `json:"b"`
	DD      []int            `json:"DD"`
	Dec     decimal.Decimal  `json:"dec"`
	PtrDec  *decimal.Decimal `json:"ptrDec"`
	Data    []struct {
		MatchId        string `json:"matchId"`
		MatchTimeStamp string `json:"matchTimeStamp"`
		MatchDate      string `json:"matchDate"`
		MatchTime      string `json:"matchTime"`
		CatId          string `json:"catId"`
		CatName        string `json:"catName"`
		LeagueId       string `json:"leagueId"`
		LeagueName     string `json:"leagueName"`
		HomeId         string `json:"homeId"`
		HomeName       string `json:"homeName"`
		HomeScore      string `json:"homeScore"`
		AwayId         string `json:"awayId"`
		AwayName       string `json:"awayName"`
		AwayScore      string `json:"awayScore"`
		LiveStreams    []struct {
			Id    string `json:"id"`
			Type  string `json:"type"`
			State bool   `json:"state"`
			Name  string `json:"name"`
		} `json:"liveStreams"`
		MatchState string `json:"matchState"`
	} `json:"data"`
	Else   []interface{} `json:"else"`
	AbcArr []Resp360Arr  `json:"abc_arr"`
}

type Resp360Arr struct {
	Abc string `json:"abc"`
}

type WaitGroupWrapper struct {
	sync.WaitGroup
}

type ConfigureAliCdnDomainReq struct {
	//主域名 例如：domain.com
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// 子域名 例如 : mega
	Sub string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	// 商户id
	TenantId int64 `protobuf:"varint,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 商户token
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type OnlyInterface struct {
	A interface{}
}

type Aewe int64

type ErrorStruct struct {
	ErrorCode model_path.ErrorCODE

	Str string

	ErrorCodePtr *model_path.ErrorCODE

	A int32

	Adfe Aewe
}

type StreamCtlServiceServer interface {
	CreateTencentDomainCname()
}
