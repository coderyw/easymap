package model
import(
	unsafe "unsafe"
	strconv "strconv"
)

func str2Bytes_ConfigureAliCdnDomainReq(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *ConfigureAliCdnDomainReq) UnMarshalMap(m map[string]string) error{
	var (
		ok bool
		val string
	)
	if val,ok=m["Domain"];ok{
		{pv := val
			v.Domain = string(pv)
		}
	}
	if val,ok=m["Sub"];ok{
		{pv := val
			v.Sub = string(pv)
		}
	}
	if val,ok=m["TenantId"];ok{
		if pv, err := strconv.ParseInt(val, 10, 64); err != nil {
			return err
		} else {
			v.TenantId = int64(pv)
		}
	}
	if val,ok=m["Token"];ok{
		{pv := val
			v.Token = string(pv)
		}
	}


	return nil
}


func (v *ConfigureAliCdnDomainReq) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["Domain"] = v.Domain
	m["Sub"] = v.Sub
	m["TenantId"] = v.TenantId
	m["Token"] = v.Token
	return m, nil
}


type ConfigureAliCdnDomainReqField string
func (v ConfigureAliCdnDomainReqField) MarshalBinary() (data []byte, err error) {
return str2Bytes_ConfigureAliCdnDomainReq(string(v)), nil
}
const(

	ConfigureAliCdnDomainReq_Domain ConfigureAliCdnDomainReqField = "Domain"
	ConfigureAliCdnDomainReq_Sub ConfigureAliCdnDomainReqField = "Sub"
	ConfigureAliCdnDomainReq_TenantId ConfigureAliCdnDomainReqField = "TenantId"
	ConfigureAliCdnDomainReq_Token ConfigureAliCdnDomainReqField = "Token"
	ConfigureAliCdnDomainReq_XXX_sizecache ConfigureAliCdnDomainReqField = "XXX_sizecache"
)
