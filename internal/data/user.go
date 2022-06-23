package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/wizard-xie/realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	return g, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.Realworld, error) {
	return nil, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.Realworld, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.Realworld, error) {
	return nil, nil
}
