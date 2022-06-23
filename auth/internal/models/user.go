package models

import (
	pb "app/auth/pkg/proto/gen"
	"crypto/sha1"
	"fmt"
)

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

func UserFromProto(user *pb.User) *User {
	return &User{
		ID:    user.GetId(),
		Name:  user.GetName(),
		Login: user.GetLogin(),
	}
}

func PasswordSHA1(pass, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
