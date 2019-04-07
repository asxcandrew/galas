package workers

import (
	"context"
	"errors"
	"time"

	"github.com/asxcandrew/galas/faults"

	"github.com/asxcandrew/galas/storage/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
)

const TokenTTL = time.Hour * 750

type AuthWorker interface {
	GenerateToken(*model.User) (string, time.Time, error)
	NewJWTParser(e endpoint.Endpoint) endpoint.Endpoint
}

type authWorker struct {
	secret []byte
}

type CustomClaims struct {
	UserID   int    `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

// NewAuthWorker creates an auth handler with necessary dependencies.
func NewAuthWorker(secret string) AuthWorker {
	return &authWorker{
		secret: []byte(secret),
	}
}

func (w *authWorker) NewJWTParser(e endpoint.Endpoint) endpoint.Endpoint {
	keys := func(_ *jwt.Token) (interface{}, error) {
		return w.secret, nil
	}
	claims := func() jwt.Claims {
		return &CustomClaims{}
	}
	return NewParser(keys, jwt.SigningMethodHS256, claims)(e)
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

func GetClaims(c context.Context) (*CustomClaims, error) {
	if claims, ok := c.Value(JWTClaimsContextKey).(*CustomClaims); !ok {
		return nil, faults.BuildRichError(faults.UnauthorisedError, errors.New("Can`t extract claims"))
	} else {
		return claims, nil
	}
}
