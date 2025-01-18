package model

import (
	fmt "fmt"
	easy_facade "github.com/coderyw/easymap/gen/facade"
	decimal "github.com/shopspring/decimal"
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
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
		case uint8:
			v.UserId = int64(acval)
		case uint16:
			v.UserId = int64(acval)
		case uint64:
			v.UserId = int64(acval)
		case int16:
			v.UserId = int64(acval)
		case int64:
			v.UserId = int64(acval)
		case int32:
			v.UserId = int64(acval)
		case uint:
			v.UserId = int64(acval)
		case uint32:
			v.UserId = int64(acval)
		case string:
			if len(acval) == 0 {
				v.UserId = 0
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
		case int:
			v.UserId = int64(acval)
		case int8:
			v.UserId = int64(acval)
		}
	}
	if val, ok = m["nickname"]; ok {
		switch acval := val.(type) {
		case string:
			v.Nickname = acval
		case int:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float64:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int8:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["avatar"]; ok {
		switch acval := val.(type) {
		case string:
			v.Avatar = acval
		case int8:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		switch acval := val.(type) {
		case string:
			v.FbAvatar = acval
		case int8:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite_code"]; ok {
		switch acval := val.(type) {
		case float64:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int32:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.InviteCode = acval
		case int:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int32:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.Phone = acval
		case int:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Phone = strconv.FormatInt(int64(acval), 10)
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
		case uint:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Email = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.Email = acval
		case int8:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Email = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Email = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Email = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case string:
			v.Phone = acval
		case int:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int8:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Phone = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["name"]; ok {
		switch acval := val.(type) {
		case uint32:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Name = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.Name = acval
		case int16:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Name = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Name = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.Name = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Name = strconv.FormatUint(uint64(acval), 10)
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
		case int64:
			v.ID = int64(acval)
		case uint:
			v.ID = int64(acval)
		case uint8:
			v.ID = int64(acval)
		case uint32:
			v.ID = int64(acval)
		case float32:
			v.ID = int64(acval)
		case int:
			v.ID = int64(acval)
		case int16:
			v.ID = int64(acval)
		case int32:
			v.ID = int64(acval)
		case string:
			if len(acval) == 0 {
				v.ID = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.ID = int64(pvv)
			}
		case float64:
			v.ID = int64(acval)
		case int8:
			v.ID = int64(acval)
		case uint16:
			v.ID = int64(acval)
		case uint64:
			v.ID = int64(acval)
		}
	}
	if val, ok = m["user_id"]; ok {
		switch acval := val.(type) {
		case uint64:
			v.UserID = int64(acval)
		case string:
			if len(acval) == 0 {
				v.UserID = 0
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
		case int:
			v.UserID = int64(acval)
		case int8:
			v.UserID = int64(acval)
		case int16:
			v.UserID = int64(acval)
		case int32:
			v.UserID = int64(acval)
		case uint32:
			v.UserID = int64(acval)
		case int64:
			v.UserID = int64(acval)
		case uint:
			v.UserID = int64(acval)
		case uint8:
			v.UserID = int64(acval)
		case uint16:
			v.UserID = int64(acval)
		}
	}
	if val, ok = m["balance"]; ok {
		switch acval := val.(type) {
		case uint16:
			v.Balance = int64(acval)
		case float64:
			v.Balance = int64(acval)
		case int:
			v.Balance = int64(acval)
		case int64:
			v.Balance = int64(acval)
		case uint8:
			v.Balance = int64(acval)
		case uint:
			v.Balance = int64(acval)
		case uint32:
			v.Balance = int64(acval)
		case uint64:
			v.Balance = int64(acval)
		case string:
			if len(acval) == 0 {
				v.Balance = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Balance = int64(pvv)
			}
		case float32:
			v.Balance = int64(acval)
		case int8:
			v.Balance = int64(acval)
		case int16:
			v.Balance = int64(acval)
		case int32:
			v.Balance = int64(acval)
		}
	}
	if val, ok = m["nickname"]; ok {
		switch acval := val.(type) {
		case string:
			v.Nickname = acval
		case int16:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint32:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case uint8:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case int:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Nickname = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Nickname = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Nickname = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["avatar"]; ok {
		switch acval := val.(type) {
		case uint64:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int64:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint32:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case int8:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Avatar = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Avatar = acval
		case int:
			v.Avatar = strconv.FormatInt(int64(acval), 10)
		case float32:
			v.Avatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["fb_avatar"]; ok {
		switch acval := val.(type) {
		case float32:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.FbAvatar = acval
		case int8:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.FbAvatar = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.FbAvatar = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.FbAvatar = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["phone"]; ok {
		switch acval := val.(type) {
		case uint:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.Phone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case uint8:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.Phone = acval
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
		case uint64:
			v.Phone = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["gcash_phone"]; ok {
		switch acval := val.(type) {
		case int:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.GcashPhone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case uint:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.GcashPhone = acval
		case int8:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.GcashPhone = strconv.FormatInt(int64(acval), 10)
		case uint64:
			v.GcashPhone = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.GcashPhone = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["gcash_customer_id"]; ok {
		switch acval := val.(type) {
		case uint:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case int:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case float64:
			v.GcashCustomerID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.GcashCustomerID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.GcashCustomerID = acval
		case int64:
			v.GcashCustomerID = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.GcashCustomerID = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["withdraw_password"]; ok {
		switch acval := val.(type) {
		case float32:
			v.WithdrawPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.WithdrawPassword = acval
		case int:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case uint32:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.WithdrawPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int32:
			v.WithdrawPassword = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.WithdrawPassword = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["login_password"]; ok {
		switch acval := val.(type) {
		case int8:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.LoginPassword = acval
		case int:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.LoginPassword = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.LoginPassword = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.LoginPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.LoginPassword = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["pay_account_id"]; ok {
		switch acval := val.(type) {
		case uint8:
			v.PayAccountID = int64(acval)
		case uint16:
			v.PayAccountID = int64(acval)
		case uint64:
			v.PayAccountID = int64(acval)
		case string:
			if len(acval) == 0 {
				v.PayAccountID = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.PayAccountID = int64(pvv)
			}
		case uint:
			v.PayAccountID = int64(acval)
		case uint32:
			v.PayAccountID = int64(acval)
		case float32:
			v.PayAccountID = int64(acval)
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
		case float64:
			v.PayAccountID = int64(acval)
		}
	}
	if val, ok = m["card_back"]; ok {
		switch acval := val.(type) {
		case int8:
			v.CardBack = int(acval)
		case int16:
			v.CardBack = int(acval)
		case int64:
			v.CardBack = int(acval)
		case uint8:
			v.CardBack = int(acval)
		case float32:
			v.CardBack = int(acval)
		case float64:
			v.CardBack = int(acval)
		case int:
			v.CardBack = int(acval)
		case int32:
			v.CardBack = int(acval)
		case uint:
			v.CardBack = int(acval)
		case uint16:
			v.CardBack = int(acval)
		case uint32:
			v.CardBack = int(acval)
		case uint64:
			v.CardBack = int(acval)
		case string:
			if len(acval) == 0 {
				v.CardBack = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CardBack = int(pvv)
			}
		}
	}
	if val, ok = m["avatar_frame"]; ok {
		switch acval := val.(type) {
		case string:
			if len(acval) == 0 {
				v.AvatarFrame = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.AvatarFrame = int(pvv)
			}
		case float64:
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
		case float32:
			v.AvatarFrame = int(acval)
		case int:
			v.AvatarFrame = int(acval)
		case int8:
			v.AvatarFrame = int(acval)
		case uint16:
			v.AvatarFrame = int(acval)
		case uint32:
			v.AvatarFrame = int(acval)
		case uint64:
			v.AvatarFrame = int(acval)
		}
	}
	if val, ok = m["app_package_name"]; ok {
		switch acval := val.(type) {
		case string:
			v.AppPackageName = acval
		case int:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case uint32:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.AppPackageName = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case uint64:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.AppPackageName = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int8:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.AppPackageName = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.AppPackageName = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["bind"]; ok {
		switch acval := val.(type) {
		case []int64:
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
		case []uint16:
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
		case []uint64:
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
		case []int:
			v.Bind = acval
		case []float32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.Bind = tmpArr
		case []int8:
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
		case []float64:
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
		}
	}
	if val, ok = m["type"]; ok {
		switch acval := val.(type) {
		case int32:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Type = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.Type = acval
		case int8:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Type = strconv.FormatUint(uint64(acval), 10)
		case int:
			v.Type = strconv.FormatInt(int64(acval), 10)
		case float32:
			v.Type = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["invite_user_id"]; ok {
		switch acval := val.(type) {
		case float32:
			v.InviteUserID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.InviteUserID = acval
		case int32:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.InviteUserID = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case uint:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.InviteUserID = strconv.FormatUint(uint64(acval), 10)
		case int:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.InviteUserID = strconv.FormatInt(int64(acval), 10)
		}
	}
	if val, ok = m["invite_code"]; ok {
		switch acval := val.(type) {
		case uint64:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int8:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case int64:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.InviteCode = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.InviteCode = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case string:
			v.InviteCode = acval
		case int:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.InviteCode = strconv.FormatInt(int64(acval), 10)
		}
	}
	if val, ok = m["invite"]; ok {
		if m1, ok := val.(map[string]interface{}); ok && len(m1) > 0 {
			var i interface{} = &UserTinyInfo{}
			if b, ok := i.(easy_facade.EasyMapInter); ok {
				if err := b.UnMarshalMapInterface(m1); err != nil {
					return err
				}
				v.Invite = i.(*UserTinyInfo)
			}
		} else if m2, ok := val.(map[string]string); ok && len(m1) > 0 {
			var i interface{} = &UserTinyInfo{}
			if b, ok := i.(easy_facade.EasyMapString); ok {
				if err := b.UnMarshalMap(m2); err != nil {
					return err
				}
				v.Invite = i.(*UserTinyInfo)
			}
		}
	}
	if val, ok = m["is_register"]; ok {
		switch acval := val.(type) {
		case uint32:
			v.IsRegister = uint8(acval)
		case float64:
			v.IsRegister = uint8(acval)
		case int16:
			v.IsRegister = uint8(acval)
		case uint:
			v.IsRegister = uint8(acval)
		case uint16:
			v.IsRegister = uint8(acval)
		case int64:
			v.IsRegister = uint8(acval)
		case uint8:
			v.IsRegister = uint8(acval)
		case uint64:
			v.IsRegister = uint8(acval)
		case string:
			if len(acval) == 0 {
				v.IsRegister = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.IsRegister = uint8(pvv)
			}
		case float32:
			v.IsRegister = uint8(acval)
		case int:
			v.IsRegister = uint8(acval)
		case int8:
			v.IsRegister = uint8(acval)
		case int32:
			v.IsRegister = uint8(acval)
		}
	}
	if val, ok = m["enable"]; ok {
		switch acval := val.(type) {
		case uint8:
			v.Enable = uint8(acval)
		case uint16:
			v.Enable = uint8(acval)
		case int:
			v.Enable = uint8(acval)
		case int8:
			v.Enable = uint8(acval)
		case int16:
			v.Enable = uint8(acval)
		case int64:
			v.Enable = uint8(acval)
		case uint:
			v.Enable = uint8(acval)
		case float64:
			v.Enable = uint8(acval)
		case int32:
			v.Enable = uint8(acval)
		case uint32:
			v.Enable = uint8(acval)
		case uint64:
			v.Enable = uint8(acval)
		case string:
			if len(acval) == 0 {
				v.Enable = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Enable = uint8(pvv)
			}
		case float32:
			v.Enable = uint8(acval)
		}
	}
	if val, ok = m["recharge_amount"]; ok {
		switch acval := val.(type) {
		case string:
			if len(acval) == 0 {
				v.RechargeAmount = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.RechargeAmount = int64(pvv)
			}
		case float64:
			v.RechargeAmount = int64(acval)
		case int8:
			v.RechargeAmount = int64(acval)
		case int64:
			v.RechargeAmount = int64(acval)
		case uint16:
			v.RechargeAmount = int64(acval)
		case uint32:
			v.RechargeAmount = int64(acval)
		case uint8:
			v.RechargeAmount = int64(acval)
		case uint64:
			v.RechargeAmount = int64(acval)
		case float32:
			v.RechargeAmount = int64(acval)
		case int:
			v.RechargeAmount = int64(acval)
		case int16:
			v.RechargeAmount = int64(acval)
		case int32:
			v.RechargeAmount = int64(acval)
		case uint:
			v.RechargeAmount = int64(acval)
		}
	}
	if val, ok = m["withdraw_amount"]; ok {
		switch acval := val.(type) {
		case int16:
			v.WithdrawAmount = int64(acval)
		case int32:
			v.WithdrawAmount = int64(acval)
		case uint:
			v.WithdrawAmount = int64(acval)
		case uint8:
			v.WithdrawAmount = int64(acval)
		case uint32:
			v.WithdrawAmount = int64(acval)
		case uint64:
			v.WithdrawAmount = int64(acval)
		case int:
			v.WithdrawAmount = int64(acval)
		case int8:
			v.WithdrawAmount = int64(acval)
		case string:
			if len(acval) == 0 {
				v.WithdrawAmount = 0
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
		case int64:
			v.WithdrawAmount = int64(acval)
		case uint16:
			v.WithdrawAmount = int64(acval)
		}
	}
	if val, ok = m["available_coins"]; ok {
		switch acval := val.(type) {
		case int32:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint16:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint32:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int8:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int16:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint64:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case int64:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case uint8:
			v.AvailableCoins = decimal.NewFromInt(int64(acval))
		case float32:
			v.AvailableCoins = decimal.NewFromFloat(float64(acval))
		case float64:
			v.AvailableCoins = decimal.NewFromFloat(float64(acval))
		case string:
			var err error
			v.AvailableCoins, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
		}
	}
	if val, ok = m["s_player"]; ok {
		switch acval := val.(type) {
		case int64:
			v.SPlayer = uint8(acval)
		case uint:
			v.SPlayer = uint8(acval)
		case uint8:
			v.SPlayer = uint8(acval)
		case uint16:
			v.SPlayer = uint8(acval)
		case int:
			v.SPlayer = uint8(acval)
		case int8:
			v.SPlayer = uint8(acval)
		case int16:
			v.SPlayer = uint8(acval)
		case int32:
			v.SPlayer = uint8(acval)
		case uint64:
			v.SPlayer = uint8(acval)
		case string:
			if len(acval) == 0 {
				v.SPlayer = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.SPlayer = uint8(pvv)
			}
		case float64:
			v.SPlayer = uint8(acval)
		case uint32:
			v.SPlayer = uint8(acval)
		case float32:
			v.SPlayer = uint8(acval)
		}
	}
	if val, ok = m["c_player"]; ok {
		switch acval := val.(type) {
		case uint32:
			v.CPlayer = uint8(acval)
		case uint64:
			v.CPlayer = uint8(acval)
		case string:
			if len(acval) == 0 {
				v.CPlayer = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CPlayer = uint8(pvv)
			}
		case int:
			v.CPlayer = uint8(acval)
		case int16:
			v.CPlayer = uint8(acval)
		case int64:
			v.CPlayer = uint8(acval)
		case uint:
			v.CPlayer = uint8(acval)
		case uint8:
			v.CPlayer = uint8(acval)
		case float32:
			v.CPlayer = uint8(acval)
		case int8:
			v.CPlayer = uint8(acval)
		case int32:
			v.CPlayer = uint8(acval)
		case uint16:
			v.CPlayer = uint8(acval)
		case float64:
			v.CPlayer = uint8(acval)
		}
	}
	if val, ok = m["withdraw_model"]; ok {
		switch acval := val.(type) {
		case uint64:
			v.WithdrawModel = int(acval)
		case int:
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
		case float64:
			v.WithdrawModel = int(acval)
		case int8:
			v.WithdrawModel = int(acval)
		case int16:
			v.WithdrawModel = int(acval)
		case string:
			if len(acval) == 0 {
				v.WithdrawModel = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.WithdrawModel = int(pvv)
			}
		case float32:
			v.WithdrawModel = int(acval)
		}
	}
	if val, ok = m["withdraw_control"]; ok {
		switch acval := val.(type) {
		case int32:
			v.WithdrawControl = int(acval)
		case uint16:
			v.WithdrawControl = int(acval)
		case uint32:
			v.WithdrawControl = int(acval)
		case float64:
			v.WithdrawControl = int(acval)
		case uint8:
			v.WithdrawControl = int(acval)
		case uint64:
			v.WithdrawControl = int(acval)
		case string:
			if len(acval) == 0 {
				v.WithdrawControl = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.WithdrawControl = int(pvv)
			}
		case int:
			v.WithdrawControl = int(acval)
		case int8:
			v.WithdrawControl = int(acval)
		case int16:
			v.WithdrawControl = int(acval)
		case int64:
			v.WithdrawControl = int(acval)
		case uint:
			v.WithdrawControl = int(acval)
		case float32:
			v.WithdrawControl = int(acval)
		}
	}
	if val, ok = m["total_rounds"]; ok {
		switch acval := val.(type) {
		case int16:
			v.TotalRounds = int(acval)
		case uint:
			v.TotalRounds = int(acval)
		case uint8:
			v.TotalRounds = int(acval)
		case uint64:
			v.TotalRounds = int(acval)
		case string:
			if len(acval) == 0 {
				v.TotalRounds = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.TotalRounds = int(pvv)
			}
		case float32:
			v.TotalRounds = int(acval)
		case int:
			v.TotalRounds = int(acval)
		case int8:
			v.TotalRounds = int(acval)
		case float64:
			v.TotalRounds = int(acval)
		case uint16:
			v.TotalRounds = int(acval)
		case uint32:
			v.TotalRounds = int(acval)
		case int32:
			v.TotalRounds = int(acval)
		case int64:
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
		case uint:
			v.Level = int(acval)
		case uint8:
			v.Level = int(acval)
		case uint16:
			v.Level = int(acval)
		case uint64:
			v.Level = int(acval)
		case string:
			if len(acval) == 0 {
				v.Level = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Level = int(pvv)
			}
		case int8:
			v.Level = int(acval)
		case int16:
			v.Level = int(acval)
		case int32:
			v.Level = int(acval)
		case float32:
			v.Level = int(acval)
		case float64:
			v.Level = int(acval)
		case int:
			v.Level = int(acval)
		case int64:
			v.Level = int(acval)
		case uint32:
			v.Level = int(acval)
		}
	}
	if val, ok = m["level_max"]; ok {
		switch acval := val.(type) {
		case int16:
			v.LevelMax = int(acval)
		case int32:
			v.LevelMax = int(acval)
		case int64:
			v.LevelMax = int(acval)
		case uint8:
			v.LevelMax = int(acval)
		case uint16:
			v.LevelMax = int(acval)
		case int:
			v.LevelMax = int(acval)
		case int8:
			v.LevelMax = int(acval)
		case uint:
			v.LevelMax = int(acval)
		case uint32:
			v.LevelMax = int(acval)
		case uint64:
			v.LevelMax = int(acval)
		case string:
			if len(acval) == 0 {
				v.LevelMax = 0
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
		case int16:
			v.Score = int(acval)
		case uint:
			v.Score = int(acval)
		case uint16:
			v.Score = int(acval)
		case int64:
			v.Score = int(acval)
		case uint8:
			v.Score = int(acval)
		case uint32:
			v.Score = int(acval)
		case uint64:
			v.Score = int(acval)
		case string:
			if len(acval) == 0 {
				v.Score = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Score = int(pvv)
			}
		case int:
			v.Score = int(acval)
		case int8:
			v.Score = int(acval)
		case int32:
			v.Score = int(acval)
		case float32:
			v.Score = int(acval)
		case float64:
			v.Score = int(acval)
		}
	}
	if val, ok = m["created_at"]; ok {
		switch acval := val.(type) {
		case uint8:
			v.CreatedAt = int(acval)
		case uint32:
			v.CreatedAt = int(acval)
		case uint64:
			v.CreatedAt = int(acval)
		case string:
			if len(acval) == 0 {
				v.CreatedAt = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.CreatedAt = int(pvv)
			}
		case int8:
			v.CreatedAt = int(acval)
		case int16:
			v.CreatedAt = int(acval)
		case int32:
			v.CreatedAt = int(acval)
		case uint:
			v.CreatedAt = int(acval)
		case float64:
			v.CreatedAt = int(acval)
		case int:
			v.CreatedAt = int(acval)
		case int64:
			v.CreatedAt = int(acval)
		case uint16:
			v.CreatedAt = int(acval)
		case float32:
			v.CreatedAt = int(acval)
		}
	}
	if val, ok = m["status"]; ok {
		switch acval := val.(type) {
		case int32:
			v.Status = int8(acval)
		case int64:
			v.Status = int8(acval)
		case uint:
			v.Status = int8(acval)
		case uint16:
			v.Status = int8(acval)
		case uint64:
			v.Status = int8(acval)
		case int:
			v.Status = int8(acval)
		case int8:
			v.Status = int8(acval)
		case uint32:
			v.Status = int8(acval)
		case string:
			if len(acval) == 0 {
				v.Status = 0
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
		case int16:
			v.Status = int8(acval)
		case uint8:
			v.Status = int8(acval)
		}
	}
	if val, ok = m["last_login_channel"]; ok {
		switch acval := val.(type) {
		case uint:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint8:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		case string:
			v.LastLoginChannel = acval
		case int:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int16:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case float64:
			v.LastLoginChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.LastLoginChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int8:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.LastLoginChannel = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.LastLoginChannel = strconv.FormatUint(uint64(acval), 10)
		}
	}
	if val, ok = m["app_channel"]; ok {
		switch acval := val.(type) {
		case int8:
			v.AppChannel = strconv.FormatInt(int64(acval), 10)
		case uint16:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.AppChannel = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.AppChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
			v.AppChannel = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
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
		case string:
			v.AppChannel = acval
		}
	}
	if val, ok = m["pay_account"]; ok {
		if m1, ok := val.(map[string]interface{}); ok && len(m1) > 0 {
			var i interface{} = &PayAccount{}
			if b, ok := i.(easy_facade.EasyMapInter); ok {
				if err := b.UnMarshalMapInterface(m1); err != nil {
					return err
				}
				v.PayAccount = *i.(*PayAccount)
			}
		} else if m2, ok := val.(map[string]string); ok && len(m1) > 0 {
			var i interface{} = &PayAccount{}
			if b, ok := i.(easy_facade.EasyMapString); ok {
				if err := b.UnMarshalMap(m2); err != nil {
					return err
				}
				v.PayAccount = *i.(*PayAccount)
			}
		}
	}
	if val, ok = m["pay_account_arr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			v.PayAccountArr = make([]PayAccount, len(acval))
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
						v.PayAccountArr[i] = *kjj.(*PayAccount)
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
						v.PayAccountArr[i] = *kjj.(*PayAccount)
					}
				}
			}
		case []map[string]string:
			v.PayAccountArr = make([]PayAccount, len(acval))
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
				v.PayAccountArr[i] = *kjj.(*PayAccount)
			}
		case []map[string]interface{}:
			v.PayAccountArr = make([]PayAccount, len(acval))
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
				v.PayAccountArr[i] = *kjj.(*PayAccount)
			}
		}
	}
	if val, ok = m["pay_account_ptr_arr"]; ok {
		switch acval := val.(type) {
		case []interface{}:
			v.PayAccountPtrArr = make([]*PayAccount, len(acval))
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
						v.PayAccountPtrArr[i] = kjj.(*PayAccount)
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
						v.PayAccountPtrArr[i] = kjj.(*PayAccount)
					}
				}
			}
		case []map[string]string:
			v.PayAccountPtrArr = make([]*PayAccount, len(acval))
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
				v.PayAccountPtrArr[i] = kjj.(*PayAccount)
			}
		case []map[string]interface{}:
			v.PayAccountPtrArr = make([]*PayAccount, len(acval))
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
				v.PayAccountPtrArr[i] = kjj.(*PayAccount)
			}
		}
	}
	if val, ok = m["float_arr"]; ok {
		switch acval := val.(type) {
		case []float32:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint8:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint32:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint64:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int32:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int64:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []interface{}:
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
		case []uint:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []uint16:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []string:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseFloat(k, 10)
				if err != nil {
					return err
				}
				tmpArr[i] = float64(tmp)
			}
			v.FloatArr = tmpArr
		case []float64:
			v.FloatArr = acval
		case []int8:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		case []int16:
			tmpArr := make([]float64, len(acval))
			for i, k := range acval {
				tmpArr[i] = float64(k)
			}
			v.FloatArr = tmpArr
		}
	}
	if val, ok = m["string_arr"]; ok {
		switch acval := val.(type) {
		case []uint8:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint16:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []string:
			v.StringArr = acval
		case []int8:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []int16:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint32:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []uint64:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatUint(uint64(k), 10)
			}
			v.StringArr = tmpArr
		case []interface{}:
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
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatFloat(float64(k), 'f', -1, 64)
			}
			v.StringArr = tmpArr
		case []int:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []int32:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		case []int64:
			tmpArr := make([]string, len(acval))
			for i, k := range acval {
				tmpArr[i] = strconv.FormatInt(int64(k), 10)
			}
			v.StringArr = tmpArr
		}
	}
	if val, ok = m["int_arr"]; ok {
		switch acval := val.(type) {
		case []uint:
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
		case []float32:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []int:
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
		case []uint16:
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
		case []int8:
			tmpArr := make([]int32, len(acval))
			for i, k := range acval {
				tmpArr[i] = int32(k)
			}
			v.IntArr = tmpArr
		case []int32:
			v.IntArr = acval
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
		case []int16:
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
		}
	}
	if val, ok = m["uint_arr"]; ok {
		switch acval := val.(type) {
		case []int16:
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
		case []uint8:
			v.UIntArr = acval
		case []int:
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
		case []float64:
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
		case []int32:
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
		case []uint16:
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
		}
	}
	if val, ok = m["connection"]; ok {
		if m1, ok := val.(map[string]interface{}); ok && len(m1) > 0 {
			var i interface{} = &Connection{}
			if b, ok := i.(easy_facade.EasyMapInter); ok {
				if err := b.UnMarshalMapInterface(m1); err != nil {
					return err
				}
				v.Connection = *i.(*Connection)
			}
		} else if m2, ok := val.(map[string]string); ok && len(m1) > 0 {
			var i interface{} = &Connection{}
			if b, ok := i.(easy_facade.EasyMapString); ok {
				if err := b.UnMarshalMap(m2); err != nil {
					return err
				}
				v.Connection = *i.(*Connection)
			}
		}
	}
	if val, ok = m["player_id"]; ok {
		switch acval := val.(type) {
		case uint16:
			v.PlayerID = int64(acval)
		case int:
			v.PlayerID = int64(acval)
		case int32:
			v.PlayerID = int64(acval)
		case uint:
			v.PlayerID = int64(acval)
		case uint8:
			v.PlayerID = int64(acval)
		case uint64:
			v.PlayerID = int64(acval)
		case string:
			if len(acval) == 0 {
				v.PlayerID = 0
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
		case int8:
			v.PlayerID = int64(acval)
		case int16:
			v.PlayerID = int64(acval)
		case int64:
			v.PlayerID = int64(acval)
		case uint32:
			v.PlayerID = int64(acval)
		}
	}
	if val, ok = m["token"]; ok {
		switch acval := val.(type) {
		case string:
			v.Token = acval
		case int:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case int16:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Token = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Token = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Token = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case float32:
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
