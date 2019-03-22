package workers

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
	"github.com/asxcandrew/galas/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthWorker interface {
	Login(string, string) (*model.User, error)
	Register(*model.User, string) error
}

type authWorker struct {
	storage     *storage.Storage
	userService *user.UserService
}

// NewAuthWorker creates an auth handler with necessary dependencies.
func NewAuthWorker(us *user.UserService, s *storage.Storage) AuthWorker {
	return &authWorker{
		userService: us,
		storage:     s,
	}
}

func (w *authWorker) Login(email, password string) (user *model.User, err error) {
	user, err = w.storage.User.GetByEmail(email)
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

func (w *authWorker) Register(user *model.User, password string) (err error) {
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

	err = w.storage.User.Create(user)

	return err
}
