package util

import (
	"userservice/internal/dto/req"
	"userservice/internal/dto/resp"

	userPb "github.com/ideaspaper/social-media-proto/user"
)

func RespUserDtoToPb(userDto *resp.UserDto) *userPb.UserResp {
	return &userPb.UserResp{
		Id:        int64(userDto.ID),
		Email:     userDto.Email,
		FirstName: userDto.Email,
		LastName:  userDto.LastName,
		CreatedAt: userDto.CreatedAt,
		UpdatedAt: userDto.UpdatedAt,
		DeletedAt: userDto.DeletedAt,
	}
}

func ReqUserPbToDto(in *userPb.Req) req.UserDto {
	return req.UserDto{
		Email:     in.GetUserReq().Email,
		Password:  in.GetUserReq().Password,
		FirstName: in.GetUserReq().FirstName,
		LastName:  in.GetUserReq().LastName,
	}
}
