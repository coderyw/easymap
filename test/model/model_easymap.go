package model

import (
	easy_facade "github.com/coderyw/easymap/facade"
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *WaitGroupWrapper) UnMarshalMap(m map[string]string) error {
	return nil
}
func (v *WaitGroupWrapper) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["WaitGroup"]; ok {
		var i interface{} = v.WaitGroup
		if m1, ok := val.(map[string]interface{}); ok {
			if b, ok := i.(easy_facade.EasyMapInter); ok {
				if err := b.UnMarshalMapInterface(m1); err != nil {
					return err
				}
			}
		} else if m2, ok := val.(map[string]string); ok {
			if b, ok := i.(easy_facade.EasyMapString); ok {
				if err := b.UnMarshalMap(m2); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (v *WaitGroupWrapper) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

type WaitGroupWrapperField string

func (v WaitGroupWrapperField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	WaitGroupWrapper_WaitGroup WaitGroupWrapperField = "WaitGroup"
)

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
			v.B = string(pv)
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
			v.B = string(val.(string))
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

	return nil
}

func (v *TestStruct) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["a"] = v.A
	m["b"] = v.B
	m["c"] = v.C
	m["d"] = v.D
	return m, nil
}

type TestStructField string

func (v TestStructField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	TestStruct_A TestStructField = "a"
	TestStruct_B TestStructField = "b"
	TestStruct_C TestStructField = "c"
	TestStruct_D TestStructField = "d"
)

func (v *Struct2) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["dd"]; ok {
		{
			pv := val
			v.DD = string(pv)
		}
	}

	return nil
}
func (v *Struct2) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["dd"]; ok {
		switch val.(type) {
		case string:
			v.DD = string(val.(string))
		}
	}

	return nil
}

func (v *Struct2) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["dd"] = v.DD
	return m, nil
}

type Struct2Field string

func (v Struct2Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	Struct2_DD Struct2Field = "dd"
)

func (v *Resp360) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["code"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Code = int(pv)
		}
	}
	if val, ok = m["message"]; ok {
		{
			pv := val
			v.Message = string(pv)
		}
	}

	return nil
}
func (v *Resp360) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["code"]; ok {
		switch val.(type) {
		case int:
			v.Code = int(val.(int))
		case int8:
			v.Code = int(val.(int8))
		case int16:
			v.Code = int(val.(int16))
		case int32:
			v.Code = int(val.(int32))
		case int64:
			v.Code = int(val.(int64))
		case uint:
			v.Code = int(val.(uint))
		case uint8:
			v.Code = int(val.(uint8))
		case uint16:
			v.Code = int(val.(uint16))
		case uint32:
			v.Code = int(val.(uint32))
		case uint64:
			v.Code = int(val.(uint64))
		}
	}
	if val, ok = m["message"]; ok {
		switch val.(type) {
		case string:
			v.Message = string(val.(string))
		}
	}

	return nil
}

func (v *Resp360) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["code"] = v.Code
	m["message"] = v.Message
	return m, nil
}

type Resp360Field string

func (v Resp360Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	Resp360_Code    Resp360Field = "code"
	Resp360_Message Resp360Field = "message"
	Resp360_Data    Resp360Field = "data"
	Resp360_Else    Resp360Field = "else"
)
