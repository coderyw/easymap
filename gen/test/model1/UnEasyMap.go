// Package model1
// @Author: yinwei
// @File: UnEasyMap
// @Version: 1.0.0
// @Date: 2024/3/13 16:19

package model1

type UnEasyMap struct {
	A int `json:"a"`
}

func (u *UnEasyMap) UnMarshalMapInterface(m map[string]interface{}) error {
	u.A = 1
	return nil
}

func (u UnEasyMap) UnMarshalMapInterface1(m map[string]interface{}) error {
	b := &u
	b.A = 1
	return nil
}
