package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/wizard-xie/realworld/internal/biz"
)

type realworldRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealworldRepo .
func NewRealworldRepo(data *Data, logger log.Logger) biz.RealwroldRepo {
	return &realworldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realworldRepo) Save(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	return g, nil
}

func (r *realworldRepo) Update(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	return g, nil
}

func (r *realworldRepo) FindByID(context.Context, int64) (*biz.Realworld, error) {
	return nil, nil
}

func (r *realworldRepo) ListByHello(context.Context, string) ([]*biz.Realworld, error) {
	return nil, nil
}

func (r *realworldRepo) ListAll(context.Context) ([]*biz.Realworld, error) {
	return nil, nil
}
