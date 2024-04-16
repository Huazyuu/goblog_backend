package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = iota + 1
	SignGitee SignStatus = iota + 1
	SignEmail SignStatus = iota + 1
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "gitee"
	case SignEmail:
		str = "Email"
	default:
		str = "其他"
	}
	return str
}
