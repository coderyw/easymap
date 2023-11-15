package model
import(
	unsafe "unsafe"
)

func str2Bytes_Fs(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *Fs) UnMarshalMap(m map[string]string) error{
	var (
		ok bool
		val string
	)
	if val,ok=m["DDD"];ok{
		{pv := val
			v.DDD = string(pv)
		}
	}


	return nil
}
func (v *Fs) UnMarshalMapInterface(m map[string]interface{}) error{
	var (
		ok bool
		val interface{}
	)
	if val,ok=m["DDD"];ok{
		switch val.(type){
		case string:
			v.DDD = string(val.(string))
		}
	}


	return nil
}


func (v *Fs) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["DDD"] = v.DDD
	return m, nil
}


type FsField string
func (v FsField) MarshalBinary() (data []byte, err error) {
return str2Bytes_Fs(string(v)), nil
}
const(

	Fs_DDD FsField = "DDD"
)
