package model

import (
	fmt "fmt"
	easymap_gen "github.com/coderyw/easymap/gen"
	easy_facade "github.com/coderyw/easymap/gen/facade"
	github_com_coderyw_easymap_gen_test_model1 "github.com/coderyw/easymap/gen/test/model1"
	decimal "github.com/shopspring/decimal"
	github_com_shopspring_decimal "github.com/shopspring/decimal"
	google_golang_org_grpc "google.golang.org/grpc"
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *CaMsgServiceClient) UnMarshalMap(m map[string]string) error {
	return nil
}
func (v *CaMsgServiceClient) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["cc"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &google_golang_org_grpc.ClientConn{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.cc = i.(*google_golang_org_grpc.ClientConn)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &google_golang_org_grpc.ClientConn{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.cc = i.(*google_golang_org_grpc.ClientConn)
				}
			}
		case google_golang_org_grpc.ClientConn:
			v.cc = &acval
		case *google_golang_org_grpc.ClientConn:
			v.cc = acval
		}

	}

	return nil
}

func (v *CaMsgServiceClient) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

func (v *CaMsgServiceClient) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	return m, nil
}

func (v *CaMsgServiceClient) CheckField(m map[string]string) error {
	if _, ok := m["cc"]; !ok {
		return fmt.Errorf("field cc is miss")
	}
	return nil
}

type CaMsgServiceClientField string

func (v CaMsgServiceClientField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	CaMsgServiceClient_cc CaMsgServiceClientField = "cc"
)

func (v *OutModel) UnMarshalMap(m map[string]string) error {
	return nil
}
func (v *OutModel) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["unEasyMap"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.UnEasyMap = *i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.UnEasyMap = *i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.UnEasyMap = acval
		case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if acval == nil {
				break
			}
			v.UnEasyMap = *acval
		}

	}
	if val, ok = m["ptrMap"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.PtrMap = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.PtrMap = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.PtrMap = &acval
		case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.PtrMap = acval
		}

	}
	if val, ok = m["un"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.un = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.un = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.un = &acval
		case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.un = acval
		}

	}

	return nil
}

func (v *OutModel) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

func (v *OutModel) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	return m, nil
}

func (v *OutModel) CheckField(m map[string]string) error {
	if _, ok := m["unEasyMap"]; !ok {
		return fmt.Errorf("field unEasyMap is miss")
	}
	if _, ok := m["ptrMap"]; !ok {
		return fmt.Errorf("field ptrMap is miss")
	}
	if _, ok := m["un"]; !ok {
		return fmt.Errorf("field un is miss")
	}
	return nil
}

type OutModelField string

func (v OutModelField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	OutModel_UnEasyMap OutModelField = "unEasyMap"
	OutModel_PtrMap    OutModelField = "ptrMap"
	OutModel_un        OutModelField = "un"
)

func (v *UserTinyInfo) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["user_id"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.UserId = int64(pv)
		}
	}
	if val, ok = m["nickname"]; ok {
		{
			pv := val
			v.Nickname = string(pv)
		}
	}
	if val, ok = m["avatar"]; ok {
		{
			pv := val
			v.Avatar = string(pv)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		{
			pv := val
			v.FbAvatar = string(pv)
		}
	}
	if val, ok = m["invite_code"]; ok {
		{
			pv := val
			v.InviteCode = string(pv)
		}
	}
	if val, ok = m["phone"]; ok {
		{
			pv := val
			v.Phone = string(pv)
		}
	}

	return nil
}
func (v *UserTinyInfo) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["user_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.UserId = int64(acval)
		case int8:
			v.UserId = int64(acval)
		case int16:
			v.UserId = int64(acval)
		case int32:
			v.UserId = int64(acval)
		case int64:
			v.UserId = int64(acval)
		case uint:
			v.UserId = int64(acval)
		case uint8:
			v.UserId = int64(acval)
		case uint16:
			v.UserId = int64(acval)
		case uint32:
			v.UserId = int64(acval)
		case uint64:
			v.UserId = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.UserId = int64(pvv)
			}
		case float32:
			v.UserId = int64(acval)
		case float64:
			v.UserId = int64(acval)
		}
	}
	if val, ok = m["nickname"]; ok {
		switch acval := val.(type) {
		case int:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Nickname = acval
		case float32:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["avatar"]; ok {
		switch acval := val.(type) {
		case int:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Avatar = acval
		case float32:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		switch acval := val.(type) {
		case int:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.FbAvatar = acval
		case float32:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite_code"]; ok {
		switch acval := val.(type) {
		case int:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.InviteCode = acval
		case float32:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case int:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Phone = acval
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}

	return nil
}

func (v *UserTinyInfo) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["user_id"] = v.UserId
	m["nickname"] = v.Nickname
	m["avatar"] = v.Avatar
	m["fb_avatar"] = v.FbAvatar
	m["invite_code"] = v.InviteCode
	m["phone"] = v.Phone
	return m, nil
}

func (v *UserTinyInfo) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["user_id"] = fmt.Sprint(v.UserId)
	m["nickname"] = fmt.Sprint(v.Nickname)
	m["avatar"] = fmt.Sprint(v.Avatar)
	m["fb_avatar"] = fmt.Sprint(v.FbAvatar)
	m["invite_code"] = fmt.Sprint(v.InviteCode)
	m["phone"] = fmt.Sprint(v.Phone)
	return m, nil
}

func (v *UserTinyInfo) CheckField(m map[string]string) error {
	if _, ok := m["user_id"]; !ok {
		return fmt.Errorf("field user_id is miss")
	}
	if _, ok := m["nickname"]; !ok {
		return fmt.Errorf("field nickname is miss")
	}
	if _, ok := m["avatar"]; !ok {
		return fmt.Errorf("field avatar is miss")
	}
	if _, ok := m["fb_avatar"]; !ok {
		return fmt.Errorf("field fb_avatar is miss")
	}
	if _, ok := m["invite_code"]; !ok {
		return fmt.Errorf("field invite_code is miss")
	}
	if _, ok := m["phone"]; !ok {
		return fmt.Errorf("field phone is miss")
	}
	return nil
}

type UserTinyInfoField string

func (v UserTinyInfoField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	UserTinyInfo_UserId     UserTinyInfoField = "user_id"
	UserTinyInfo_Nickname   UserTinyInfoField = "nickname"
	UserTinyInfo_Avatar     UserTinyInfoField = "avatar"
	UserTinyInfo_FbAvatar   UserTinyInfoField = "fb_avatar"
	UserTinyInfo_InviteCode UserTinyInfoField = "invite_code"
	UserTinyInfo_Phone      UserTinyInfoField = "phone"
)

func (v *Connection) UnMarshalMap(m map[string]string) error {
	return nil
}
func (v *Connection) UnMarshalMapInterface(m map[string]interface{}) error {
	return nil
}

func (v *Connection) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

func (v *Connection) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	return m, nil
}

func (v *Connection) CheckField(m map[string]string) error {
	return nil
}

type ConnectionField string

func (v ConnectionField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const ()

func (v *PayAccount) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["email"]; ok {
		{
			pv := val
			v.Email = string(pv)
		}
	}
	if val, ok = m["phone"]; ok {
		{
			pv := val
			v.Phone = string(pv)
		}
	}
	if val, ok = m["name"]; ok {
		{
			pv := val
			v.Name = string(pv)
		}
	}

	return nil
}
func (v *PayAccount) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["email"]; ok {
		switch acval := val.(type) {
		case int:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Email = acval
		case float32:
			v.Email = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Email = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case int:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Phone = acval
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["name"]; ok {
		switch acval := val.(type) {
		case int:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Name = acval
		case float32:
			v.Name = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Name = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}

	return nil
}

func (v *PayAccount) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["email"] = v.Email
	m["phone"] = v.Phone
	m["name"] = v.Name
	return m, nil
}

func (v *PayAccount) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["email"] = fmt.Sprint(v.Email)
	m["phone"] = fmt.Sprint(v.Phone)
	m["name"] = fmt.Sprint(v.Name)
	return m, nil
}

func (v *PayAccount) CheckField(m map[string]string) error {
	if _, ok := m["email"]; !ok {
		return fmt.Errorf("field email is miss")
	}
	if _, ok := m["phone"]; !ok {
		return fmt.Errorf("field phone is miss")
	}
	if _, ok := m["name"]; !ok {
		return fmt.Errorf("field name is miss")
	}
	return nil
}

type PayAccountField string

func (v PayAccountField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	PayAccount_Email PayAccountField = "email"
	PayAccount_Phone PayAccountField = "phone"
	PayAccount_Name  PayAccountField = "name"
)

func (v *UserCache) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["id"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.ID = int64(pv)
		}
	}
	if val, ok = m["user_id"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.UserID = int64(pv)
		}
	}
	if val, ok = m["balance"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Balance = int64(pv)
		}
	}
	if val, ok = m["nickname"]; ok {
		{
			pv := val
			v.Nickname = string(pv)
		}
	}
	if val, ok = m["avatar"]; ok {
		{
			pv := val
			v.Avatar = string(pv)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		{
			pv := val
			v.FbAvatar = string(pv)
		}
	}
	if val, ok = m["phone"]; ok {
		{
			pv := val
			v.Phone = string(pv)
		}
	}
	if val, ok = m["gcash_phone"]; ok {
		{
			pv := val
			v.GcashPhone = string(pv)
		}
	}
	if val, ok = m["gcash_customer_id"]; ok {
		{
			pv := val
			v.GcashCustomerID = string(pv)
		}
	}
	if val, ok = m["withdraw_password"]; ok {
		{
			pv := val
			v.WithdrawPassword = string(pv)
		}
	}
	if val, ok = m["login_password"]; ok {
		{
			pv := val
			v.LoginPassword = string(pv)
		}
	}
	if val, ok = m["pay_account_id"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.PayAccountID = int64(pv)
		}
	}
	if val, ok = m["card_back"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.CardBack = int(pv)
		}
	}
	if val, ok = m["avatar_frame"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.AvatarFrame = int(pv)
		}
	}
	if val, ok = m["app_package_name"]; ok {
		{
			pv := val
			v.AppPackageName = string(pv)
		}
	}
	if val, ok = m["type"]; ok {
		{
			pv := val
			v.Type = string(pv)
		}
	}
	if val, ok = m["invite_user_id"]; ok {
		{
			pv := val
			v.InviteUserID = string(pv)
		}
	}
	if val, ok = m["invite_code"]; ok {
		{
			pv := val
			v.InviteCode = string(pv)
		}
	}
	if val, ok = m["is_register"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.IsRegister = uint8(pv)
		}
	}
	if val, ok = m["enable"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.Enable = uint8(pv)
		}
	}
	if val, ok = m["recharge_amount"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.RechargeAmount = int64(pv)
		}
	}
	if val, ok = m["withdraw_amount"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.WithdrawAmount = int64(pv)
		}
	}
	if val, ok = m["available_coins"]; ok {
		var err error
		v.AvailableCoins, err = decimal.NewFromString(val)
		if err != nil {
			return err
		}
	}
	if val, ok = m["s_player"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.SPlayer = uint8(pv)
		}
	}
	if val, ok = m["c_player"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.CPlayer = uint8(pv)
		}
	}
	if val, ok = m["withdraw_model"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.WithdrawModel = int(pv)
		}
	}
	if val, ok = m["withdraw_control"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.WithdrawControl = int(pv)
		}
	}
	if val, ok = m["total_rounds"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.TotalRounds = int(pv)
		}
	}
	if val, ok = m["first_rw_reward"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			v.FirstRwReward = bool(pv)
		}
	}
	if val, ok = m["level"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Level = int(pv)
		}
	}
	if val, ok = m["level_max"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.LevelMax = int(pv)
		}
	}
	if val, ok = m["score"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Score = int(pv)
		}
	}
	if val, ok = m["created_at"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.CreatedAt = int(pv)
		}
	}
	if val, ok = m["status"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Status = int8(pv)
		}
	}
	if val, ok = m["last_login_channel"]; ok {
		{
			pv := val
			v.LastLoginChannel = string(pv)
		}
	}
	if val, ok = m["app_channel"]; ok {
		{
			pv := val
			v.AppChannel = string(pv)
		}
	}
	if val, ok = m["player_id"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.PlayerID = int64(pv)
		}
	}
	if val, ok = m["token"]; ok {
		{
			pv := val
			v.Token = string(pv)
		}
	}

	return nil
}
func (v *UserCache) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["id"]; ok {
		switch acval := val.(type) {
		case int:
			v.ID = int64(acval)
		case int8:
			v.ID = int64(acval)
		case int16:
			v.ID = int64(acval)
		case int32:
			v.ID = int64(acval)
		case int64:
			v.ID = int64(acval)
		case uint:
			v.ID = int64(acval)
		case uint8:
			v.ID = int64(acval)
		case uint16:
			v.ID = int64(acval)
		case uint32:
			v.ID = int64(acval)
		case uint64:
			v.ID = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.ID = int64(pvv)
			}
		case float32:
			v.ID = int64(acval)
		case float64:
			v.ID = int64(acval)
		}
	}
	if val, ok = m["user_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.UserID = int64(acval)
		case int8:
			v.UserID = int64(acval)
		case int16:
			v.UserID = int64(acval)
		case int32:
			v.UserID = int64(acval)
		case int64:
			v.UserID = int64(acval)
		case uint:
			v.UserID = int64(acval)
		case uint8:
			v.UserID = int64(acval)
		case uint16:
			v.UserID = int64(acval)
		case uint32:
			v.UserID = int64(acval)
		case uint64:
			v.UserID = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.UserID = int64(pvv)
			}
		case float32:
			v.UserID = int64(acval)
		case float64:
			v.UserID = int64(acval)
		}
	}
	if val, ok = m["balance"]; ok {
		switch acval := val.(type) {
		case int:
			v.Balance = int64(acval)
		case int8:
			v.Balance = int64(acval)
		case int16:
			v.Balance = int64(acval)
		case int32:
			v.Balance = int64(acval)
		case int64:
			v.Balance = int64(acval)
		case uint:
			v.Balance = int64(acval)
		case uint8:
			v.Balance = int64(acval)
		case uint16:
			v.Balance = int64(acval)
		case uint32:
			v.Balance = int64(acval)
		case uint64:
			v.Balance = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Balance = int64(pvv)
			}
		case float32:
			v.Balance = int64(acval)
		case float64:
			v.Balance = int64(acval)
		}
	}
	if val, ok = m["nickname"]; ok {
		switch acval := val.(type) {
		case int:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Nickname = acval
		case float32:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["avatar"]; ok {
		switch acval := val.(type) {
		case int:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Avatar = acval
		case float32:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		switch acval := val.(type) {
		case int:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.FbAvatar = acval
		case float32:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case int:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Phone = acval
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["gcash_phone"]; ok {
		switch acval := val.(type) {
		case int:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.GcashPhone = acval
		case float32:
			v.GcashPhone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.GcashPhone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["gcash_customer_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.GcashCustomerID = acval
		case float32:
			v.GcashCustomerID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.GcashCustomerID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["withdraw_password"]; ok {
		switch acval := val.(type) {
		case int:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.WithdrawPassword = acval
		case float32:
			v.WithdrawPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.WithdrawPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["login_password"]; ok {
		switch acval := val.(type) {
		case int:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.LoginPassword = acval
		case float32:
			v.LoginPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.LoginPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["pay_account_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.PayAccountID = int64(acval)
		case int8:
			v.PayAccountID = int64(acval)
		case int16:
			v.PayAccountID = int64(acval)
		case int32:
			v.PayAccountID = int64(acval)
		case int64:
			v.PayAccountID = int64(acval)
		case uint:
			v.PayAccountID = int64(acval)
		case uint8:
			v.PayAccountID = int64(acval)
		case uint16:
			v.PayAccountID = int64(acval)
		case uint32:
			v.PayAccountID = int64(acval)
		case uint64:
			v.PayAccountID = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.PayAccountID = int64(pvv)
			}
		case float32:
			v.PayAccountID = int64(acval)
		case float64:
			v.PayAccountID = int64(acval)
		}
	}
	if val, ok = m["card_back"]; ok {
		switch acval := val.(type) {
		case int:
			v.CardBack = int(acval)
		case int8:
			v.CardBack = int(acval)
		case int16:
			v.CardBack = int(acval)
		case int32:
			v.CardBack = int(acval)
		case int64:
			v.CardBack = int(acval)
		case uint:
			v.CardBack = int(acval)
		case uint8:
			v.CardBack = int(acval)
		case uint16:
			v.CardBack = int(acval)
		case uint32:
			v.CardBack = int(acval)
		case uint64:
			v.CardBack = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CardBack = int(pvv)
			}
		case float32:
			v.CardBack = int(acval)
		case float64:
			v.CardBack = int(acval)
		}
	}
	if val, ok = m["avatar_frame"]; ok {
		switch acval := val.(type) {
		case int:
			v.AvatarFrame = int(acval)
		case int8:
			v.AvatarFrame = int(acval)
		case int16:
			v.AvatarFrame = int(acval)
		case int32:
			v.AvatarFrame = int(acval)
		case int64:
			v.AvatarFrame = int(acval)
		case uint:
			v.AvatarFrame = int(acval)
		case uint8:
			v.AvatarFrame = int(acval)
		case uint16:
			v.AvatarFrame = int(acval)
		case uint32:
			v.AvatarFrame = int(acval)
		case uint64:
			v.AvatarFrame = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.AvatarFrame = int(pvv)
			}
		case float32:
			v.AvatarFrame = int(acval)
		case float64:
			v.AvatarFrame = int(acval)
		}
	}
	if val, ok = m["app_package_name"]; ok {
		switch acval := val.(type) {
		case int:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.AppPackageName = acval
		case float32:
			v.AppPackageName = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.AppPackageName = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["bind"]; ok {
		switch acval := val.(type) {
		case []int:
			v.Bind = acval
		case []int8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []float64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []int16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []int32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []int64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []uint:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []uint8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []uint32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []uint64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []interface{}:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = int(trueK)
				case int8:
					tmpArr[i] = int(trueK)
				case int16:
					tmpArr[i] = int(trueK)
				case int32:
					tmpArr[i] = int(trueK)
				case int64:
					tmpArr[i] = int(trueK)
				case uint:
					tmpArr[i] = int(trueK)
				case uint8:
					tmpArr[i] = int(trueK)
				case uint16:
					tmpArr[i] = int(trueK)
				case uint32:
					tmpArr[i] = int(trueK)
				case uint64:
					tmpArr[i] = int(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = int(tmp)
				case float32:
					tmpArr[i] = int(trueK)
				case float64:
					tmpArr[i] = int(trueK)
				default:
					break
				}
			}
			v.Bind = tmpArr
		case []uint16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []string:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = int(tmp)
			}
			v.Bind = tmpArr
		case []float32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		}
	}
	if val, ok = m["type"]; ok {
		switch acval := val.(type) {
		case int:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Type = acval
		case float32:
			v.Type = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Type = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite_user_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.InviteUserID = acval
		case float32:
			v.InviteUserID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.InviteUserID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite_code"]; ok {
		switch acval := val.(type) {
		case int:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.InviteCode = acval
		case float32:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &UserTinyInfo{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.Invite = i.(*UserTinyInfo)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &UserTinyInfo{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.Invite = i.(*UserTinyInfo)
				}
			}
		case UserTinyInfo:
			v.Invite = &acval
		case *UserTinyInfo:
			v.Invite = acval
		}

	}
	if val, ok = m["is_register"]; ok {
		switch acval := val.(type) {
		case int:
			v.IsRegister = uint8(acval)
		case int8:
			v.IsRegister = uint8(acval)
		case int16:
			v.IsRegister = uint8(acval)
		case int32:
			v.IsRegister = uint8(acval)
		case int64:
			v.IsRegister = uint8(acval)
		case uint:
			v.IsRegister = uint8(acval)
		case uint8:
			v.IsRegister = uint8(acval)
		case uint16:
			v.IsRegister = uint8(acval)
		case uint32:
			v.IsRegister = uint8(acval)
		case uint64:
			v.IsRegister = uint8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.IsRegister = uint8(pvv)
			}
		case float32:
			v.IsRegister = uint8(acval)
		case float64:
			v.IsRegister = uint8(acval)
		}
	}
	if val, ok = m["enable"]; ok {
		switch acval := val.(type) {
		case int:
			v.Enable = uint8(acval)
		case int8:
			v.Enable = uint8(acval)
		case int16:
			v.Enable = uint8(acval)
		case int32:
			v.Enable = uint8(acval)
		case int64:
			v.Enable = uint8(acval)
		case uint:
			v.Enable = uint8(acval)
		case uint8:
			v.Enable = uint8(acval)
		case uint16:
			v.Enable = uint8(acval)
		case uint32:
			v.Enable = uint8(acval)
		case uint64:
			v.Enable = uint8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Enable = uint8(pvv)
			}
		case float32:
			v.Enable = uint8(acval)
		case float64:
			v.Enable = uint8(acval)
		}
	}
	if val, ok = m["recharge_amount"]; ok {
		switch acval := val.(type) {
		case int:
			v.RechargeAmount = int64(acval)
		case int8:
			v.RechargeAmount = int64(acval)
		case int16:
			v.RechargeAmount = int64(acval)
		case int32:
			v.RechargeAmount = int64(acval)
		case int64:
			v.RechargeAmount = int64(acval)
		case uint:
			v.RechargeAmount = int64(acval)
		case uint8:
			v.RechargeAmount = int64(acval)
		case uint16:
			v.RechargeAmount = int64(acval)
		case uint32:
			v.RechargeAmount = int64(acval)
		case uint64:
			v.RechargeAmount = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.RechargeAmount = int64(pvv)
			}
		case float32:
			v.RechargeAmount = int64(acval)
		case float64:
			v.RechargeAmount = int64(acval)
		}
	}
	if val, ok = m["withdraw_amount"]; ok {
		switch acval := val.(type) {
		case int:
			v.WithdrawAmount = int64(acval)
		case int8:
			v.WithdrawAmount = int64(acval)
		case int16:
			v.WithdrawAmount = int64(acval)
		case int32:
			v.WithdrawAmount = int64(acval)
		case int64:
			v.WithdrawAmount = int64(acval)
		case uint:
			v.WithdrawAmount = int64(acval)
		case uint8:
			v.WithdrawAmount = int64(acval)
		case uint16:
			v.WithdrawAmount = int64(acval)
		case uint32:
			v.WithdrawAmount = int64(acval)
		case uint64:
			v.WithdrawAmount = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.WithdrawAmount = int64(pvv)
			}
		case float32:
			v.WithdrawAmount = int64(acval)
		case float64:
			v.WithdrawAmount = int64(acval)
		}
	}
	if val, ok = m["available_coins"]; ok {
		switch acval := val.(type) {
		case int:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int8:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int16:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int32:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int64:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint:
			v.AvailableCoins = decimal.NewFromUint64(uint64(acval))
		case uint8:
			v.AvailableCoins = decimal.NewFromUint64(uint64(acval))
		case uint16:
			v.AvailableCoins = decimal.NewFromUint64(uint64(acval))
		case uint32:
			v.AvailableCoins = decimal.NewFromUint64(uint64(acval))
		case uint64:
			v.AvailableCoins = decimal.NewFromUint64(uint64(acval))
		case string:
			var err error
			v.AvailableCoins, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
		case float32:
			v.AvailableCoins = decimal.NewFromFloat(float64(acval))
		case float64:
			v.AvailableCoins = decimal.NewFromFloat(float64(acval))
		}
	}
	if val, ok = m["s_player"]; ok {
		switch acval := val.(type) {
		case int:
			v.SPlayer = uint8(acval)
		case int8:
			v.SPlayer = uint8(acval)
		case int16:
			v.SPlayer = uint8(acval)
		case int32:
			v.SPlayer = uint8(acval)
		case int64:
			v.SPlayer = uint8(acval)
		case uint:
			v.SPlayer = uint8(acval)
		case uint8:
			v.SPlayer = uint8(acval)
		case uint16:
			v.SPlayer = uint8(acval)
		case uint32:
			v.SPlayer = uint8(acval)
		case uint64:
			v.SPlayer = uint8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.SPlayer = uint8(pvv)
			}
		case float32:
			v.SPlayer = uint8(acval)
		case float64:
			v.SPlayer = uint8(acval)
		}
	}
	if val, ok = m["c_player"]; ok {
		switch acval := val.(type) {
		case int:
			v.CPlayer = uint8(acval)
		case int8:
			v.CPlayer = uint8(acval)
		case int16:
			v.CPlayer = uint8(acval)
		case int32:
			v.CPlayer = uint8(acval)
		case int64:
			v.CPlayer = uint8(acval)
		case uint:
			v.CPlayer = uint8(acval)
		case uint8:
			v.CPlayer = uint8(acval)
		case uint16:
			v.CPlayer = uint8(acval)
		case uint32:
			v.CPlayer = uint8(acval)
		case uint64:
			v.CPlayer = uint8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CPlayer = uint8(pvv)
			}
		case float32:
			v.CPlayer = uint8(acval)
		case float64:
			v.CPlayer = uint8(acval)
		}
	}
	if val, ok = m["withdraw_model"]; ok {
		switch acval := val.(type) {
		case int:
			v.WithdrawModel = int(acval)
		case int8:
			v.WithdrawModel = int(acval)
		case int16:
			v.WithdrawModel = int(acval)
		case int32:
			v.WithdrawModel = int(acval)
		case int64:
			v.WithdrawModel = int(acval)
		case uint:
			v.WithdrawModel = int(acval)
		case uint8:
			v.WithdrawModel = int(acval)
		case uint16:
			v.WithdrawModel = int(acval)
		case uint32:
			v.WithdrawModel = int(acval)
		case uint64:
			v.WithdrawModel = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.WithdrawModel = int(pvv)
			}
		case float32:
			v.WithdrawModel = int(acval)
		case float64:
			v.WithdrawModel = int(acval)
		}
	}
	if val, ok = m["withdraw_control"]; ok {
		switch acval := val.(type) {
		case int:
			v.WithdrawControl = int(acval)
		case int8:
			v.WithdrawControl = int(acval)
		case int16:
			v.WithdrawControl = int(acval)
		case int32:
			v.WithdrawControl = int(acval)
		case int64:
			v.WithdrawControl = int(acval)
		case uint:
			v.WithdrawControl = int(acval)
		case uint8:
			v.WithdrawControl = int(acval)
		case uint16:
			v.WithdrawControl = int(acval)
		case uint32:
			v.WithdrawControl = int(acval)
		case uint64:
			v.WithdrawControl = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.WithdrawControl = int(pvv)
			}
		case float32:
			v.WithdrawControl = int(acval)
		case float64:
			v.WithdrawControl = int(acval)
		}
	}
	if val, ok = m["total_rounds"]; ok {
		switch acval := val.(type) {
		case int:
			v.TotalRounds = int(acval)
		case int8:
			v.TotalRounds = int(acval)
		case int16:
			v.TotalRounds = int(acval)
		case int32:
			v.TotalRounds = int(acval)
		case int64:
			v.TotalRounds = int(acval)
		case uint:
			v.TotalRounds = int(acval)
		case uint8:
			v.TotalRounds = int(acval)
		case uint16:
			v.TotalRounds = int(acval)
		case uint32:
			v.TotalRounds = int(acval)
		case uint64:
			v.TotalRounds = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.TotalRounds = int(pvv)
			}
		case float32:
			v.TotalRounds = int(acval)
		case float64:
			v.TotalRounds = int(acval)
		}
	}
	if val, ok = m["first_rw_reward"]; ok {
		switch acval := val.(type) {
		case bool:
			v.FirstRwReward = bool(acval)
		}
	}
	if val, ok = m["level"]; ok {
		switch acval := val.(type) {
		case int:
			v.Level = int(acval)
		case int8:
			v.Level = int(acval)
		case int16:
			v.Level = int(acval)
		case int32:
			v.Level = int(acval)
		case int64:
			v.Level = int(acval)
		case uint:
			v.Level = int(acval)
		case uint8:
			v.Level = int(acval)
		case uint16:
			v.Level = int(acval)
		case uint32:
			v.Level = int(acval)
		case uint64:
			v.Level = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Level = int(pvv)
			}
		case float32:
			v.Level = int(acval)
		case float64:
			v.Level = int(acval)
		}
	}
	if val, ok = m["level_max"]; ok {
		switch acval := val.(type) {
		case int:
			v.LevelMax = int(acval)
		case int8:
			v.LevelMax = int(acval)
		case int16:
			v.LevelMax = int(acval)
		case int32:
			v.LevelMax = int(acval)
		case int64:
			v.LevelMax = int(acval)
		case uint:
			v.LevelMax = int(acval)
		case uint8:
			v.LevelMax = int(acval)
		case uint16:
			v.LevelMax = int(acval)
		case uint32:
			v.LevelMax = int(acval)
		case uint64:
			v.LevelMax = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.LevelMax = int(pvv)
			}
		case float32:
			v.LevelMax = int(acval)
		case float64:
			v.LevelMax = int(acval)
		}
	}
	if val, ok = m["score"]; ok {
		switch acval := val.(type) {
		case int:
			v.Score = int(acval)
		case int8:
			v.Score = int(acval)
		case int16:
			v.Score = int(acval)
		case int32:
			v.Score = int(acval)
		case int64:
			v.Score = int(acval)
		case uint:
			v.Score = int(acval)
		case uint8:
			v.Score = int(acval)
		case uint16:
			v.Score = int(acval)
		case uint32:
			v.Score = int(acval)
		case uint64:
			v.Score = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Score = int(pvv)
			}
		case float32:
			v.Score = int(acval)
		case float64:
			v.Score = int(acval)
		}
	}
	if val, ok = m["created_at"]; ok {
		switch acval := val.(type) {
		case int:
			v.CreatedAt = int(acval)
		case int8:
			v.CreatedAt = int(acval)
		case int16:
			v.CreatedAt = int(acval)
		case int32:
			v.CreatedAt = int(acval)
		case int64:
			v.CreatedAt = int(acval)
		case uint:
			v.CreatedAt = int(acval)
		case uint8:
			v.CreatedAt = int(acval)
		case uint16:
			v.CreatedAt = int(acval)
		case uint32:
			v.CreatedAt = int(acval)
		case uint64:
			v.CreatedAt = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CreatedAt = int(pvv)
			}
		case float32:
			v.CreatedAt = int(acval)
		case float64:
			v.CreatedAt = int(acval)
		}
	}
	if val, ok = m["status"]; ok {
		switch acval := val.(type) {
		case int:
			v.Status = int8(acval)
		case int8:
			v.Status = int8(acval)
		case int16:
			v.Status = int8(acval)
		case int32:
			v.Status = int8(acval)
		case int64:
			v.Status = int8(acval)
		case uint:
			v.Status = int8(acval)
		case uint8:
			v.Status = int8(acval)
		case uint16:
			v.Status = int8(acval)
		case uint32:
			v.Status = int8(acval)
		case uint64:
			v.Status = int8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Status = int8(pvv)
			}
		case float32:
			v.Status = int8(acval)
		case float64:
			v.Status = int8(acval)
		}
	}
	if val, ok = m["last_login_channel"]; ok {
		switch acval := val.(type) {
		case int:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.LastLoginChannel = acval
		case float32:
			v.LastLoginChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.LastLoginChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["app_channel"]; ok {
		switch acval := val.(type) {
		case int:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.AppChannel = acval
		case float32:
			v.AppChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.AppChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["pay_account"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &PayAccount{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.PayAccount = *i.(*PayAccount)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &PayAccount{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.PayAccount = *i.(*PayAccount)
				}
			}
		case PayAccount:
			v.PayAccount = acval
		case *PayAccount:
			if acval == nil {
				break
			}
			v.PayAccount = *acval
		}

	}
	if val, ok = m["pay_account_arr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]PayAccount, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &PayAccount{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*PayAccount)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &PayAccount{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*PayAccount)
					}
				case PayAccount:
					tmpArr[i] = trueK
				case *PayAccount:
					tmpArr[i] = *trueK
				}
			}
			v.PayAccountArr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]PayAccount, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &PayAccount{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*PayAccount)
			}
			v.PayAccountArr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]PayAccount, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &PayAccount{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*PayAccount)
			}
			v.PayAccountArr = tmpArr

		case []PayAccount:
			if len(acval) > 0 {
				tmpArr := make([]PayAccount, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.PayAccountArr = tmpArr
			}

		case []*PayAccount:
			if len(acval) > 0 {
				tmpArr := make([]PayAccount, len(acval))
				for i, k := range acval {
					tmpArr[i] = *k
				}
				v.PayAccountArr = tmpArr
			}
		}
	}
	if val, ok = m["pay_account_ptr_arr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*PayAccount, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &PayAccount{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*PayAccount)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &PayAccount{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*PayAccount)
					}
				case PayAccount:
					tmpArr[i] = &trueK
				case *PayAccount:
					tmpArr[i] = trueK
				}
			}
			v.PayAccountPtrArr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*PayAccount, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &PayAccount{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*PayAccount)
			}
			v.PayAccountPtrArr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*PayAccount, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &PayAccount{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*PayAccount)
			}
			v.PayAccountPtrArr = tmpArr

		case []PayAccount:
			if len(acval) > 0 {
				tmpArr := make([]*PayAccount, len(acval))
				for i, k := range acval {
					tmpArr[i] = &k
				}
				v.PayAccountPtrArr = tmpArr
			}

		case []*PayAccount:
			if len(acval) > 0 {
				tmpArr := make([]*PayAccount, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.PayAccountPtrArr = tmpArr
			}
		}
	}
	if val, ok = m["float_arr"]; ok {
		switch acval := val.(type) {
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []float32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			v.FloatArr = acval
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = float64(trueK)
				case float64:
					tmpArr[i] = float64(trueK)
				case int:
					tmpArr[i] = float64(trueK)
				case int8:
					tmpArr[i] = float64(trueK)
				case int16:
					tmpArr[i] = float64(trueK)
				case int32:
					tmpArr[i] = float64(trueK)
				case int64:
					tmpArr[i] = float64(trueK)
				case uint:
					tmpArr[i] = float64(trueK)
				case uint8:
					tmpArr[i] = float64(trueK)
				case uint16:
					tmpArr[i] = float64(trueK)
				case uint32:
					tmpArr[i] = float64(trueK)
				case uint64:
					tmpArr[i] = float64(trueK)
				case string:
					tmp, err := strconv.ParseFloat(trueK, 10)
					if err != nil {
						return err
					}
					tmpArr[i] = float64(tmp)
				default:
					break
				}
			}
			v.FloatArr = tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpArr[i] = float64(tmp)
			}
			v.FloatArr = tmpArr
		}
	}
	if val, ok = m["string_arr"]; ok {
		switch acval := val.(type) {
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case float64:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case int:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int8:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int16:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int32:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int64:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case uint:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint8:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint16:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint32:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint64:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case string:
					tmpArr[i] = trueK
				default:
					break
				}
			}
			v.StringArr = tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatFloat(float64(k), 'f', -1, 64)
			}
			v.StringArr = tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			v.StringArr = acval
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		}
	}
	if val, ok = m["int_arr"]; ok {
		switch acval := val.(type) {
		case []int:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []int8:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []float32:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []int32:
			v.IntArr = acval
		case []uint16:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []string:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = int32(tmp)
			}
			v.IntArr = tmpArr
		case []uint:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []uint64:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []interface{}:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = int32(trueK)
				case int8:
					tmpArr[i] = int32(trueK)
				case int16:
					tmpArr[i] = int32(trueK)
				case int32:
					tmpArr[i] = int32(trueK)
				case int64:
					tmpArr[i] = int32(trueK)
				case uint:
					tmpArr[i] = int32(trueK)
				case uint8:
					tmpArr[i] = int32(trueK)
				case uint16:
					tmpArr[i] = int32(trueK)
				case uint32:
					tmpArr[i] = int32(trueK)
				case uint64:
					tmpArr[i] = int32(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = int32(tmp)
				case float32:
					tmpArr[i] = int32(trueK)
				case float64:
					tmpArr[i] = int32(trueK)
				default:
					break
				}
			}
			v.IntArr = tmpArr
		case []int16:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []int64:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []uint8:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []uint32:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []float64:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		}
	}
	if val, ok = m["uint_arr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = uint8(trueK)
				case int8:
					tmpArr[i] = uint8(trueK)
				case int16:
					tmpArr[i] = uint8(trueK)
				case int32:
					tmpArr[i] = uint8(trueK)
				case int64:
					tmpArr[i] = uint8(trueK)
				case uint:
					tmpArr[i] = uint8(trueK)
				case uint8:
					tmpArr[i] = uint8(trueK)
				case uint16:
					tmpArr[i] = uint8(trueK)
				case uint32:
					tmpArr[i] = uint8(trueK)
				case uint64:
					tmpArr[i] = uint8(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = uint8(tmp)
				case float32:
					tmpArr[i] = uint8(trueK)
				case float64:
					tmpArr[i] = uint8(trueK)
				default:
					break
				}
			}
			v.UIntArr = tmpArr
		case []float32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []int8:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []uint8:
			v.UIntArr = acval
		case []string:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseUint(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = uint8(tmp)
			}
			v.UIntArr = tmpArr
		case []float64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []int16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []int64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []uint64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []int:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []uint:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []uint16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []int32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		case []uint32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.UIntArr = tmpArr
		}
	}
	if val, ok = m["connection"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &Connection{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.Connection = *i.(*Connection)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &Connection{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.Connection = *i.(*Connection)
				}
			}
		case Connection:
			v.Connection = acval
		case *Connection:
			if acval == nil {
				break
			}
			v.Connection = *acval
		}

	}
	if val, ok = m["player_id"]; ok {
		switch acval := val.(type) {
		case int:
			v.PlayerID = int64(acval)
		case int8:
			v.PlayerID = int64(acval)
		case int16:
			v.PlayerID = int64(acval)
		case int32:
			v.PlayerID = int64(acval)
		case int64:
			v.PlayerID = int64(acval)
		case uint:
			v.PlayerID = int64(acval)
		case uint8:
			v.PlayerID = int64(acval)
		case uint16:
			v.PlayerID = int64(acval)
		case uint32:
			v.PlayerID = int64(acval)
		case uint64:
			v.PlayerID = int64(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.PlayerID = int64(pvv)
			}
		case float32:
			v.PlayerID = int64(acval)
		case float64:
			v.PlayerID = int64(acval)
		}
	}
	if val, ok = m["token"]; ok {
		switch acval := val.(type) {
		case int:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Token = acval
		case float32:
			v.Token = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Token = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}

	return nil
}

func (v *UserCache) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["user_id"] = v.UserID
	m["balance"] = v.Balance
	m["nickname"] = v.Nickname
	m["avatar"] = v.Avatar
	m["fb_avatar"] = v.FbAvatar
	m["phone"] = v.Phone
	m["gcash_phone"] = v.GcashPhone
	m["gcash_customer_id"] = v.GcashCustomerID
	m["withdraw_password"] = v.WithdrawPassword
	m["login_password"] = v.LoginPassword
	m["pay_account_id"] = v.PayAccountID
	m["card_back"] = v.CardBack
	m["avatar_frame"] = v.AvatarFrame
	m["app_package_name"] = v.AppPackageName
	m["type"] = v.Type
	m["invite_user_id"] = v.InviteUserID
	m["invite_code"] = v.InviteCode
	m["is_register"] = v.IsRegister
	m["enable"] = v.Enable
	m["recharge_amount"] = v.RechargeAmount
	m["withdraw_amount"] = v.WithdrawAmount
	m["available_coins"] = v.AvailableCoins.String()
	m["s_player"] = v.SPlayer
	m["c_player"] = v.CPlayer
	m["withdraw_model"] = v.WithdrawModel
	m["withdraw_control"] = v.WithdrawControl
	m["total_rounds"] = v.TotalRounds
	m["first_rw_reward"] = v.FirstRwReward
	m["level"] = v.Level
	m["level_max"] = v.LevelMax
	m["score"] = v.Score
	m["created_at"] = v.CreatedAt
	m["status"] = v.Status
	m["last_login_channel"] = v.LastLoginChannel
	m["app_channel"] = v.AppChannel
	m["player_id"] = v.PlayerID
	m["token"] = v.Token
	return m, nil
}

func (v *UserCache) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["id"] = fmt.Sprint(v.ID)
	m["user_id"] = fmt.Sprint(v.UserID)
	m["balance"] = fmt.Sprint(v.Balance)
	m["nickname"] = fmt.Sprint(v.Nickname)
	m["avatar"] = fmt.Sprint(v.Avatar)
	m["fb_avatar"] = fmt.Sprint(v.FbAvatar)
	m["phone"] = fmt.Sprint(v.Phone)
	m["gcash_phone"] = fmt.Sprint(v.GcashPhone)
	m["gcash_customer_id"] = fmt.Sprint(v.GcashCustomerID)
	m["withdraw_password"] = fmt.Sprint(v.WithdrawPassword)
	m["login_password"] = fmt.Sprint(v.LoginPassword)
	m["pay_account_id"] = fmt.Sprint(v.PayAccountID)
	m["card_back"] = fmt.Sprint(v.CardBack)
	m["avatar_frame"] = fmt.Sprint(v.AvatarFrame)
	m["app_package_name"] = fmt.Sprint(v.AppPackageName)
	m["type"] = fmt.Sprint(v.Type)
	m["invite_user_id"] = fmt.Sprint(v.InviteUserID)
	m["invite_code"] = fmt.Sprint(v.InviteCode)
	m["is_register"] = fmt.Sprint(v.IsRegister)
	m["enable"] = fmt.Sprint(v.Enable)
	m["recharge_amount"] = fmt.Sprint(v.RechargeAmount)
	m["withdraw_amount"] = fmt.Sprint(v.WithdrawAmount)
	m["available_coins"] = v.AvailableCoins.String()
	m["s_player"] = fmt.Sprint(v.SPlayer)
	m["c_player"] = fmt.Sprint(v.CPlayer)
	m["withdraw_model"] = fmt.Sprint(v.WithdrawModel)
	m["withdraw_control"] = fmt.Sprint(v.WithdrawControl)
	m["total_rounds"] = fmt.Sprint(v.TotalRounds)
	m["first_rw_reward"] = fmt.Sprint(v.FirstRwReward)
	m["level"] = fmt.Sprint(v.Level)
	m["level_max"] = fmt.Sprint(v.LevelMax)
	m["score"] = fmt.Sprint(v.Score)
	m["created_at"] = fmt.Sprint(v.CreatedAt)
	m["status"] = fmt.Sprint(v.Status)
	m["last_login_channel"] = fmt.Sprint(v.LastLoginChannel)
	m["app_channel"] = fmt.Sprint(v.AppChannel)
	m["player_id"] = fmt.Sprint(v.PlayerID)
	m["token"] = fmt.Sprint(v.Token)
	return m, nil
}

func (v *UserCache) CheckField(m map[string]string) error {
	if _, ok := m["id"]; !ok {
		return fmt.Errorf("field id is miss")
	}
	if _, ok := m["user_id"]; !ok {
		return fmt.Errorf("field user_id is miss")
	}
	if _, ok := m["balance"]; !ok {
		return fmt.Errorf("field balance is miss")
	}
	if _, ok := m["nickname"]; !ok {
		return fmt.Errorf("field nickname is miss")
	}
	if _, ok := m["avatar"]; !ok {
		return fmt.Errorf("field avatar is miss")
	}
	if _, ok := m["fb_avatar"]; !ok {
		return fmt.Errorf("field fb_avatar is miss")
	}
	if _, ok := m["phone"]; !ok {
		return fmt.Errorf("field phone is miss")
	}
	if _, ok := m["gcash_phone"]; !ok {
		return fmt.Errorf("field gcash_phone is miss")
	}
	if _, ok := m["gcash_customer_id"]; !ok {
		return fmt.Errorf("field gcash_customer_id is miss")
	}
	if _, ok := m["withdraw_password"]; !ok {
		return fmt.Errorf("field withdraw_password is miss")
	}
	if _, ok := m["login_password"]; !ok {
		return fmt.Errorf("field login_password is miss")
	}
	if _, ok := m["pay_account_id"]; !ok {
		return fmt.Errorf("field pay_account_id is miss")
	}
	if _, ok := m["card_back"]; !ok {
		return fmt.Errorf("field card_back is miss")
	}
	if _, ok := m["avatar_frame"]; !ok {
		return fmt.Errorf("field avatar_frame is miss")
	}
	if _, ok := m["app_package_name"]; !ok {
		return fmt.Errorf("field app_package_name is miss")
	}
	if _, ok := m["bind"]; !ok {
		return fmt.Errorf("field bind is miss")
	}
	if _, ok := m["type"]; !ok {
		return fmt.Errorf("field type is miss")
	}
	if _, ok := m["invite_user_id"]; !ok {
		return fmt.Errorf("field invite_user_id is miss")
	}
	if _, ok := m["invite_code"]; !ok {
		return fmt.Errorf("field invite_code is miss")
	}
	if _, ok := m["invite"]; !ok {
		return fmt.Errorf("field invite is miss")
	}
	if _, ok := m["is_register"]; !ok {
		return fmt.Errorf("field is_register is miss")
	}
	if _, ok := m["enable"]; !ok {
		return fmt.Errorf("field enable is miss")
	}
	if _, ok := m["recharge_amount"]; !ok {
		return fmt.Errorf("field recharge_amount is miss")
	}
	if _, ok := m["withdraw_amount"]; !ok {
		return fmt.Errorf("field withdraw_amount is miss")
	}
	if _, ok := m["available_coins"]; !ok {
		return fmt.Errorf("field available_coins is miss")
	}
	if _, ok := m["s_player"]; !ok {
		return fmt.Errorf("field s_player is miss")
	}
	if _, ok := m["c_player"]; !ok {
		return fmt.Errorf("field c_player is miss")
	}
	if _, ok := m["withdraw_model"]; !ok {
		return fmt.Errorf("field withdraw_model is miss")
	}
	if _, ok := m["withdraw_control"]; !ok {
		return fmt.Errorf("field withdraw_control is miss")
	}
	if _, ok := m["total_rounds"]; !ok {
		return fmt.Errorf("field total_rounds is miss")
	}
	if _, ok := m["first_rw_reward"]; !ok {
		return fmt.Errorf("field first_rw_reward is miss")
	}
	if _, ok := m["level"]; !ok {
		return fmt.Errorf("field level is miss")
	}
	if _, ok := m["level_max"]; !ok {
		return fmt.Errorf("field level_max is miss")
	}
	if _, ok := m["score"]; !ok {
		return fmt.Errorf("field score is miss")
	}
	if _, ok := m["created_at"]; !ok {
		return fmt.Errorf("field created_at is miss")
	}
	if _, ok := m["status"]; !ok {
		return fmt.Errorf("field status is miss")
	}
	if _, ok := m["last_login_channel"]; !ok {
		return fmt.Errorf("field last_login_channel is miss")
	}
	if _, ok := m["app_channel"]; !ok {
		return fmt.Errorf("field app_channel is miss")
	}
	if _, ok := m["pay_account"]; !ok {
		return fmt.Errorf("field pay_account is miss")
	}
	if _, ok := m["pay_account_arr"]; !ok {
		return fmt.Errorf("field pay_account_arr is miss")
	}
	if _, ok := m["pay_account_ptr_arr"]; !ok {
		return fmt.Errorf("field pay_account_ptr_arr is miss")
	}
	if _, ok := m["float_arr"]; !ok {
		return fmt.Errorf("field float_arr is miss")
	}
	if _, ok := m["string_arr"]; !ok {
		return fmt.Errorf("field string_arr is miss")
	}
	if _, ok := m["int_arr"]; !ok {
		return fmt.Errorf("field int_arr is miss")
	}
	if _, ok := m["uint_arr"]; !ok {
		return fmt.Errorf("field uint_arr is miss")
	}
	if _, ok := m["connection"]; !ok {
		return fmt.Errorf("field connection is miss")
	}
	if _, ok := m["player_id"]; !ok {
		return fmt.Errorf("field player_id is miss")
	}
	if _, ok := m["token"]; !ok {
		return fmt.Errorf("field token is miss")
	}
	return nil
}

type UserCacheField string

func (v UserCacheField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	UserCache_ID               UserCacheField = "id"
	UserCache_UserID           UserCacheField = "user_id"
	UserCache_Balance          UserCacheField = "balance"
	UserCache_Nickname         UserCacheField = "nickname"
	UserCache_Avatar           UserCacheField = "avatar"
	UserCache_FbAvatar         UserCacheField = "fb_avatar"
	UserCache_Phone            UserCacheField = "phone"
	UserCache_GcashPhone       UserCacheField = "gcash_phone"
	UserCache_GcashCustomerID  UserCacheField = "gcash_customer_id"
	UserCache_WithdrawPassword UserCacheField = "withdraw_password"
	UserCache_LoginPassword    UserCacheField = "login_password"
	UserCache_PayAccountID     UserCacheField = "pay_account_id"
	UserCache_CardBack         UserCacheField = "card_back"
	UserCache_AvatarFrame      UserCacheField = "avatar_frame"
	UserCache_AppPackageName   UserCacheField = "app_package_name"
	UserCache_Bind             UserCacheField = "bind"
	UserCache_Type             UserCacheField = "type"
	UserCache_InviteUserID     UserCacheField = "invite_user_id"
	UserCache_InviteCode       UserCacheField = "invite_code"
	UserCache_Invite           UserCacheField = "invite"
	UserCache_IsRegister       UserCacheField = "is_register"
	UserCache_Enable           UserCacheField = "enable"
	UserCache_RechargeAmount   UserCacheField = "recharge_amount"
	UserCache_WithdrawAmount   UserCacheField = "withdraw_amount"
	UserCache_AvailableCoins   UserCacheField = "available_coins"
	UserCache_SPlayer          UserCacheField = "s_player"
	UserCache_CPlayer          UserCacheField = "c_player"
	UserCache_WithdrawModel    UserCacheField = "withdraw_model"
	UserCache_WithdrawControl  UserCacheField = "withdraw_control"
	UserCache_TotalRounds      UserCacheField = "total_rounds"
	UserCache_FirstRwReward    UserCacheField = "first_rw_reward"
	UserCache_Level            UserCacheField = "level"
	UserCache_LevelMax         UserCacheField = "level_max"
	UserCache_Score            UserCacheField = "score"
	UserCache_CreatedAt        UserCacheField = "created_at"
	UserCache_Status           UserCacheField = "status"
	UserCache_LastLoginChannel UserCacheField = "last_login_channel"
	UserCache_AppChannel       UserCacheField = "app_channel"
	UserCache_PayAccount       UserCacheField = "pay_account"
	UserCache_PayAccountArr    UserCacheField = "pay_account_arr"
	UserCache_PayAccountPtrArr UserCacheField = "pay_account_ptr_arr"
	UserCache_FloatArr         UserCacheField = "float_arr"
	UserCache_StringArr        UserCacheField = "string_arr"
	UserCache_IntArr           UserCacheField = "int_arr"
	UserCache_UIntArr          UserCacheField = "uint_arr"
	UserCache_Connection       UserCacheField = "connection"
	UserCache_PlayerID         UserCacheField = "player_id"
	UserCache_Token            UserCacheField = "token"
)

func (v *InModel) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["abc"]; ok {
		{
			pv := val
			v.Abc = string(pv)
		}
	}

	return nil
}
func (v *InModel) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["abc"]; ok {
		switch acval := val.(type) {
		case int:
			v.Abc = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Abc = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Abc = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Abc = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Abc = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Abc = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Abc = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Abc = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Abc = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Abc = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Abc = acval
		case float32:
			v.Abc = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Abc = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}

	return nil
}

func (v *InModel) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["abc"] = v.Abc
	return m, nil
}

func (v *InModel) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["abc"] = fmt.Sprint(v.Abc)
	return m, nil
}

func (v *InModel) CheckField(m map[string]string) error {
	if _, ok := m["abc"]; !ok {
		return fmt.Errorf("field abc is miss")
	}
	return nil
}

type InModelField string

func (v InModelField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	InModel_Abc InModelField = "abc"
)

func (v *Resp360) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["fieldInt"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.FieldInt = int(pv)
		}
	}
	if val, ok = m["FieldIntPtr"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			pvv := int(pv)
			v.FieldIntPtr = &pvv
		}
	}
	if val, ok = m["fieldUint8"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.FieldUint8 = uint8(pv)
		}
	}
	if val, ok = m["fieldUint8Ptr"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			pvv := uint8(pv)
			v.FieldUint8Ptr = &pvv
		}
	}
	if val, ok = m["fieldFloat32"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.FieldFloat32 = float32(pv)
		}
	}
	if val, ok = m["fieldFloat32Ptr"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			pvv := float32(pv)
			v.FieldFloat32Ptr = &pvv
		}
	}
	if val, ok = m["fieldBool"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			v.FieldBool = bool(pv)
		}
	}
	if val, ok = m["fieldBoolPtr"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			pvv := bool(pv)
			v.FieldBoolPtr = &pvv
		}
	}
	if val, ok = m["fieldString"]; ok {
		{
			pv := val
			v.FieldString = string(pv)
		}
	}
	if val, ok = m["fieldStringPtr"]; ok {
		{
			pv := val
			pvv := string(pv)
			v.FieldStringPtr = &pvv
		}
	}
	if val, ok = m["decimal"]; ok {
		var err error
		v.Decimal, err = decimal.NewFromString(val)
		if err != nil {
			return err
		}
	}
	if val, ok = m["decimalPtr"]; ok {
		var err error
		var decc decimal.Decimal
		decc, err = decimal.NewFromString(val)
		if err != nil {
			return err
		}
		v.DecimalPtr = &decc
	}

	return nil
}
func (v *Resp360) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["fieldInt"]; ok {
		switch acval := val.(type) {
		case int:
			v.FieldInt = int(acval)
		case int8:
			v.FieldInt = int(acval)
		case int16:
			v.FieldInt = int(acval)
		case int32:
			v.FieldInt = int(acval)
		case int64:
			v.FieldInt = int(acval)
		case uint:
			v.FieldInt = int(acval)
		case uint8:
			v.FieldInt = int(acval)
		case uint16:
			v.FieldInt = int(acval)
		case uint32:
			v.FieldInt = int(acval)
		case uint64:
			v.FieldInt = int(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.FieldInt = int(pvv)
			}
		case float32:
			v.FieldInt = int(acval)
		case float64:
			v.FieldInt = int(acval)
		}
	}
	if val, ok = m["FieldIntPtr"]; ok {
		switch acval := val.(type) {
		case int:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case int8:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case int16:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case int32:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case int64:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case uint:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case uint8:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case uint16:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case uint32:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case uint64:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.FieldIntPtr = easymap_gen.BaseToPtr(int(pvv))
			}
		case float32:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		case float64:
			pvv := int(acval)
			v.FieldIntPtr = &pvv
		}
	}
	if val, ok = m["fieldIntArr"]; ok {
		switch acval := val.(type) {
		case []uint64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []string:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = int(tmp)
			}
			v.FieldIntArr = tmpArr
		case []interface{}:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = int(trueK)
				case int8:
					tmpArr[i] = int(trueK)
				case int16:
					tmpArr[i] = int(trueK)
				case int32:
					tmpArr[i] = int(trueK)
				case int64:
					tmpArr[i] = int(trueK)
				case uint:
					tmpArr[i] = int(trueK)
				case uint8:
					tmpArr[i] = int(trueK)
				case uint16:
					tmpArr[i] = int(trueK)
				case uint32:
					tmpArr[i] = int(trueK)
				case uint64:
					tmpArr[i] = int(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = int(tmp)
				case float32:
					tmpArr[i] = int(trueK)
				case float64:
					tmpArr[i] = int(trueK)
				default:
					break
				}
			}
			v.FieldIntArr = tmpArr
		case []int:
			v.FieldIntArr = acval
		case []uint8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []uint:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []uint16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []uint32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []float32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []int16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []int64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []float64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []int8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		case []int32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntArr = tmpArr
		}
	}
	if val, ok = m["FieldIntArrPtr"]; ok {
		switch acval := val.(type) {
		case []uint:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []uint32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []string:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpA := int(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldIntArrPtr = tmpArr
		case []interface{}:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpA := int(tmp)
					tmpArr[i] = &tmpA
				case float32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				default:
					break
				}
			}
			v.FieldIntArrPtr = tmpArr
		case []float32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []float64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []int16:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []int32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []uint8:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []uint64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []int64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []uint16:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []int:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		case []int8:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntArrPtr = tmpArr
		}
	}
	if val, ok = m["FieldIntPtrArr"]; ok {
		switch acval := val.(type) {
		case []uint:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []uint8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []float64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []interface{}:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = int(trueK)
				case int8:
					tmpArr[i] = int(trueK)
				case int16:
					tmpArr[i] = int(trueK)
				case int32:
					tmpArr[i] = int(trueK)
				case int64:
					tmpArr[i] = int(trueK)
				case uint:
					tmpArr[i] = int(trueK)
				case uint8:
					tmpArr[i] = int(trueK)
				case uint16:
					tmpArr[i] = int(trueK)
				case uint32:
					tmpArr[i] = int(trueK)
				case uint64:
					tmpArr[i] = int(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = int(tmp)
				case float32:
					tmpArr[i] = int(trueK)
				case float64:
					tmpArr[i] = int(trueK)
				default:
					break
				}
			}
			v.FieldIntPtrArr = &tmpArr
		case []float32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []int16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []uint64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []string:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = int(tmp)
			}
			v.FieldIntPtrArr = &tmpArr
		case []int64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []uint16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []uint32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []int:
			v.FieldIntPtrArr = &acval
		case []int8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		case []int32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.FieldIntPtrArr = &tmpArr
		}
	}
	if val, ok = m["FieldIntPtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []int64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []float32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []int:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []int32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []uint64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []uint:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []uint16:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []uint32:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []string:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpA := int(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []interface{}:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpA := int(tmp)
					tmpArr[i] = &tmpA
				case float32:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := int(trueK)
					tmpArr[i] = &tmp
				default:
					break
				}
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []float64:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []int8:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []int16:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		case []uint8:
			tmpArr := make([]*int, len(acval))
			for i, k := range acval {
				tmp := int(k)
				tmpArr[i] = &tmp
			}
			v.FieldIntPtrArrPtr = &tmpArr
		}
	}
	if val, ok = m["fieldUint8"]; ok {
		switch acval := val.(type) {
		case int:
			v.FieldUint8 = uint8(acval)
		case int8:
			v.FieldUint8 = uint8(acval)
		case int16:
			v.FieldUint8 = uint8(acval)
		case int32:
			v.FieldUint8 = uint8(acval)
		case int64:
			v.FieldUint8 = uint8(acval)
		case uint:
			v.FieldUint8 = uint8(acval)
		case uint8:
			v.FieldUint8 = uint8(acval)
		case uint16:
			v.FieldUint8 = uint8(acval)
		case uint32:
			v.FieldUint8 = uint8(acval)
		case uint64:
			v.FieldUint8 = uint8(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.FieldUint8 = uint8(pvv)
			}
		case float32:
			v.FieldUint8 = uint8(acval)
		case float64:
			v.FieldUint8 = uint8(acval)
		}
	}
	if val, ok = m["fieldUint8Ptr"]; ok {
		switch acval := val.(type) {
		case int:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case int8:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case int16:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case int32:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case int64:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case uint:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case uint8:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case uint16:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case uint32:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case uint64:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.FieldUint8Ptr = easymap_gen.BaseToPtr(uint8(pvv))
			}
		case float32:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		case float64:
			pvv := uint8(acval)
			v.FieldUint8Ptr = &pvv
		}
	}
	if val, ok = m["fieldUint8Arr"]; ok {
		switch acval := val.(type) {
		case []uint:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []string:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseUint(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = uint8(tmp)
			}
			v.FieldUint8Arr = tmpArr
		case []int8:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []int64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []uint8:
			v.FieldUint8Arr = acval
		case []uint16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []uint64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []float32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []float64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []int:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []int32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []int16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []uint32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8Arr = tmpArr
		case []interface{}:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = uint8(trueK)
				case int8:
					tmpArr[i] = uint8(trueK)
				case int16:
					tmpArr[i] = uint8(trueK)
				case int32:
					tmpArr[i] = uint8(trueK)
				case int64:
					tmpArr[i] = uint8(trueK)
				case uint:
					tmpArr[i] = uint8(trueK)
				case uint8:
					tmpArr[i] = uint8(trueK)
				case uint16:
					tmpArr[i] = uint8(trueK)
				case uint32:
					tmpArr[i] = uint8(trueK)
				case uint64:
					tmpArr[i] = uint8(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = uint8(tmp)
				case float32:
					tmpArr[i] = uint8(trueK)
				case float64:
					tmpArr[i] = uint8(trueK)
				default:
					break
				}
			}
			v.FieldUint8Arr = tmpArr
		}
	}
	if val, ok = m["fieldUint8ArrPtr"]; ok {
		switch acval := val.(type) {
		case []uint32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []string:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseUint(k, 10, 64)
				if err != nil {
					return err
				}
				tmpA := uint8(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldUint8ArrPtr = tmpArr
		case []int:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []int32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []int64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []uint:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []interface{}:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpA := uint8(tmp)
					tmpArr[i] = &tmpA
				case float32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				default:
					break
				}
			}
			v.FieldUint8ArrPtr = tmpArr
		case []float32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []float64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []uint64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []int8:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []int16:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []uint8:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		case []uint16:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8ArrPtr = tmpArr
		}
	}
	if val, ok = m["fieldUint8PtrArr"]; ok {
		switch acval := val.(type) {
		case []uint16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []uint32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []uint8:
			v.FieldUint8PtrArr = &acval
		case []int64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []uint64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []string:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseUint(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = uint8(tmp)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []int:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []int32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []uint:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []float32:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []float64:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []int8:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		case []interface{}:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = uint8(trueK)
				case int8:
					tmpArr[i] = uint8(trueK)
				case int16:
					tmpArr[i] = uint8(trueK)
				case int32:
					tmpArr[i] = uint8(trueK)
				case int64:
					tmpArr[i] = uint8(trueK)
				case uint:
					tmpArr[i] = uint8(trueK)
				case uint8:
					tmpArr[i] = uint8(trueK)
				case uint16:
					tmpArr[i] = uint8(trueK)
				case uint32:
					tmpArr[i] = uint8(trueK)
				case uint64:
					tmpArr[i] = uint8(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = uint8(tmp)
				case float32:
					tmpArr[i] = uint8(trueK)
				case float64:
					tmpArr[i] = uint8(trueK)
				default:
					break
				}
			}
			v.FieldUint8PtrArr = &tmpArr
		case []int16:
			tmpArr := make([]uint8, len(acval))
			for i, k := range acval {
				tmpArr[i] = uint8(k)
			}
			v.FieldUint8PtrArr = &tmpArr
		}
	}
	if val, ok = m["fieldUint8PtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []int8:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []uint:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []uint8:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []int16:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []int32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []uint32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []uint64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []interface{}:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpA := uint8(tmp)
					tmpArr[i] = &tmpA
				case float32:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := uint8(trueK)
					tmpArr[i] = &tmp
				default:
					break
				}
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []float32:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []string:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseUint(k, 10, 64)
				if err != nil {
					return err
				}
				tmpA := uint8(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []float64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []int:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []int64:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		case []uint16:
			tmpArr := make([]*uint8, len(acval))
			for i, k := range acval {
				tmp := uint8(k)
				tmpArr[i] = &tmp
			}
			v.FieldUint8PtrArrPtr = &tmpArr
		}
	}
	if val, ok = m["fieldFloat32"]; ok {
		switch acval := val.(type) {
		case int:
			v.FieldFloat32 = float32(acval)
		case int8:
			v.FieldFloat32 = float32(acval)
		case int16:
			v.FieldFloat32 = float32(acval)
		case int32:
			v.FieldFloat32 = float32(acval)
		case int64:
			v.FieldFloat32 = float32(acval)
		case uint:
			v.FieldFloat32 = float32(acval)
		case uint8:
			v.FieldFloat32 = float32(acval)
		case uint16:
			v.FieldFloat32 = float32(acval)
		case uint32:
			v.FieldFloat32 = float32(acval)
		case uint64:
			v.FieldFloat32 = float32(acval)
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseFloat(acval, 10)
				if err != nil {
					return err
				}
				v.FieldFloat32 = float32(pvv)
			}
		case float32:
			v.FieldFloat32 = float32(acval)
		case float64:
			v.FieldFloat32 = float32(acval)
		}
	}
	if val, ok = m["fieldFloat32Ptr"]; ok {
		switch acval := val.(type) {
		case int:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case int8:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case int16:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case int32:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case int64:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case uint:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case uint8:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case uint16:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case uint32:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case uint64:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case string:
			if len(acval) == 0 {
				break
			} else {
				pvv, err := strconv.ParseFloat(acval, 10)
				if err != nil {
					return err
				}
				v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(pvv))
			}
		case float32:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		case float64:
			v.FieldFloat32Ptr = easymap_gen.BaseToPtr(float32(acval))
		}
	}
	if val, ok = m["fieldFloat32Arr"]; ok {
		switch acval := val.(type) {
		case []float32:
			if len(acval) == 0 {
				break
			}
			v.FieldFloat32Arr = acval
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = float32(trueK)
				case float64:
					tmpArr[i] = float32(trueK)
				case int:
					tmpArr[i] = float32(trueK)
				case int8:
					tmpArr[i] = float32(trueK)
				case int16:
					tmpArr[i] = float32(trueK)
				case int32:
					tmpArr[i] = float32(trueK)
				case int64:
					tmpArr[i] = float32(trueK)
				case uint:
					tmpArr[i] = float32(trueK)
				case uint8:
					tmpArr[i] = float32(trueK)
				case uint16:
					tmpArr[i] = float32(trueK)
				case uint32:
					tmpArr[i] = float32(trueK)
				case uint64:
					tmpArr[i] = float32(trueK)
				case string:
					tmp, err := strconv.ParseFloat(trueK, 10)
					if err != nil {
						return err
					}
					tmpArr[i] = float32(tmp)
				default:
					break
				}
			}
			v.FieldFloat32Arr = tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpArr[i] = float32(tmp)
			}
			v.FieldFloat32Arr = tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32Arr = tmpArr
		}
	}
	if val, ok = m["fieldFloat32ArrPtr"]; ok {
		switch acval := val.(type) {
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []float32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpA := float32(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldFloat32ArrPtr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseFloat(trueK, 10)
					if err != nil {
						return err
					}
					tmpA := float32(tmp)
					tmpArr[i] = &tmpA
				default:
					break
				}
			}
			v.FieldFloat32ArrPtr = tmpArr
		}
	}
	if val, ok = m["fieldFloat32PtrArr"]; ok {
		switch acval := val.(type) {
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = float32(trueK)
				case float64:
					tmpArr[i] = float32(trueK)
				case int:
					tmpArr[i] = float32(trueK)
				case int8:
					tmpArr[i] = float32(trueK)
				case int16:
					tmpArr[i] = float32(trueK)
				case int32:
					tmpArr[i] = float32(trueK)
				case int64:
					tmpArr[i] = float32(trueK)
				case uint:
					tmpArr[i] = float32(trueK)
				case uint8:
					tmpArr[i] = float32(trueK)
				case uint16:
					tmpArr[i] = float32(trueK)
				case uint32:
					tmpArr[i] = float32(trueK)
				case uint64:
					tmpArr[i] = float32(trueK)
				case string:
					tmp, err := strconv.ParseFloat(trueK, 10)
					if err != nil {
						return err
					}
					tmpArr[i] = float32(tmp)
				default:
					break
				}
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []float32:
			if len(acval) == 0 {
				break
			}
			v.FieldFloat32PtrArr = &acval
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmpArr[i] = float32(k)
			}
			v.FieldFloat32PtrArr = &tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]float32, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpArr[i] = float32(tmp)
			}
			v.FieldFloat32PtrArr = &tmpArr
		}
	}
	if val, ok = m["fieldFloat32PtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []float32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpA := float32(tmp)
				tmpArr[i] = &tmpA
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				tmp := float32(k)
				tmpArr[i] = &tmp
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*float32, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case float64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int8:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int16:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case int64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint8:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint16:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint32:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case uint64:
					tmp := float32(trueK)
					tmpArr[i] = &tmp
				case string:
					tmp, err := strconv.ParseFloat(trueK, 10)
					if err != nil {
						return err
					}
					tmpA := float32(tmp)
					tmpArr[i] = &tmpA
				default:
					break
				}
			}
			v.FieldFloat32PtrArrPtr = &tmpArr
		}
	}
	if val, ok = m["fieldBool"]; ok {
		switch acval := val.(type) {
		case bool:
			v.FieldBool = bool(acval)
		}
	}
	if val, ok = m["fieldBoolPtr"]; ok {
		switch acval := val.(type) {
		case bool:
			pvv := bool(acval)
			v.FieldBoolPtr = &pvv
		}
	}
	if val, ok = m["fieldString"]; ok {
		switch acval := val.(type) {
		case int:
			v.FieldString = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.FieldString = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.FieldString = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.FieldString = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.FieldString = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.FieldString = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.FieldString = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.FieldString = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.FieldString = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.FieldString = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.FieldString = acval
		case float32:
			v.FieldString = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.FieldString = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["fieldStringPtr"]; ok {
		switch acval := val.(type) {
		case int:
			pvv := strconv.FormatInt(int64(acval), 10)
			v.FieldStringPtr = &pvv
		case int8:
			pvv := strconv.FormatInt(int64(acval), 10)
			v.FieldStringPtr = &pvv
		case int16:
			pvv := strconv.FormatInt(int64(acval), 10)
			v.FieldStringPtr = &pvv
		case int32:
			pvv := strconv.FormatInt(int64(acval), 10)
			v.FieldStringPtr = &pvv
		case int64:
			pvv := strconv.FormatInt(int64(acval), 10)
			v.FieldStringPtr = &pvv
		case uint:
			pvv := strconv.FormatUint(uint64(acval), 10)
			v.FieldStringPtr = &pvv
		case uint8:
			pvv := strconv.FormatUint(uint64(acval), 10)
			v.FieldStringPtr = &pvv
		case uint16:
			pvv := strconv.FormatUint(uint64(acval), 10)
			v.FieldStringPtr = &pvv
		case uint32:
			pvv := strconv.FormatUint(uint64(acval), 10)
			v.FieldStringPtr = &pvv
		case uint64:
			pvv := strconv.FormatUint(uint64(acval), 10)
			v.FieldStringPtr = &pvv
		case string:
			pvv := acval
			v.FieldStringPtr = &pvv
		case float32:
			pvv := strconv.FormatFloat(float64(acval), 'f', -1, 64)
			v.FieldStringPtr = &pvv
		case float64:
			pvv := strconv.FormatFloat(float64(acval), 'f', -1, 64)
			v.FieldStringPtr = &pvv
		}
	}
	if val, ok = m["fieldStringArr"]; ok {
		switch acval := val.(type) {
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			v.FieldStringArr = acval
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatFloat(float64(k), 'f', -1, 64)
			}
			v.FieldStringArr = tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case float64:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case int:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int8:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int16:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int32:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int64:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case uint:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint8:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint16:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint32:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint64:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case string:
					tmpArr[i] = trueK
				default:
					break
				}
			}
			v.FieldStringArr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringArr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringArr = tmpArr
		}
	}
	if val, ok = m["fieldStringArrPtr"]; ok {
		switch acval := val.(type) {
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := k
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatFloat(float64(k), 'f', -1, 64)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmp := strconv.FormatFloat(float64(trueK), 'f', -1, 64)
					tmpArr[i] = &tmp
				case float64:
					tmp := strconv.FormatFloat(float64(trueK), 'f', -1, 64)
					tmpArr[i] = &tmp
				case int:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int8:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int16:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int32:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int64:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case uint:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint8:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint16:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint32:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint64:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case string:
					tmpArr[i] = &trueK
				default:
					break
				}
			}
			v.FieldStringArrPtr = tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringArrPtr = tmpArr
		}
	}
	if val, ok = m["fieldStringPtrArr"]; ok {
		switch acval := val.(type) {
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			v.FieldStringPtrArr = &acval
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatFloat(float64(k), 'f', -1, 64)
			}
			v.FieldStringPtrArr = &tmpArr
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case float64:
					tmpArr[i] = strconv.FormatFloat(float64(trueK), 'f', -1, 64)
				case int:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int8:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int16:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int32:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case int64:
					tmpArr[i] = strconv.FormatInt(int64(trueK), 10)
				case uint:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint8:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint16:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint32:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case uint64:
					tmpArr[i] = strconv.FormatUint(uint64(trueK), 10)
				case string:
					tmpArr[i] = trueK
				default:
					break
				}
			}
			v.FieldStringPtrArr = &tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.FieldStringPtrArr = &tmpArr
		}
	}
	if val, ok = m["fieldStringPtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []uint8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []uint32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case float32:
					tmp := strconv.FormatFloat(float64(trueK), 'f', -1, 64)
					tmpArr[i] = &tmp
				case float64:
					tmp := strconv.FormatFloat(float64(trueK), 'f', -1, 64)
					tmpArr[i] = &tmp
				case int:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int8:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int16:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int32:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case int64:
					tmp := strconv.FormatInt(int64(trueK), 10)
					tmpArr[i] = &tmp
				case uint:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint8:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint16:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint32:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case uint64:
					tmp := strconv.FormatUint(uint64(trueK), 10)
					tmpArr[i] = &tmp
				case string:
					tmpArr[i] = &trueK
				default:
					break
				}
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := k
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []int8:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []int64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []uint:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []uint16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []uint64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatUint(uint64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []float64:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatFloat(float64(k), 'f', -1, 64)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []int:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []int16:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		case []int32:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*string, len(acval))
			for i, k := range acval {
				tmp := strconv.FormatInt(int64(k), 10)
				tmpArr[i] = &tmp
			}
			v.FieldStringPtrArrPtr = &tmpArr
		}
	}
	if val, ok = m["decimal"]; ok {
		switch acval := val.(type) {
		case int:
			v.Decimal = decimal.NewFromInt(int64(acval))
		case int8:
			v.Decimal = decimal.NewFromInt(int64(acval))
		case int16:
			v.Decimal = decimal.NewFromInt(int64(acval))
		case int32:
			v.Decimal = decimal.NewFromInt(int64(acval))
		case int64:
			v.Decimal = decimal.NewFromInt(int64(acval))
		case uint:
			v.Decimal = decimal.NewFromUint64(uint64(acval))
		case uint8:
			v.Decimal = decimal.NewFromUint64(uint64(acval))
		case uint16:
			v.Decimal = decimal.NewFromUint64(uint64(acval))
		case uint32:
			v.Decimal = decimal.NewFromUint64(uint64(acval))
		case uint64:
			v.Decimal = decimal.NewFromUint64(uint64(acval))
		case string:
			var err error
			v.Decimal, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
		case float32:
			v.Decimal = decimal.NewFromFloat(float64(acval))
		case float64:
			v.Decimal = decimal.NewFromFloat(float64(acval))
		}
	}
	if val, ok = m["decimalPtr"]; ok {
		switch acval := val.(type) {
		case int:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.DecimalPtr = &decc
		case int8:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.DecimalPtr = &decc
		case int16:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.DecimalPtr = &decc
		case int32:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.DecimalPtr = &decc
		case int64:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.DecimalPtr = &decc
		case uint:
			var decc decimal.Decimal
			decc = decimal.NewFromUint64(uint64(acval))
			v.DecimalPtr = &decc
		case uint8:
			var decc decimal.Decimal
			decc = decimal.NewFromUint64(uint64(acval))
			v.DecimalPtr = &decc
		case uint16:
			var decc decimal.Decimal
			decc = decimal.NewFromUint64(uint64(acval))
			v.DecimalPtr = &decc
		case uint32:
			var decc decimal.Decimal
			decc = decimal.NewFromUint64(uint64(acval))
			v.DecimalPtr = &decc
		case uint64:
			var decc decimal.Decimal
			decc = decimal.NewFromUint64(uint64(acval))
			v.DecimalPtr = &decc
		case string:
			var err error
			var decc decimal.Decimal
			decc, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
			v.DecimalPtr = &decc
		case float32:
			var decc decimal.Decimal
			decc = decimal.NewFromFloat(float64(acval))
			v.DecimalPtr = &decc
		case float64:
			var decc decimal.Decimal
			decc = decimal.NewFromFloat(float64(acval))
			v.DecimalPtr = &decc
		}
	}
	if val, ok = m["decimalArr"]; ok {
		switch acval := val.(type) {
		case []float32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(k))
				}
				v.DecimalArr = decArr
			}
		case []float64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(k))
				}
				v.DecimalArr = decArr
			}
		case []int:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalArr = decArr
			}
		case []int8:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalArr = decArr
			}
		case []int16:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalArr = decArr
			}
		case []int32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalArr = decArr
			}
		case []int64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalArr = decArr
			}
		case []uint:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalArr = decArr
			}
		case []uint8:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalArr = decArr
			}
		case []uint16:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalArr = decArr
			}
		case []uint32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalArr = decArr
			}
		case []uint64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalArr = decArr
			}
		case []string:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp, err := github_com_shopspring_decimal.NewFromString(k)
					if err != nil {
						return err
					}
					decArr[i] = tmp
				}
				v.DecimalArr = decArr
			}
		case []interface{}:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					switch trueK := k.(type) {
					case float32:
						decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(trueK))
					case float64:
						decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(trueK))
					case int:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int8:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int16:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int32:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int64:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case uint:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint8:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint16:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint32:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint64:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case string:
						tmp, err := github_com_shopspring_decimal.NewFromString(trueK)
						if err != nil {
							return err
						}
						decArr[i] = tmp
					}
				}
				v.DecimalArr = decArr
			}
		}
	}
	if val, ok = m["DecimalPtrArr"]; ok {
		switch acval := val.(type) {
		case []float32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []float64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []int:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []int8:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []int16:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []int32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []int64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []uint:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []uint8:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []uint16:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []uint32:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []uint64:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(k))
				}
				v.DecimalPtrArr = &decArr
			}
		case []string:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp, err := github_com_shopspring_decimal.NewFromString(k)
					if err != nil {
						return err
					}
					decArr[i] = tmp
				}
				v.DecimalPtrArr = &decArr
			}
		case []interface{}:
			if len(acval) > 0 {
				decArr := make([]github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					switch trueK := k.(type) {
					case float32:
						decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(trueK))
					case float64:
						decArr[i] = github_com_shopspring_decimal.NewFromFloat(float64(trueK))
					case int:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int8:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int16:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int32:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case int64:
						decArr[i] = github_com_shopspring_decimal.NewFromInt(int64(trueK))
					case uint:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint8:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint16:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint32:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case uint64:
						decArr[i] = github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
					case string:
						tmp, err := github_com_shopspring_decimal.NewFromString(trueK)
						if err != nil {
							return err
						}
						decArr[i] = tmp
					}
				}
				v.DecimalPtrArr = &decArr
			}
		}
	}
	if val, ok = m["decimalArrPtr"]; ok {
		switch acval := val.(type) {
		case []float32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromFloat(float64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []float64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromFloat(float64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []int:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []int8:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []int16:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []int32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []int64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []uint:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []uint8:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []uint16:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []uint32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []uint64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []string:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp, err := github_com_shopspring_decimal.NewFromString(k)
					if err != nil {
						return err
					}
					decArr[i] = &tmp
				}
				v.DecimalArrPtr = decArr
			}
		case []interface{}:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					switch trueK := k.(type) {
					case float32:
						tmp := github_com_shopspring_decimal.NewFromFloat(float64(trueK))
						decArr[i] = &tmp
					case float64:
						tmp := github_com_shopspring_decimal.NewFromFloat(float64(trueK))
						decArr[i] = &tmp
					case int:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int8:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int16:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int32:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int64:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case uint:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint8:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint16:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint32:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint64:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case string:
						tmp, err := github_com_shopspring_decimal.NewFromString(trueK)
						if err != nil {
							return err
						}
						decArr[i] = &tmp
					}
				}
				v.DecimalArrPtr = decArr
			}
		}
	}
	if val, ok = m["DecimalPtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []float32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromFloat(float64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []float64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromFloat(float64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []int:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []int8:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []int16:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []int32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []int64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromInt(int64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []uint:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []uint8:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []uint16:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []uint32:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []uint64:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp := github_com_shopspring_decimal.NewFromUint64(uint64(k))
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []string:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					tmp, err := github_com_shopspring_decimal.NewFromString(k)
					if err != nil {
						return err
					}
					decArr[i] = &tmp
				}
				v.DecimalPtrArrPtr = &decArr
			}
		case []interface{}:
			if len(acval) > 0 {
				decArr := make([]*github_com_shopspring_decimal.Decimal, len(acval))
				for i, k := range acval {
					switch trueK := k.(type) {
					case float32:
						tmp := github_com_shopspring_decimal.NewFromFloat(float64(trueK))
						decArr[i] = &tmp
					case float64:
						tmp := github_com_shopspring_decimal.NewFromFloat(float64(trueK))
						decArr[i] = &tmp
					case int:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int8:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int16:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int32:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case int64:
						tmp := github_com_shopspring_decimal.NewFromInt(int64(trueK))
						decArr[i] = &tmp
					case uint:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint8:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint16:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint32:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case uint64:
						tmp := github_com_shopspring_decimal.NewFromUint64(uint64(trueK))
						decArr[i] = &tmp
					case string:
						tmp, err := github_com_shopspring_decimal.NewFromString(trueK)
						if err != nil {
							return err
						}
						decArr[i] = &tmp
					}
				}
				v.DecimalPtrArrPtr = &decArr
			}
		}
	}
	if val, ok = m["structModel"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &InModel{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.StructModel = *i.(*InModel)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &InModel{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.StructModel = *i.(*InModel)
				}
			}
		case InModel:
			v.StructModel = acval
		case *InModel:
			if acval == nil {
				break
			}
			v.StructModel = *acval
		}

	}
	if val, ok = m["structModelPtr"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &InModel{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.StructModelPtr = i.(*InModel)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &InModel{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.StructModelPtr = i.(*InModel)
				}
			}
		case InModel:
			v.StructModelPtr = &acval
		case *InModel:
			v.StructModelPtr = acval
		}

	}
	if val, ok = m["structModelArr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]InModel, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*InModel)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*InModel)
					}
				case InModel:
					tmpArr[i] = trueK
				case *InModel:
					tmpArr[i] = *trueK
				}
			}
			v.StructModelArr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*InModel)
			}
			v.StructModelArr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*InModel)
			}
			v.StructModelArr = tmpArr

		case []InModel:
			if len(acval) > 0 {
				tmpArr := make([]InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.StructModelArr = tmpArr
			}

		case []*InModel:
			if len(acval) > 0 {
				tmpArr := make([]InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = *k
				}
				v.StructModelArr = tmpArr
			}
		}
	}
	if val, ok = m["structModelArrPtr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*InModel)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*InModel)
					}
				case InModel:
					tmpArr[i] = &trueK
				case *InModel:
					tmpArr[i] = trueK
				}
			}
			v.StructModelArrPtr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*InModel)
			}
			v.StructModelArrPtr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*InModel)
			}
			v.StructModelArrPtr = tmpArr

		case []InModel:
			if len(acval) > 0 {
				tmpArr := make([]*InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = &k
				}
				v.StructModelArrPtr = tmpArr
			}

		case []*InModel:
			if len(acval) > 0 {
				tmpArr := make([]*InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.StructModelArrPtr = tmpArr
			}
		}
	}
	if val, ok = m["structModelPtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*InModel)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &InModel{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*InModel)
					}
				case InModel:
					tmpArr[i] = &trueK
				case *InModel:
					tmpArr[i] = trueK
				}
			}
			v.StructModelPtrArrPtr = &tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*InModel)
			}
			v.StructModelPtrArrPtr = &tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*InModel, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &InModel{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*InModel)
			}
			v.StructModelPtrArrPtr = &tmpArr

		case []InModel:
			if len(acval) > 0 {
				tmpArr := make([]*InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = &k
				}
				v.StructModelPtrArrPtr = &tmpArr
			}

		case []*InModel:
			if len(acval) > 0 {
				tmpArr := make([]*InModel, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.StructModelPtrArrPtr = &tmpArr
			}
		}
	}
	if val, ok = m["outModel"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.OutModel = *i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.OutModel = *i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.OutModel = acval
		case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if acval == nil {
				break
			}
			v.OutModel = *acval
		}

	}
	if val, ok = m["outModelPtr"]; ok {

		switch acval := val.(type) {
		case map[string]interface{}:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(acval); err != nil {
						return err
					}
					v.OutModelPtr = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case map[string]string:
			if len(acval) > 0 {
				var i interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := i.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(acval); err != nil {
						return err
					}
					v.OutModelPtr = i.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
				}
			}
		case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.OutModelPtr = &acval
		case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			v.OutModelPtr = acval
		}

	}
	if val, ok = m["outModelArr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = *kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = trueK
				case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = *trueK
				}
			}
			v.OutModelArr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelArr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = *kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelArr = tmpArr

		case []github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.OutModelArr = tmpArr
			}

		case []*github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = *k
				}
				v.OutModelArr = tmpArr
			}
		}
	}
	if val, ok = m["outModelArrPtr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = &trueK
				case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = trueK
				}
			}
			v.OutModelArrPtr = tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelArrPtr = tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelArrPtr = tmpArr

		case []github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = &k
				}
				v.OutModelArrPtr = tmpArr
			}

		case []*github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.OutModelArrPtr = tmpArr
			}
		}
	}
	if val, ok = m["OutModelPtrArrPtr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case map[string]string:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case map[string]interface{}:
					if len(trueK) == 0 {
						break
					}
					var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(trueK); err != nil {
							return err
						}
						tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
					}
				case github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = &trueK
				case *github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
					tmpArr[i] = trueK
				}
			}
			v.OutModelPtrArrPtr = &tmpArr
		case []map[string]string:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelPtrArrPtr = &tmpArr
		case []map[string]interface{}:
			if len(acval) == 0 {
				break
			}
			tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
			for i, k := range acval {
				if len(k) == 0 {
					continue
				}
				var kjj interface{} = &github_com_coderyw_easymap_gen_test_model1.UnEasyMap{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				tmpArr[i] = kjj.(*github_com_coderyw_easymap_gen_test_model1.UnEasyMap)
			}
			v.OutModelPtrArrPtr = &tmpArr

		case []github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = &k
				}
				v.OutModelPtrArrPtr = &tmpArr
			}

		case []*github_com_coderyw_easymap_gen_test_model1.UnEasyMap:
			if len(acval) > 0 {
				tmpArr := make([]*github_com_coderyw_easymap_gen_test_model1.UnEasyMap, len(acval))
				for i, k := range acval {
					tmpArr[i] = k
				}
				v.OutModelPtrArrPtr = &tmpArr
			}
		}
	}

	return nil
}

func (v *Resp360) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["fieldInt"] = v.FieldInt
	if v.FieldIntPtr != nil {
		m["FieldIntPtr"] = *v.FieldIntPtr
	}
	m["fieldUint8"] = v.FieldUint8
	if v.FieldUint8Ptr != nil {
		m["fieldUint8Ptr"] = *v.FieldUint8Ptr
	}
	m["fieldFloat32"] = v.FieldFloat32
	if v.FieldFloat32Ptr != nil {
		m["fieldFloat32Ptr"] = *v.FieldFloat32Ptr
	}
	m["fieldBool"] = v.FieldBool
	if v.FieldBoolPtr != nil {
		m["fieldBoolPtr"] = *v.FieldBoolPtr
	}
	m["fieldString"] = v.FieldString
	if v.FieldStringPtr != nil {
		m["fieldStringPtr"] = *v.FieldStringPtr
	}
	m["decimal"] = v.Decimal.String()
	m["decimalPtr"] = v.DecimalPtr.String()
	return m, nil
}

func (v *Resp360) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["fieldInt"] = fmt.Sprint(v.FieldInt)
	if v.FieldIntPtr != nil {
		m["FieldIntPtr"] = fmt.Sprint(*v.FieldIntPtr)
	}
	m["fieldUint8"] = fmt.Sprint(v.FieldUint8)
	if v.FieldUint8Ptr != nil {
		m["fieldUint8Ptr"] = fmt.Sprint(*v.FieldUint8Ptr)
	}
	m["fieldFloat32"] = fmt.Sprint(v.FieldFloat32)
	if v.FieldFloat32Ptr != nil {
		m["fieldFloat32Ptr"] = fmt.Sprint(*v.FieldFloat32Ptr)
	}
	m["fieldBool"] = fmt.Sprint(v.FieldBool)
	if v.FieldBoolPtr != nil {
		m["fieldBoolPtr"] = fmt.Sprint(*v.FieldBoolPtr)
	}
	m["fieldString"] = fmt.Sprint(v.FieldString)
	if v.FieldStringPtr != nil {
		m["fieldStringPtr"] = fmt.Sprint(*v.FieldStringPtr)
	}
	m["decimal"] = v.Decimal.String()
	m["decimalPtr"] = v.DecimalPtr.String()
	return m, nil
}

func (v *Resp360) CheckField(m map[string]string) error {
	if _, ok := m["fieldInt"]; !ok {
		return fmt.Errorf("field fieldInt is miss")
	}
	if _, ok := m["FieldIntPtr"]; !ok {
		return fmt.Errorf("field FieldIntPtr is miss")
	}
	if _, ok := m["fieldIntArr"]; !ok {
		return fmt.Errorf("field fieldIntArr is miss")
	}
	if _, ok := m["FieldIntArrPtr"]; !ok {
		return fmt.Errorf("field FieldIntArrPtr is miss")
	}
	if _, ok := m["FieldIntPtrArr"]; !ok {
		return fmt.Errorf("field FieldIntPtrArr is miss")
	}
	if _, ok := m["FieldIntPtrArrPtr"]; !ok {
		return fmt.Errorf("field FieldIntPtrArrPtr is miss")
	}
	if _, ok := m["fieldUint8"]; !ok {
		return fmt.Errorf("field fieldUint8 is miss")
	}
	if _, ok := m["fieldUint8Ptr"]; !ok {
		return fmt.Errorf("field fieldUint8Ptr is miss")
	}
	if _, ok := m["fieldUint8Arr"]; !ok {
		return fmt.Errorf("field fieldUint8Arr is miss")
	}
	if _, ok := m["fieldUint8ArrPtr"]; !ok {
		return fmt.Errorf("field fieldUint8ArrPtr is miss")
	}
	if _, ok := m["fieldUint8PtrArr"]; !ok {
		return fmt.Errorf("field fieldUint8PtrArr is miss")
	}
	if _, ok := m["fieldUint8PtrArrPtr"]; !ok {
		return fmt.Errorf("field fieldUint8PtrArrPtr is miss")
	}
	if _, ok := m["fieldFloat32"]; !ok {
		return fmt.Errorf("field fieldFloat32 is miss")
	}
	if _, ok := m["fieldFloat32Ptr"]; !ok {
		return fmt.Errorf("field fieldFloat32Ptr is miss")
	}
	if _, ok := m["fieldFloat32Arr"]; !ok {
		return fmt.Errorf("field fieldFloat32Arr is miss")
	}
	if _, ok := m["fieldFloat32ArrPtr"]; !ok {
		return fmt.Errorf("field fieldFloat32ArrPtr is miss")
	}
	if _, ok := m["fieldFloat32PtrArr"]; !ok {
		return fmt.Errorf("field fieldFloat32PtrArr is miss")
	}
	if _, ok := m["fieldFloat32PtrArrPtr"]; !ok {
		return fmt.Errorf("field fieldFloat32PtrArrPtr is miss")
	}
	if _, ok := m["fieldBool"]; !ok {
		return fmt.Errorf("field fieldBool is miss")
	}
	if _, ok := m["fieldBoolPtr"]; !ok {
		return fmt.Errorf("field fieldBoolPtr is miss")
	}
	if _, ok := m["fieldBoolArr"]; !ok {
		return fmt.Errorf("field fieldBoolArr is miss")
	}
	if _, ok := m["fieldBoolArrPtr"]; !ok {
		return fmt.Errorf("field fieldBoolArrPtr is miss")
	}
	if _, ok := m["fieldBoolPtrArr"]; !ok {
		return fmt.Errorf("field fieldBoolPtrArr is miss")
	}
	if _, ok := m["fieldBoolPtrArrPtr"]; !ok {
		return fmt.Errorf("field fieldBoolPtrArrPtr is miss")
	}
	if _, ok := m["fieldString"]; !ok {
		return fmt.Errorf("field fieldString is miss")
	}
	if _, ok := m["fieldStringPtr"]; !ok {
		return fmt.Errorf("field fieldStringPtr is miss")
	}
	if _, ok := m["fieldStringArr"]; !ok {
		return fmt.Errorf("field fieldStringArr is miss")
	}
	if _, ok := m["fieldStringArrPtr"]; !ok {
		return fmt.Errorf("field fieldStringArrPtr is miss")
	}
	if _, ok := m["fieldStringPtrArr"]; !ok {
		return fmt.Errorf("field fieldStringPtrArr is miss")
	}
	if _, ok := m["fieldStringPtrArrPtr"]; !ok {
		return fmt.Errorf("field fieldStringPtrArrPtr is miss")
	}
	if _, ok := m["decimal"]; !ok {
		return fmt.Errorf("field decimal is miss")
	}
	if _, ok := m["decimalPtr"]; !ok {
		return fmt.Errorf("field decimalPtr is miss")
	}
	if _, ok := m["decimalArr"]; !ok {
		return fmt.Errorf("field decimalArr is miss")
	}
	if _, ok := m["DecimalPtrArr"]; !ok {
		return fmt.Errorf("field DecimalPtrArr is miss")
	}
	if _, ok := m["decimalArrPtr"]; !ok {
		return fmt.Errorf("field decimalArrPtr is miss")
	}
	if _, ok := m["DecimalPtrArrPtr"]; !ok {
		return fmt.Errorf("field DecimalPtrArrPtr is miss")
	}
	if _, ok := m["structModel"]; !ok {
		return fmt.Errorf("field structModel is miss")
	}
	if _, ok := m["structModelPtr"]; !ok {
		return fmt.Errorf("field structModelPtr is miss")
	}
	if _, ok := m["structModelArr"]; !ok {
		return fmt.Errorf("field structModelArr is miss")
	}
	if _, ok := m["structModelArrPtr"]; !ok {
		return fmt.Errorf("field structModelArrPtr is miss")
	}
	if _, ok := m["structModelPtrArrPtr"]; !ok {
		return fmt.Errorf("field structModelPtrArrPtr is miss")
	}
	if _, ok := m["outModel"]; !ok {
		return fmt.Errorf("field outModel is miss")
	}
	if _, ok := m["outModelPtr"]; !ok {
		return fmt.Errorf("field outModelPtr is miss")
	}
	if _, ok := m["outModelArr"]; !ok {
		return fmt.Errorf("field outModelArr is miss")
	}
	if _, ok := m["outModelArrPtr"]; !ok {
		return fmt.Errorf("field outModelArrPtr is miss")
	}
	if _, ok := m["OutModelPtrArrPtr"]; !ok {
		return fmt.Errorf("field OutModelPtrArrPtr is miss")
	}
	if _, ok := m["inter"]; !ok {
		return fmt.Errorf("field inter is miss")
	}
	if _, ok := m["InterArr"]; !ok {
		return fmt.Errorf("field InterArr is miss")
	}
	if _, ok := m["InterPtrArr"]; !ok {
		return fmt.Errorf("field InterPtrArr is miss")
	}
	if _, ok := m["data"]; !ok {
		return fmt.Errorf("field data is miss")
	}
	return nil
}

type Resp360Field string

func (v Resp360Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	Resp360_FieldInt              Resp360Field = "fieldInt"
	Resp360_FieldIntPtr           Resp360Field = "FieldIntPtr"
	Resp360_FieldIntArr           Resp360Field = "fieldIntArr"
	Resp360_FieldIntArrPtr        Resp360Field = "FieldIntArrPtr"
	Resp360_FieldIntPtrArr        Resp360Field = "FieldIntPtrArr"
	Resp360_FieldIntPtrArrPtr     Resp360Field = "FieldIntPtrArrPtr"
	Resp360_FieldUint8            Resp360Field = "fieldUint8"
	Resp360_FieldUint8Ptr         Resp360Field = "fieldUint8Ptr"
	Resp360_FieldUint8Arr         Resp360Field = "fieldUint8Arr"
	Resp360_FieldUint8ArrPtr      Resp360Field = "fieldUint8ArrPtr"
	Resp360_FieldUint8PtrArr      Resp360Field = "fieldUint8PtrArr"
	Resp360_FieldUint8PtrArrPtr   Resp360Field = "fieldUint8PtrArrPtr"
	Resp360_FieldFloat32          Resp360Field = "fieldFloat32"
	Resp360_FieldFloat32Ptr       Resp360Field = "fieldFloat32Ptr"
	Resp360_FieldFloat32Arr       Resp360Field = "fieldFloat32Arr"
	Resp360_FieldFloat32ArrPtr    Resp360Field = "fieldFloat32ArrPtr"
	Resp360_FieldFloat32PtrArr    Resp360Field = "fieldFloat32PtrArr"
	Resp360_FieldFloat32PtrArrPtr Resp360Field = "fieldFloat32PtrArrPtr"
	Resp360_FieldBool             Resp360Field = "fieldBool"
	Resp360_FieldBoolPtr          Resp360Field = "fieldBoolPtr"
	Resp360_FieldBoolArr          Resp360Field = "fieldBoolArr"
	Resp360_FieldBoolArrPtr       Resp360Field = "fieldBoolArrPtr"
	Resp360_FieldBoolPtrArr       Resp360Field = "fieldBoolPtrArr"
	Resp360_FieldBoolPtrArrPtr    Resp360Field = "fieldBoolPtrArrPtr"
	Resp360_FieldString           Resp360Field = "fieldString"
	Resp360_FieldStringPtr        Resp360Field = "fieldStringPtr"
	Resp360_FieldStringArr        Resp360Field = "fieldStringArr"
	Resp360_FieldStringArrPtr     Resp360Field = "fieldStringArrPtr"
	Resp360_FieldStringPtrArr     Resp360Field = "fieldStringPtrArr"
	Resp360_FieldStringPtrArrPtr  Resp360Field = "fieldStringPtrArrPtr"
	Resp360_Decimal               Resp360Field = "decimal"
	Resp360_DecimalPtr            Resp360Field = "decimalPtr"
	Resp360_DecimalArr            Resp360Field = "decimalArr"
	Resp360_DecimalPtrArr         Resp360Field = "DecimalPtrArr"
	Resp360_DecimalArrPtr         Resp360Field = "decimalArrPtr"
	Resp360_DecimalPtrArrPtr      Resp360Field = "DecimalPtrArrPtr"
	Resp360_StructModel           Resp360Field = "structModel"
	Resp360_StructModelPtr        Resp360Field = "structModelPtr"
	Resp360_StructModelArr        Resp360Field = "structModelArr"
	Resp360_StructModelArrPtr     Resp360Field = "structModelArrPtr"
	Resp360_StructModelPtrArrPtr  Resp360Field = "structModelPtrArrPtr"
	Resp360_OutModel              Resp360Field = "outModel"
	Resp360_OutModelPtr           Resp360Field = "outModelPtr"
	Resp360_OutModelArr           Resp360Field = "outModelArr"
	Resp360_OutModelArrPtr        Resp360Field = "outModelArrPtr"
	Resp360_OutModelPtrArrPtr     Resp360Field = "OutModelPtrArrPtr"
	Resp360_Inter                 Resp360Field = "inter"
	Resp360_InterArr              Resp360Field = "InterArr"
	Resp360_InterPtrArr           Resp360Field = "InterPtrArr"
	Resp360_Data                  Resp360Field = "data"
)
