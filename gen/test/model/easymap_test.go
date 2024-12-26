// Package model
// @Author: yinwei
// @File: easymap_test
// @Version: 1.0.0
// @Date: 2024/12/18 18:03

package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_UnMarshalMap(t *testing.T) {
	str := `{"id":"35811495334178821","bind":[1,2,3],"user_id":"35811495334178821","nickname":"G35811495334178821","avatar":"1","fb_avatar":"","phone":"9103005002","gcash_phone":"","gcash_customer_id":"","withdraw_password":0,"login_password":123,"pay_account_id":"","card_back":0,"avatar_frame":0,"app_package_name":"com.playmate.playzone","bind":[1],"type":"16384","invite_user_id":"","invite_code":"35811495334178821","invite":[],"is_register":0,"enable":1,"recharge_amount":0,"withdraw_amount":0,"s_player":0,"c_player":0,"withdraw_model":2,"total_rounds":0,"first_rw_reward":false,"level":0,"level_max":0,"score":0,"created_at":1734498988,"status":null,"last_login_channel":"Maya","grayscale":0,"app_channel":"Maya","is_vip":"","vip_start":"0","vip_end":"0","player_id":"35811495334178821","token":"35811495334178821:1734499218:0:ee630f29d3aa3bdfc82f6d8f31068144:env-PRE","pay_account":{"email":"hcg@rummy.com","phone":"9103005002","name":"G35811495334178821"}}`
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	bu := &UserCache{}
	err = bu.UnMarshalMapInterface(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", bu)
}
