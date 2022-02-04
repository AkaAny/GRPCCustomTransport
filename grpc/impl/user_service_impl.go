package impl

import (
	"context"
	"grpc_custom_transport/grpc/def"
)

type UserServiceImpl struct {
	def.UnimplementedUserServiceServer
}

func (u *UserServiceImpl) GetUserInfo(ctx context.Context, in *def.GetUserInfoRequest) (*def.GetUserInfoResponse, error) {
	return &def.GetUserInfoResponse{
		UserName: "user1",
		Location: "hangzhou",
	}, nil
}
