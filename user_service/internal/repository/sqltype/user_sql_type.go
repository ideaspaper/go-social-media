package sqltype

import (
	"database/sql"
	"userservice/internal/model"
)

type User struct {
	ID        sql.NullInt64
	Email     sql.NullString
	Password  sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

func (u User) ToModel() *model.User {
	if !u.ID.Valid {
		return nil
	}
	result := &model.User{
		ID:        int(u.ID.Int64),
		Email:     u.Email.String,
		Password:  u.Password.String,
		FirstName: u.FirstName.String,
		LastName:  u.LastName.String,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
	if u.DeletedAt.Valid {
		result.DeletedAt = &u.DeletedAt.Time
	}
	return result
}
