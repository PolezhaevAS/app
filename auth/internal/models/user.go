package models

import pb "app/auth/pkg/proto/gen"

type User struct {
	ID       uint64
	Name     string
	Login    string
	Password string
}

func (u *User) Proto() *pb.User {
	return &pb.User{
		Id:    u.ID,
		Name:  u.Name,
		Login: u.Login,
	}
}
