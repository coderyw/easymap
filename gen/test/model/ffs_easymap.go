package model

import (
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_fs(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *fs) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["bs"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Bs = int(pv)
		}
	}

	return nil
}
func (v *fs) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["bs"]; ok {
		switch val.(type) {
		case int:
			v.Bs = int(val.(int))
		case int8:
			v.Bs = int(val.(int8))
		case int16:
			v.Bs = int(val.(int16))
		case int32:
			v.Bs = int(val.(int32))
		case int64:
			v.Bs = int(val.(int64))
		case uint:
			v.Bs = int(val.(uint))
		case uint8:
			v.Bs = int(val.(uint8))
		case uint16:
			v.Bs = int(val.(uint16))
		case uint32:
			v.Bs = int(val.(uint32))
		case uint64:
			v.Bs = int(val.(uint64))
		}
	}

	return nil
}

func (v *fs) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["bs"] = v.Bs
	return m, nil
}

type fsField string

func (v fsField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_fs(string(v)), nil
}

const (
	fs_Bs fsField = "bs"
)
