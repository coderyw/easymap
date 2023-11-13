package model
import(
	unsafe "unsafe"
	strconv "strconv"
)

func str2Bytes_Resp360(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *Resp360) UnMarshalMap(m map[string]string) error{
	var (
		ok bool
		val string
	)
	if val,ok=m["Code"];ok{
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.Code = int(pv)
		}
	}
	if val,ok=m["Message"];ok{
		{pv := val
			v.Message = string(pv)
		}
	}


	return nil
}


func (v *Resp360) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["Code"] = v.Code
	m["Message"] = v.Message
	return m, nil
}


type Resp360Field string
func (v Resp360Field) MarshalBinary() (data []byte, err error) {
return str2Bytes_Resp360(string(v)), nil
}
const(

	Resp360_Code Resp360Field = "Code"
	Resp360_Message Resp360Field = "Message"
)
