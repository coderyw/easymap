package model
import(
	strconv "strconv"
)
func (v *TestStruct) UnMarshalMap(m map[string]string) error{
	var (
		ok bool
		val string
	)
	if val,ok=m["a"];ok{
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.A = int(pv)
		}
	}
	if val,ok=m["b"];ok{
		v.B = &val
	}
	if val,ok=m["c"];ok{
		if pv, err := strconv.ParseFloat(val, 10); err != nil {
			return err
		} else {
			v.C = pv
		}
	}
	if val,ok=m["d"];ok{
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
	m["a"] = v.A
	if v.B != nil {
		m["b"] = *v.B
	}
	m["c"] = v.C
	m["d"] = v.D
	return m, nil
}


const(

	TestStruct_A = "a"
	TestStruct_B = "b"
	TestStruct_C = "c"
	TestStruct_D = "d"
)
func (v *Struct2) UnMarshalMap(m map[string]string) error{
	var (
		ok bool
		val string
	)
	if val,ok=m["dd"];ok{
		v.DD = val
	}

	return nil
}


func (v *Struct2) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["dd"] = v.DD
	return m, nil
}


const(

	Struct2_DD = "dd"
)
func (v *Resp360) UnMarshalMap(m map[string]string) error{
	return nil
}
