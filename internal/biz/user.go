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

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *Realworld) (*Realworld, error)
	Update(context.Context, *Realworld) (*Realworld, error)
	FindByID(context.Context, int64) (*Realworld, error)
	ListByHello(context.Context, string) ([]*Realworld, error)
	ListAll(context.Context) ([]*Realworld, error)
}

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new user.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *Realworld) (*Realworld, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
