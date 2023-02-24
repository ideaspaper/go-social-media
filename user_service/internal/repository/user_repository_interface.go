package repository

import (
	"context"
	"userservice/internal/dto/req"
	"userservice/internal/model"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id int) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, userDto *req.UserDto) (*model.User, error)
	DeleteByID(ctx context.Context, id int) (*model.User, error)
	DeletePermanentlyByID(ctx context.Context, id int) (*model.User, error)
}
