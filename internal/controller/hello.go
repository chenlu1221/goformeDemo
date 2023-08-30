package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"user/api/v1"
	"user/internal/consts"
	"user/internal/model"
	"user/internal/service"
)

var (
	UserApi = cUser{}
)

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code": consts.UserDuplicationCode,
			"data": err,
		})
		return nil, err
	}
	return
}
