package repository

import "github.com/scyna/core/examples/chat/account/proto"

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (u *User) FromDTO(user *proto.User) {
	u.ID = user.Id
	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password
}

func (u *User) ToDTO() *proto.User {
	return &proto.User{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
