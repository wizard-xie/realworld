package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/wizard-xie/realworld/api/realworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Realworld is a Realworld model.
type Realworld struct {
	Hello string
}

// RealwroldRepo is a Greater repo.
type RealwroldRepo interface {
	Save(context.Context, *Realworld) (*Realworld, error)
	Update(context.Context, *Realworld) (*Realworld, error)
	FindByID(context.Context, int64) (*Realworld, error)
	ListByHello(context.Context, string) ([]*Realworld, error)
	ListAll(context.Context) ([]*Realworld, error)
}

// RealworldUsecase is a Greeter usecase.
type RealworldUsecase struct {
	repo RealwroldRepo
	log  *log.Helper
}

// NewRealworldUsecase new a Greeter usecase.
func NewRealworldUsecase(repo RealwroldRepo, logger log.Logger) *RealworldUsecase {
	return &RealworldUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *RealworldUsecase) CreateGreeter(ctx context.Context, g *Realworld) (*Realworld, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
