package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type (
	UserRole int16
)

const (
	_ UserRole = iota
	UserRoleGlobalAdmin
	UserRoleAdmin
	UserRoleModerator
	UserRoleUser
)

type (
	User struct {
		Id        int64      `json:"id"`
		Username  string     `json:"username"`
		Email     string     `json:"email"`
		Password  string     `json:"-"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		Role      UserRole   `json:"role"`
	}

	CreateUserRequest struct {
		Id       int64    `json:"id"`
		Username string   `json:"username"`
		Email    string   `json:"email"`
		Password string   `json:"-"`
		Role     UserRole `json:"role"`
	}
	CreateUserResponse struct {
		User *User `json:"data"`
	}

	GetUserRequest struct {
		UserId int64 `json:"-"`
	}
	GetUserResponse struct {
		User *User `json:"data"`
	}
	GetUsersRequest struct {
		UserId int64 `json:"-"`
	}
	GetUsersResponse struct {
		Users []User `json:"data"`
	}
	UpdateUserRequest struct {
		UserId   int64     `json:"user_id"`
		Username *string   `json:"username"`
		Email    *string   `json:"email"`
		Password *string   `json:"password"`
		Role     *UserRole `json:"role"`
	}
	UpdateUserResponse struct {
		User *User `json:"data"`
	}

	DeleteUserRequest struct {
		UserId int64 `json:"user_id"`
	}
	DeleteUserResponse struct{}
)

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
