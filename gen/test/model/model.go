package model

import (
	"github.com/coderyw/easymap/gen/test/model/model_path"
	"github.com/coderyw/easymap/gen/test/model1"
	"github.com/shopspring/decimal"
	grpc "google.golang.org/grpc"
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
type EasyMAP_exporter_InModel *InModel
type EasyMAP_exporter_ConfigureAliCdnDomainReq *ConfigureAliCdnDomainReq
type EasyMAP_exporter_OnlyInterface *OnlyInterface
type EasyMAP_exporter_ErrorStruct *ErrorStruct
type EasyMAP_exporter_StreamCtlServiceServer *StreamCtlServiceServer
type EasyMAP_exporter_caMsgServiceClient *CaMsgServiceClient
type EasyMAP_exporter_UserCache *UserCache
type EasyMAP_exporter_PayAccount *PayAccount
type EasyMAP_exporter_Connection *Connection
type EasyMAP_exporter_UserTinyInfo *UserTinyInfo

type Struct2 struct {
	DD string `json:"dd"`
}

type CaMsgServiceClient struct {
	cc *grpc.ClientConn `json:"cc,omitempty"`
}

type Resp360 struct {
	FieldInt          int     `json:"fieldInt"`
	FieldIntPtr       *int    `json:"FieldIntPtr"`
	FieldIntArr       []int   `json:"fieldIntArr"`
	FieldIntArrPtr    []*int  `json:"fieldIntArrPtr"`
	FieldIntPtrArr    *[]int  `json:"fieldIntPtrArr"`
	FieldIntPtrArrPtr *[]*int `json:"fieldIntPtrArrPtr"`

	FieldUint8          uint8     `json:"fieldUint8"`
	FieldUint8Ptr       *uint8    `json:"fieldUint8Ptr"`
	FieldUint8Arr       []uint8   `json:"fieldUint8Arr"`
	FieldUint8ArrPtr    []*uint8  `json:"fieldUint8ArrPtr"`
	FieldUint8PtrArr    *[]uint8  `json:"fieldUint8PtrArr"`
	FieldUint8PtrArrPtr *[]*uint8 `json:"fieldUint8PtrArrPtr"`

	FieldFloat32          float32     `json:"fieldFloat32"`
	FieldFloat32Ptr       *float32    `json:"fieldFloat32Ptr"`
	FieldFloat32Arr       []float32   `json:"fieldFloat32Arr"`
	FieldFloat32ArrPtr    []*float32  `json:"fieldFloat32ArrPtr"`
	FieldFloat32PtrArr    *[]float32  `json:"fieldFloat32PtrArr"`
	FieldFloat32PtrArrPtr *[]*float32 `json:"fieldFloat32PtrArrPtr"`

	FieldBool          bool     `json:"fieldBool"`
	FieldBoolPtr       *bool    `json:"fieldBoolPtr"`
	FieldBoolArr       []bool   `json:"fieldBoolArr"`
	FieldBoolArrPtr    []*bool  `json:"fieldBoolArrPtr"`
	FieldBoolPtrArr    *[]bool  `json:"fieldBoolPtrArr"`
	FieldBoolPtrArrPtr *[]*bool `json:"fieldBoolPtrArrPtr"`

	FieldString          string     `json:"fieldString"`
	FieldStringPtr       *string    `json:"fieldStringPtr"`
	FieldStringArr       []string   `json:"fieldStringArr"`
	FieldStringArrPtr    []*string  `json:"fieldStringArrPtr"`
	FieldStringPtrArr    *[]string  `json:"fieldStringPtrArr"`
	FieldStringPtrArrPtr *[]*string `json:"fieldStringPtrArrPtr"`

	Decimal          decimal.Decimal     `json:"decimal"`
	DecimalPtr       *decimal.Decimal    `json:"decimalPtr"`
	DecimalArr       []decimal.Decimal   `json:"decimalArr"`
	DecimalPtrArr    *[]decimal.Decimal  `json:"DecimalPtrArr"`
	DecimalArrPtr    []*decimal.Decimal  `json:"decimalArrPtr"`
	DecimalPtrArrPtr *[]*decimal.Decimal `json:"DecimalPtrArrPtr"`

	StructModel          InModel     `json:"structModel"`
	StructModelPtr       *InModel    `json:"structModelPtr"`
	StructModelArr       []InModel   `json:"structModelArr"`
	StructModelArrPtr    []*InModel  `json:"structModelArrPtr"`
	StructModelPtrArrPtr *[]*InModel `json:"structModelPtrArrPtr"`

	OutModel          model1.UnEasyMap     `json:"outModel"`
	OutModelPtr       *model1.UnEasyMap    `json:"outModelPtr"`
	OutModelArr       []model1.UnEasyMap   `json:"outModelArr"`
	OutModelArrPtr    []*model1.UnEasyMap  `json:"outModelArrPtr"`
	OutModelPtrArrPtr *[]*model1.UnEasyMap `json:"OutModelPtrArrPtr"`

	Inter       interface{}    `json:"inter"`
	InterArr    []interface{}  `json:"InterArr"`
	InterPtrArr *[]interface{} `json:"InterPtrArr"`

	Data []struct {
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
}

type InModel struct {
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

type UserCache struct {
	ID               int64           `json:"id,string"`
	UserID           int64           `json:"user_id,string"`
	Balance          int64           `json:"balance"`
	Nickname         string          `json:"nickname"`
	Avatar           string          `json:"avatar"`
	FbAvatar         string          `json:"fb_avatar"`
	Phone            string          `json:"phone"`
	GcashPhone       string          `json:"gcash_phone"`
	GcashCustomerID  string          `json:"gcash_customer_id"`
	WithdrawPassword string          `json:"withdraw_password"`
	LoginPassword    string          `json:"login_password"`
	PayAccountID     int64           `json:"pay_account_id,string"`
	CardBack         int             `json:"card_back,string"`
	AvatarFrame      int             `json:"avatar_frame,string"`
	AppPackageName   string          `json:"app_package_name"`
	Bind             []int           `json:"bind"`
	Type             string          `json:"type"`
	InviteUserID     string          `json:"invite_user_id"`
	InviteCode       string          `json:"invite_code"`
	Invite           *UserTinyInfo   `json:"invite"`
	IsRegister       uint8           `json:"is_register,string"`
	Enable           uint8           `json:"enable,string"`
	RechargeAmount   int64           `json:"recharge_amount"`
	WithdrawAmount   int64           `json:"withdraw_amount"`
	AvailableCoins   decimal.Decimal `json:"available_coins"`
	SPlayer          uint8           `json:"s_player,string"`
	CPlayer          uint8           `json:"c_player,string"`
	WithdrawModel    int             `json:"withdraw_model,string"`
	WithdrawControl  int             `json:"withdraw_control,string"`
	TotalRounds      int             `json:"total_rounds,string"`
	FirstRwReward    bool            `json:"first_rw_reward"`
	Level            int             `json:"level,string"`
	LevelMax         int             `json:"level_max,string"`
	Score            int             `json:"score,string"`
	CreatedAt        int             `json:"created_at,string"`
	Status           int8            `json:"status,string"`
	LastLoginChannel string          `json:"last_login_channel"`
	AppChannel       string          `json:"app_channel"`
	PayAccount       PayAccount      `json:"pay_account"`
	PayAccountArr    []PayAccount    `json:"pay_account_arr"`
	PayAccountPtrArr []*PayAccount   `json:"pay_account_ptr_arr"`
	FloatArr         []float64       `json:"float_arr"`
	StringArr        []string        `json:"string_arr"`
	IntArr           []int32         `json:"int_arr"`
	UIntArr          []uint8         `json:"uint_arr"`
	Connection       Connection      `json:"connection"`
	PlayerID         int64           `json:"player_id,string"`
	Token            string          `json:"token"`
}

type PayAccount struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
type Connection struct {
}

type UserTinyInfo struct {
	UserId     int64  `json:"user_id,string"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	FbAvatar   string `json:"fb_avatar"`
	InviteCode string `json:"invite_code"`
	Phone      string `json:"phone"`
}
