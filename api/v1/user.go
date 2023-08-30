package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta    `path:"/setup" tags:"setup" method:"post" summary:"注册"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type RegisterRes struct {
}
type LoginReq struct {
	g.Meta `path:"/login" tags:"login" method:"post" summary:"登录"`
}
type LoginRes struct {
}
