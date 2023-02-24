package util

import (
	"userservice/internal/dto/resp"

	userPb "github.com/ideaspaper/social-media-proto/user"
)

func RespUserDtoToPb(userDto *resp.UserDto) *userPb.UserResp {
	return &userPb.UserResp{
		Id:        int64(userDto.ID),
		Email:     userDto.Email,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		CreatedAt: userDto.CreatedAt,
		UpdatedAt: userDto.UpdatedAt,
		DeletedAt: userDto.DeletedAt,
	}
}
