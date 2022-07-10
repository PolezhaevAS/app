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
	Access   map[string][]string
}

func (u *User) Proto() *pb.User {

	access := make(map[string]*pb.Methods)

	for service, methods := range u.Access {
		access[service] = &pb.Methods{Name: methods}
	}

	return &pb.User{
		Id:     u.ID,
		Name:   u.Name,
		Login:  u.Login,
		Email:  u.Email,
		Access: access,
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
