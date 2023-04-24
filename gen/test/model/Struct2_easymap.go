package model

import (
	unsafe "unsafe"
)

func str2Bytes_Struct2(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *Struct2) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["DD"]; ok {
		v.DD = val
	}

	return nil
}

func (v *Struct2) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["DD"] = v.DD
	return m, nil
}

type Struct2Field string

func (v Struct2Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_Struct2(string(v)), nil
}

const (
	Struct2_DD Struct2Field = "DD"
)
