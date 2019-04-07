package user

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Get(string) (*model.User, error)
	Login(string, string) (*model.User, error)
	Register(*model.User, string) error
}

type userService struct {
	storage storage.Storage
}

// NewUserService creates an usre service with necessary dependencies.
func NewUserService(storage storage.Storage) UserService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) Get(username string) (user *model.User, err error) {
	user, err = s.storage.User.GetByUsername(username)

	return user, err
}

func (s *userService) Login(email, password string) (user *model.User, err error) {
	user, err = s.storage.User.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	byteHash := []byte(user.EncryptedPassword)
	bytePass := []byte(password)

	err = bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Register(user *model.User, password string) (err error) {
	bytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.EncryptedPassword = string(hashedBytes[:])
	err = user.Validate()

	if err != nil {
		return err
	}

	err = s.storage.User.Create(user)

	return err
}
