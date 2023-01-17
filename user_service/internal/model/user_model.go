package model

import (
	"time"
	"userservice/internal/dto/resp"
)

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u User) ToDto() *resp.UserDto {
	result := &resp.UserDto{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
	if u.DeletedAt != nil {
		deletedAtString := u.DeletedAt.String()
		result.DeletedAt = &deletedAtString
	}
	return result
}
