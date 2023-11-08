package user_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToUserResponse(user entities.User) UserResponse {
	return UserResponse{
		Id:       float64(user.ID),
		UserName: user.UserName,
	}
}
