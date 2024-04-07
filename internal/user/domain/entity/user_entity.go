package entity

import "github.com/google/uuid"

type UserEntity struct {
	id    string
	name  string
	email string
}

func NewUserEntity(name, email string) (*UserEntity, error) {
	return &UserEntity{
		id:    uuid.New().String(),
		name:  name,
		email: email,
	}, nil
}

func (u *UserEntity) GetID() string {
	return u.id
}

func (u *UserEntity) GetName() string {
	return u.name
}

func (u *UserEntity) GetEmail() string {
	return u.email
}
