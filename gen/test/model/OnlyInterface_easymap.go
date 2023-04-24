package model

import (
	unsafe "unsafe"
)

func str2Bytes_OnlyInterface(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *OnlyInterface) UnMarshalMap(m map[string]string) error {
	return nil
}

func (v *OnlyInterface) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

type OnlyInterfaceField string

func (v OnlyInterfaceField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_OnlyInterface(string(v)), nil
}

const ()
