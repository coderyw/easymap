package model

import (
	fmt "fmt"
	easy_facade "github.com/coderyw/easymap/gen/facade"
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
	if val, ok = m["float"]; ok {
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.Float = float32(pv)
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
		case int:
			v.Code = int(acval)
		case int8:
			v.Code = int(acval)
		case int16:
			v.Code = int(acval)
		case int32:
			v.Code = int(acval)
		case uint8:
			v.Code = int(acval)
		case uint16:
			v.Code = int(acval)
		case int64:
			v.Code = int(acval)
		case uint:
			v.Code = int(acval)
		case uint32:
			v.Code = int(acval)
		case uint64:
			v.Code = int(acval)
		case string:
			if len(acval) == 0 {
				v.Code = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.Code = int(pvv)
			}
		}
	}
	if val, ok = m["message"]; ok {
		switch acval := val.(type) {
		case string:
			v.Message = acval
		case int16:
			v.Message = strconv.FormatInt(int64(acval), 10)
		case int64:
			v.Message = strconv.FormatInt(int64(acval), 10)
		case uint8:
			v.Message = strconv.FormatUint(uint64(acval), 10)
		case uint32:
			v.Message = strconv.FormatUint(uint64(acval), 10)
		case float64:
			v.Message = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		case int:
			v.Message = strconv.FormatInt(int64(acval), 10)
		case int8:
			v.Message = strconv.FormatInt(int64(acval), 10)
		case int32:
			v.Message = strconv.FormatInt(int64(acval), 10)
		case uint:
			v.Message = strconv.FormatUint(uint64(acval), 10)
		case uint16:
			v.Message = strconv.FormatUint(uint64(acval), 10)
		case uint64:
			v.Message = strconv.FormatUint(uint64(acval), 10)
		case float32:
			v.Message = strconv.FormatFloat(float64(acval), 'f', -1, 64)
		}
	}
	if val, ok = m["b"]; ok {
		switch acval := val.(type) {
		case int8:
			v.B = int64(acval)
		case int32:
			v.B = int64(acval)
		case uint8:
			v.B = int64(acval)
		case uint16:
			v.B = int64(acval)
		case uint32:
			v.B = int64(acval)
		case int:
			v.B = int64(acval)
		case int16:
			v.B = int64(acval)
		case int64:
			v.B = int64(acval)
		case uint:
			v.B = int64(acval)
		case uint64:
			v.B = int64(acval)
		case string:
			if len(acval) == 0 {
				v.B = 0
			} else {
				pvv, err := strconv.ParseInt(acval, 10, 64)
				if err != nil {
					return err
				}
				v.B = int64(pvv)
			}
		}
	}
	if val, ok = m["float"]; ok {
		switch acval := val.(type) {
		case uint8:
			v.Float = float32(acval)
		case float32:
			v.Float = float32(acval)
		case int:
			v.Float = float32(acval)
		case int16:
			v.Float = float32(acval)
		case int32:
			v.Float = float32(acval)
		case uint16:
			v.Float = float32(acval)
		case uint32:
			v.Float = float32(acval)
		case uint64:
			v.Float = float32(acval)
		case string:
			if len(acval) == 0 {
				v.Float = 0
			} else {
				pvv, err := strconv.ParseFloat(acval, 10)
				if err != nil {
					return err
				}
				v.Float = float32(pvv)
			}
		case float64:
			v.Float = float32(acval)
		case int8:
			v.Float = float32(acval)
		case int64:
			v.Float = float32(acval)
		case uint:
			v.Float = float32(acval)
		}
	}
	if val, ok = m["DD"]; ok {
		switch acval := val.(type) {
		case []int:
			v.DD = acval
		case []int8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []int32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []int64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []uint8:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []uint64:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []int16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []uint:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []uint16:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []uint32:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmpArr[i] = int(k)
			}
			v.DD = tmpArr
		case []string:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				tmp, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return err
				}
				tmpArr[i] = int(tmp)
			}
			v.DD = tmpArr
		case []interface{}:
			tmpArr := make([]int, len(acval))
			for i, k := range acval {
				switch trueK := k.(type) {
				case int:
					tmpArr[i] = trueK
				case int8:
					tmpArr[i] = int(trueK)
				case int16:
					tmpArr[i] = int(trueK)
				case int32:
					tmpArr[i] = int(trueK)
				case int64:
					tmpArr[i] = int(trueK)
				case uint8:
					tmpArr[i] = int(trueK)
				case uint16:
					tmpArr[i] = int(trueK)
				case uint32:
					tmpArr[i] = int(trueK)
				case uint64:
					tmpArr[i] = int(trueK)
				case string:
					tmp, err := strconv.ParseInt(trueK, 10, 64)
					if err != nil {
						return err
					}
					tmpArr[i] = int(tmp)
				case struct{}:

				}
			}
			v.DD = tmpArr
		}
	}
	if val, ok = m["dec"]; ok {
		switch acval := val.(type) {
		case uint16:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint64:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int16:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int32:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int64:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint8:
			v.Dec = decimal.NewFromInt(int64(acval))
		case uint32:
			v.Dec = decimal.NewFromInt(int64(acval))
		case int8:
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
		case int:
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
		case int8:
			var decc decimal.Decimal
			decc = decimal.NewFromInt(int64(acval))
			v.PtrDec = &decc
		case int16:
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
		case float64:
			var decc decimal.Decimal
			decc = decimal.NewFromFloat(float64(acval))
			v.PtrDec = &decc
		case float32:
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
	if val, ok = m["abc_arr"]; ok {
		switch acval := val.(type) {
		case []map[string]string:
			v.AbcArr = make([]Resp360Arr, len(acval))
			for i, k := range acval {
				var kjj interface{} = Resp360Arr{}
				if b, ok := kjj.(easy_facade.EasyMapString); ok {
					if err := b.UnMarshalMap(k); err != nil {
						return err
					}
				}
				v.AbcArr[i] = kjj.(Resp360Arr)
			}
		case []map[string]interface{}:
			v.AbcArr = make([]Resp360Arr, len(acval))
			for i, k := range acval {
				var kjj interface{} = Resp360Arr{}
				if b, ok := kjj.(easy_facade.EasyMapInter); ok {
					if err := b.UnMarshalMapInterface(k); err != nil {
						return err
					}
				}
				v.AbcArr[i] = kjj.(Resp360Arr)
			}
		case []interface{}:
			v.AbcArr = make([]Resp360Arr, len(acval))
			for i, k := range acval {
				switch kt := k.(type) {
				case map[string]interface{}:
					var kjj interface{} = Resp360Arr{}
					if b, ok := kjj.(easy_facade.EasyMapInter); ok {
						if err := b.UnMarshalMapInterface(kt); err != nil {
							return err
						}
					}
					v.AbcArr[i] = kjj.(Resp360Arr)
				case map[string]string:
					var kjj interface{} = Resp360Arr{}
					if b, ok := kjj.(easy_facade.EasyMapString); ok {
						if err := b.UnMarshalMap(kt); err != nil {
							return err
						}
					}
					v.AbcArr[i] = kjj.(Resp360Arr)
				}

			}
		}
	}

	return nil
}

func (v *Resp360) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["code"] = v.Code
	m["message"] = v.Message
	m["b"] = v.B
	m["float"] = v.Float
	m["dec"] = v.Dec.String()
	m["ptrDec"] = v.PtrDec.String()
	return m, nil
}

func (v *Resp360) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["code"] = fmt.Sprint(v.Code)
	m["message"] = fmt.Sprint(v.Message)
	m["b"] = fmt.Sprint(v.B)
	m["float"] = fmt.Sprint(v.Float)
	m["dec"] = v.Dec.String()
	m["ptrDec"] = v.PtrDec.String()
	return m, nil
}

func (v *Resp360) CheckField(m map[string]string) error {
	if _, ok := m["code"]; !ok {
		return fmt.Errorf("field code is miss")
	}
	if _, ok := m["message"]; !ok {
		return fmt.Errorf("field message is miss")
	}
	if _, ok := m["b"]; !ok {
		return fmt.Errorf("field b is miss")
	}
	if _, ok := m["float"]; !ok {
		return fmt.Errorf("field float is miss")
	}
	if _, ok := m["DD"]; !ok {
		return fmt.Errorf("field DD is miss")
	}
	if _, ok := m["dec"]; !ok {
		return fmt.Errorf("field dec is miss")
	}
	if _, ok := m["ptrDec"]; !ok {
		return fmt.Errorf("field ptrDec is miss")
	}
	if _, ok := m["data"]; !ok {
		return fmt.Errorf("field data is miss")
	}
	if _, ok := m["else"]; !ok {
		return fmt.Errorf("field else is miss")
	}
	if _, ok := m["abc_arr"]; !ok {
		return fmt.Errorf("field abc_arr is miss")
	}
	return nil
}

type Resp360Field string

func (v Resp360Field) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	Resp360_Code    Resp360Field = "code"
	Resp360_Message Resp360Field = "message"
	Resp360_B       Resp360Field = "b"
	Resp360_Float   Resp360Field = "float"
	Resp360_DD      Resp360Field = "DD"
	Resp360_Dec     Resp360Field = "dec"
	Resp360_PtrDec  Resp360Field = "ptrDec"
	Resp360_Data    Resp360Field = "data"
	Resp360_Else    Resp360Field = "else"
	Resp360_AbcArr  Resp360Field = "abc_arr"
)
