package models

import (
	pb "app/auth/pkg/proto/gen"
)

type User struct {
	ID       uint64
	Name     string
	Login    string
	Password string
	Email    string
}

func (u *User) Proto() *pb.User {

	return &pb.User{
		Id:    u.ID,
		Name:  u.Name,
		Login: u.Login,
		Email: u.Email,
	}
}

func UserFromProto(user *pb.User) *User {
	return &User{
		ID:    user.GetId(),
		Name:  user.GetName(),
		Login: user.GetLogin(),
		Email: user.GetEmail(),
	}
}
