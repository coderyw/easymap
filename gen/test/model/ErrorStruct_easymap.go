package model

import (
	fmt "fmt"
	decimal "github.com/shopspring/decimal"
	strconv "strconv"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
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
	if val, ok = m["b"]; ok {
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.B = int64(pv)
		}
	}
	if val, ok = m["dec"]; ok {
		var err error
		v.Dec, err = decimal.NewFromString(val)
		if err != nil {
			return err
		}
	}
	if val, ok = m["ptrDec"]; ok {
		var err error
		var decc decimal.Decimal
		decc, err = decimal.NewFromString(val)
		if err != nil {
			return err
		}
		v.PtrDec = &decc
	}

	return nil
}
func (v *Resp360) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["code"]; ok {
		switch acval := val.(type) {
		case int64:
			v.Code = int(acval)
		case uint:
			v.Code = int(acval)
		case uint8:
			v.Code = int(acval)
		case uint16:
			v.Code = int(acval)
		case uint64:
			v.Code = int(acval)
		case int:
			v.Code = int(acval)
		case int32:
			v.Code = int(acval)
		case uint32:
			v.Code = int(acval)
		case int8:
			v.Code = int(acval)
		case int16:
			v.Code = int(acval)
		}
	}
	if val, ok = m["message"]; ok {
		switch acval := val.(type) {
		case string:
			v.Message = string(acval)
		}
	}
	if val, ok = m["b"]; ok {
		switch acval := val.(type) {
		case int:
			v.B = int64(acval)
		case int8:
			v.B = int64(acval)
		case int16:
			v.B = int64(acval)
		case int32:
			v.B = int64(acval)
		case uint:
			v.B = int64(acval)
		case uint8:
			v.B = int64(acval)
		case uint32:
			v.B = int64(acval)
		case int64:
			v.B = int64(acval)
		case uint16:
			v.B = int64(acval)
		case uint64:
			v.B = int64(acval)
		}
	}
	if val, ok = m["dec"]; ok {
		switch acval := val.(type) {
		case int8:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int32:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int64:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint16:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int16:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint8:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint32:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint64:
			v.Dec = decimal.NewFromInt(int64(acval))
		case float32:
			v.Dec = decimal.NewFromFloat(float64(acval))
		case float64:
			v.Dec = decimal.NewFromFloat(float64(acval))
		case string:
			var err error
			v.Dec, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
		}
	}
	if val, ok = m["ptrDec"]; ok {
		switch acval := val.(type) {
		case int8:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case int64:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case uint8:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case uint64:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case int:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case int16:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case int32:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case uint:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case uint16:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case uint32:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case float32:
			var decc decimal.Decimal
			decc = decimal.NewFromFloat(float64(acval))
			v.PtrDec = &decc
		case float64:
			var decc decimal.Decimal
			decc = decimal.NewFromFloat(float64(acval))
			v.PtrDec = &decc
		case string:
			var err error
			var decc decimal.Decimal
			decc, err = decimal.NewFromString(acval)
			if err != nil {
				return err
			}
			v.PtrDec = &decc
		}
	}

	return nil
}

func (v *Resp360) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["code"] = v.Code
	m["message"] = v.Message
	m["b"] = v.B
	m["dec"] = v.Dec.String()
	m["ptrDec"] = v.PtrDec.String()
	return m, nil
}

func (v *Resp360) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["code"] = fmt.Sprint(v.Code)
	m["message"] = fmt.Sprint(v.Message)
	m["b"] = fmt.Sprint(v.B)
	m["dec"] = v.Dec.String()
	m["ptrDec"] = v.PtrDec.String()
	return m, nil
}

type Resp360Field string

func (v Resp360Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	Resp360_Code    Resp360Field = "code"
	Resp360_Message Resp360Field = "message"
	Resp360_B       Resp360Field = "b"
	Resp360_DD      Resp360Field = "DD"
	Resp360_Dec     Resp360Field = "dec"
	Resp360_PtrDec  Resp360Field = "ptrDec"
	Resp360_Data    Resp360Field = "data"
	Resp360_Else    Resp360Field = "else"
	Resp360_AbcArr  Resp360Field = "abc_arr"
)
