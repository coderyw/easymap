package model

import (
	model_path "github.com/coderyw/easymap/gen/test/model/model_path"
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *ErrorStruct) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["ErrorCode"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.ErrorCode = model_path.ErrorCODE(pv)
		}
	}
	if val, ok = m["Str"]; ok {
		{
			pv := val
			v.Str = string(pv)
		}
	}
	if val, ok = m["ErrorCodePtr"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			pvv := model_path.ErrorCODE(pv)
			v.ErrorCodePtr = &pvv
		}
	}
	if val, ok = m["A"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.A = int32(pv)
		}
	}
	if val, ok = m["Adfe"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Adfe = Aewe(pv)
		}
	}

	return nil
}

func (v *ErrorStruct) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["ErrorCode"] = v.ErrorCode
	m["Str"] = v.Str
	if v.ErrorCodePtr != nil {
		m["ErrorCodePtr"] = *v.ErrorCodePtr
	}
	m["A"] = v.A
	m["Adfe"] = v.Adfe
	return m, nil
}

type ErrorStructField string

func (v ErrorStructField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	ErrorStruct_ErrorCode    ErrorStructField = "ErrorCode"
	ErrorStruct_Str          ErrorStructField = "Str"
	ErrorStruct_ErrorCodePtr ErrorStructField = "ErrorCodePtr"
	ErrorStruct_A            ErrorStructField = "A"
	ErrorStruct_Adfe         ErrorStructField = "Adfe"
)
