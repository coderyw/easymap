package model

import (
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_TestStruct(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *TestStruct) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["A"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.A = int(pv)
		}
	}
	if val, ok = m["B"]; ok {
		{
			pv := val
			pvv := string(pv)
			v.B = &pvv
		}
	}
	if val, ok = m["C"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.C = float64(pv)
		}
	}
	if val, ok = m["D"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			v.D = bool(pv)
		}
	}
	if val, ok = m["E"]; ok {
		if pv, err := strconv.ParseUint(val, 10, 64); err != nil {
			return err
		} else {
			v.E = uint8(pv)
		}
	}

	return nil
}
func (v *TestStruct) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["A"]; ok {
		switch val.(type) {
		case int:
			v.A = int(val.(int))
		case int8:
			v.A = int(val.(int8))
		case int16:
			v.A = int(val.(int16))
		case int32:
			v.A = int(val.(int32))
		case int64:
			v.A = int(val.(int64))
		case uint:
			v.A = int(val.(uint))
		case uint8:
			v.A = int(val.(uint8))
		case uint16:
			v.A = int(val.(uint16))
		case uint32:
			v.A = int(val.(uint32))
		case uint64:
			v.A = int(val.(uint64))
		}
	}
	if val, ok = m["B"]; ok {
		switch val.(type) {
		case string:
			pvv := string(val.(string))
			v.B = &pvv
		}
	}
	if val, ok = m["C"]; ok {
		switch val.(type) {
		case float32:
			v.C = float64(val.(float32))
		case float64:
			v.C = float64(val.(float64))
		}
	}
	if val, ok = m["D"]; ok {
		switch val.(type) {
		case bool:
			v.D = bool(val.(bool))
		}
	}
	if val, ok = m["E"]; ok {
		switch val.(type) {
		case int:
			v.E = uint8(val.(int))
		case int8:
			v.E = uint8(val.(int8))
		case int16:
			v.E = uint8(val.(int16))
		case int32:
			v.E = uint8(val.(int32))
		case int64:
			v.E = uint8(val.(int64))
		case uint:
			v.E = uint8(val.(uint))
		case uint8:
			v.E = uint8(val.(uint8))
		case uint16:
			v.E = uint8(val.(uint16))
		case uint32:
			v.E = uint8(val.(uint32))
		case uint64:
			v.E = uint8(val.(uint64))
		}
	}

	return nil
}

func (v *TestStruct) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["A"] = v.A
	if v.B != nil {
		m["B"] = *v.B
	}
	m["C"] = v.C
	m["D"] = v.D
	m["E"] = v.E
	return m, nil
}

type TestStructField string

func (v TestStructField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_TestStruct(string(v)), nil
}

const (
	TestStruct_A TestStructField = "A"
	TestStruct_B TestStructField = "B"
	TestStruct_C TestStructField = "C"
	TestStruct_D TestStructField = "D"
	TestStruct_E TestStructField = "E"
)
