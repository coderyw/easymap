package model

import (
	easy_facade "github.com/coderyw/easymap/gen/facade"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *OutModel) UnMarshalMap(m map[string]string) error {
	return nil
}
func (v *OutModel) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["unEasyMap"]; ok {
		var i interface{} = v.UnEasyMap
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
	if val, ok = m["ptrMap"]; ok {
		var i interface{} = v.PtrMap
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
	if val, ok = m["un"]; ok {
		var i interface{} = v.un
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

func (v *OutModel) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	return m, nil
}

func (v *OutModel) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	return m, nil
}

type OutModelField string

func (v OutModelField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	OutModel_UnEasyMap OutModelField = "unEasyMap"
	OutModel_PtrMap    OutModelField = "ptrMap"
	OutModel_un        OutModelField = "un"
)
