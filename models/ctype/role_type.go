package ctype

import (
	"encoding/json"
)

type Role int

const (
	PermissionAdmin        Role = iota + 1
	PermissionUser         Role = iota + 1
	PermissionVisitor      Role = iota + 1
	PermissionDisabledUser Role = iota + 1
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisabledUser:
		str = "被禁用的用户"
	default:
		str = "其他"
	}
	return str
}
