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
	if val, ok = m["a"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.A = int(pv)
		}
	}
	if val, ok = m["b"]; ok {
		{
			pv := val
			pvv := string(pv)
			v.B = &pvv
		}
	}
	if val, ok = m["c"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.C = float64(pv)
		}
	}
	if val, ok = m["d"]; ok {
		if pv, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			v.D = bool(pv)
		}
	}
	if val, ok = m["e"]; ok {
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
	if val, ok = m["a"]; ok {
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
	if val, ok = m["b"]; ok {
		switch val.(type) {
		case string:
			pvv := string(val.(string))
			v.B = &pvv
		}
	}
	if val, ok = m["c"]; ok {
		switch val.(type) {
		case float32:
			v.C = float64(val.(float32))
		case float64:
			v.C = float64(val.(float64))
		}
	}
	if val, ok = m["d"]; ok {
		switch val.(type) {
		case bool:
			v.D = bool(val.(bool))
		}
	}
	if val, ok = m["e"]; ok {
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
	if val, ok = m["f"]; ok {
		if m1, ok := val.(map[string]interface{}); ok {
			if err := v.F.UnMarshalMapInterface(m1); err != nil {
				return err
			}
		} else if m2, ok := val.(map[string]string); ok {
			if err := v.F.UnMarshalMap(m2); err != nil {
				return err
			}
		}
	}
	if val, ok = m["g"]; ok {
		if m1, ok := val.(map[string]interface{}); ok {
			if err := v.G.UnMarshalMapInterface(m1); err != nil {
				return err
			}
		} else if m2, ok := val.(map[string]string); ok {
			if err := v.G.UnMarshalMap(m2); err != nil {
				return err
			}
		}
	}
	if val, ok = m["hh"]; ok {
		if m1, ok := val.([]map[string]interface{}); ok {
			vv := make([]Fs, 0)
			for _, v1 := range m1 {
				ab := Fs{}
				if err := ab.UnMarshalMapInterface(v1); err != nil {
					return err
				}
				vv = append(vv, ab)
			}
			v.HH = vv
		} else if m1, ok := val.([]map[string]string); ok {
			vv := make([]Fs, 0)
			for _, v1 := range m1 {
				ab := Fs{}
				if err := ab.UnMarshalMap(v1); err != nil {
					return err
				}
				vv = append(vv, ab)
			}
			v.HH = vv
		}
	}
	if val, ok = m["hhs"]; ok {
		if m1, ok := val.([]map[string]interface{}); ok {
			vv := make([]*Fs, 0)
			for _, v1 := range m1 {
				ab := new(Fs)
				if err := ab.UnMarshalMapInterface(v1); err != nil {
					return err
				}
				vv = append(vv, ab)
			}
			v.HHS = vv
		} else if m1, ok := val.([]map[string]string); ok {
			vv := make([]*Fs, 0)
			for _, v1 := range m1 {
				ab := new(Fs)
				if err := ab.UnMarshalMap(v1); err != nil {
					return err
				}
				vv = append(vv, ab)
			}
			v.HHS = vv
		}
	}

	return nil
}

func (v *TestStruct) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["a"] = v.A
	if v.B != nil {
		m["b"] = *v.B
	}
	m["c"] = v.C
	m["d"] = v.D
	m["e"] = v.E
	return m, nil
}

type TestStructField string

func (v TestStructField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_TestStruct(string(v)), nil
}

const (
	TestStruct_A   TestStructField = "a"
	TestStruct_B   TestStructField = "b"
	TestStruct_C   TestStructField = "c"
	TestStruct_D   TestStructField = "d"
	TestStruct_E   TestStructField = "e"
	TestStruct_F   TestStructField = "f"
	TestStruct_G   TestStructField = "g"
	TestStruct_HH  TestStructField = "hh"
	TestStruct_HHS TestStructField = "hhs"
)
