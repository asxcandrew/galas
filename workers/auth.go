package workers

import (
	"context"
	"errors"
	"time"

	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
	"github.com/asxcandrew/galas/user"
	jwt "github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

const TokenTTL = time.Hour * 750

type AuthWorker interface {
	Login(string, string) (*model.User, error)
	Register(*model.User, string) error
	GenerateToken(*model.User) (string, time.Time, error)
	NewJWTParser(e endpoint.Endpoint) endpoint.Endpoint
}

type authWorker struct {
	storage     *storage.Storage
	userService *user.UserService
	secret      []byte
}

type CustomClaims struct {
	UserID   int    `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

// NewAuthWorker creates an auth handler with necessary dependencies.
func NewAuthWorker(us *user.UserService, s *storage.Storage, secret string) AuthWorker {
	return &authWorker{
		userService: us,
		storage:     s,
		secret:      []byte(secret),
	}
}

func GetClaims(c context.Context) (*CustomClaims, error) {
	if claims, ok := c.Value(gokitjwt.JWTClaimsContextKey).(*CustomClaims); !ok {
		return nil, errors.New("claims err")
	} else {
		return claims, nil
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

func (w *authWorker) NewJWTParser(e endpoint.Endpoint) endpoint.Endpoint {
	keys := func(_ *jwt.Token) (interface{}, error) {
		return w.secret, nil
	}
	claims := func() jwt.Claims {
		return &CustomClaims{}
	}
	return gokitjwt.NewParser(keys, jwt.SigningMethodHS256, claims)(e)
}

// GenerateToken generates token with expiration date for a given user
func (w *authWorker) GenerateToken(user *model.User) (string, time.Time, error) {
	expiration := time.Now().Add(TokenTTL)

	claims := CustomClaims{
		user.ID,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(w.secret)

	return tokenString, expiration, err
}
