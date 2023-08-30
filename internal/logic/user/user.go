package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"user/internal/dao"
	"user/internal/model"
	"user/internal/service"
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

type (
	sUser struct{}
)

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	one, err := dao.User.Ctx(ctx).One("nickname = ?", in.Nickname)
	if err != nil {
		return err
	}
	println(one.Map())
	if len(one) == 0 {
		_, err = dao.User.Ctx(ctx).Insert(&in)
		if err != nil {
			return err
		}
	} else {
		println(one)
		return gerror.New("用户已经存在")
	}
	return
}
