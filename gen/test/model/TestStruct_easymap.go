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
		v.B = &val
	}
	if val, ok = m["C"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.C = pv
		}
	}
	if val, ok = m["D"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			v.D = pv
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
)
