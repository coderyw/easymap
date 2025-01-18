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
	str := `{"id":"41569190421422085","user_id":"41569190421422085","nickname":"G41569190421422085","avatar":"1","fb_avatar":"","phone":"9550031347","gcash_phone":"","gcash_customer_id":"","withdraw_password":0,"login_password":0,"pay_account_id":"","card_back":0,"avatar_frame":0,"app_package_name":"com.playmate.playzone","bind":[1],"type":"16384","invite_user_id":"","invite_code":"41569190421422085","invite":{},"is_register":1,"enable":1,"recharge_amount":0,"withdraw_amount":0,"s_player":0,"c_player":0,"withdraw_model":2,"total_rounds":0,"first_rw_reward":false,"level":0,"level_max":0,"score":0,"created_at":1735871730,"status":null,"last_login_channel":"Web","grayscale":0,"app_channel":"Web","is_vip":"0","vip_start":"0","vip_end":"0","pay_account":{"email":"hcg@rummy.com","phone":"9550031347","name":"G41569190421422085"},"pay_account_arr":[{"email":"hcg@rummy.com","phone":"9550031347","name":"G41569190421422085"}],"pay_account_ptr_arr":[{"email":"hcg@rummy.com","phone":"9550031347","name":"G41569190421422085"}],"player_id":"41569190421422085","token":"41569190421422085:1735871730:0:cf77e51182a430d8596c1d076c843a47:env-TEST","test_arr":[{"a":1}]}
`
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
