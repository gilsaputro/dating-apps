package user

import (
	"gilsaputro/dating-apps/internal/store/user"
	"gilsaputro/dating-apps/pkg/hash"
	"gilsaputro/dating-apps/pkg/token"
	"strings"
)

// UserServiceMethod is list method for User Service
type UserServiceMethod interface {
	AddUser(AddUserServiceRequest) error
	DeleteUser(DeleteUserServiceRequest) error
	UpdateUser(UpdateUserServiceRequest) (UserServiceInfo, error)
	GetUserByID(GetByIDServiceRequest) (UserServiceInfo, error)
}

// UserService is list dependencies for user service
type UserService struct {
	store user.UserStoreMethod
	token token.TokenMethod
	hash  hash.HashMethod
}

// NewUserService is func to generate UserServiceMethod interface
func NewUserService(store user.UserStoreMethod, token token.TokenMethod, hash hash.HashMethod) UserServiceMethod {
	return &UserService{
		hash:  hash,
		token: token,
		store: store,
	}
}

// AddUser is service layer func to validate and creating user to database if the user is not exists and the token is verified
func (u *UserService) AddUser(request AddUserServiceRequest) error {
	_, err := u.token.ValidateToken(request.TokenRequest)
	if err != nil {
		return ErrUnauthorized
	}

	userInfo, err := u.store.GetUserInfoByUsername(request.Username)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}

	if userInfo.UserId > 0 {
		return ErrUserNameAlreadyExists
	}

	hashPassword, err := u.hash.HashValue(request.Password)
	if err != nil {
		return err
	}

	return u.store.CreateUser(user.UserStoreInfo{
		Username: request.Username,
		Password: string(hashPassword),
		Fullname: request.Fullname,
		Email:    request.Email,
	})
}

// DeleteUser is service level func to validate and delete user info in database
func (u *UserService) DeleteUser(request DeleteUserServiceRequest) error {
	value, err := u.token.ValidateToken(request.TokenRequest)
	if err != nil {
		return ErrUnauthorized
	}

	if value.Username != request.Username {
		return ErrCannotDeleteOtherUser
	}

	userInfo, err := u.store.GetUserInfoByUsername(request.Username)
	if err != nil || userInfo.UserId <= 0 {
		if (err == nil && userInfo.UserId <= 0) || strings.Contains(err.Error(), "not found") {
			return ErrUserNameNotExists
		}
		return err
	}

	if !u.hash.CompareValue(userInfo.Password, request.Password) {
		return ErrPasswordIsIncorrect
	}

	return u.store.DeleteUser(userInfo.UserId)
}

// UpdateUser is service level func to validate and update user info in database
func (u *UserService) UpdateUser(request UpdateUserServiceRequest) (UserServiceInfo, error) {
	value, err := u.token.ValidateToken(request.TokenRequest)
	if err != nil {
		return UserServiceInfo{}, ErrUnauthorized
	}

	if value.Username != request.Username {
		return UserServiceInfo{}, ErrCannotUpdateOtherUser
	}

	userInfo, err := u.store.GetUserInfoByUsername(request.Username)
	if err != nil || userInfo.UserId <= 0 {
		if (err == nil && userInfo.UserId <= 0) || strings.Contains(err.Error(), "not found") {
			return UserServiceInfo{}, ErrUserNameNotExists
		}
		return UserServiceInfo{}, err
	}

	if len(request.Password) > 0 {
		hashPassword, err := u.hash.HashValue(request.Password)
		if err != nil {
			return UserServiceInfo{}, err
		}

		userInfo.Password = string(hashPassword)
	}

	if len(request.Email) > 0 {
		userInfo.Email = request.Email
	}

	if len(request.Fullname) > 0 {
		userInfo.Fullname = request.Fullname
	}

	err = u.store.UpdateUser(userInfo)
	if err != nil {
		return UserServiceInfo{}, err
	}

	return UserServiceInfo{
		UserId:      userInfo.UserId,
		Username:    userInfo.Username,
		Fullname:    userInfo.Fullname,
		Email:       userInfo.Email,
		CreatedDate: userInfo.CreatedDate,
	}, nil
}

// GetUserByID is service level func to validate and get all user based id
func (u *UserService) GetUserByID(request GetByIDServiceRequest) (UserServiceInfo, error) {
	value, err := u.token.ValidateToken(request.TokenRequest)
	if err != nil {
		return UserServiceInfo{}, ErrUnauthorized
	}

	if value.UserID != int(request.UserId) {
		return UserServiceInfo{}, ErrCannotGetOtherUser
	}

	userInfo, err := u.store.GetUserInfoByID(int(request.UserId))
	if err != nil || userInfo.UserId <= 0 {
		if (err == nil && userInfo.UserId <= 0) || strings.Contains(err.Error(), "not found") {
			return UserServiceInfo{}, ErrUserNameNotExists
		}
		return UserServiceInfo{}, err
	}

	return UserServiceInfo{
		UserId:      userInfo.UserId,
		Username:    userInfo.Username,
		Fullname:    userInfo.Fullname,
		Email:       userInfo.Email,
		CreatedDate: userInfo.CreatedDate,
	}, nil
}
