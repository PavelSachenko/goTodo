package user

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"newExp/internal/app"
	"newExp/internal/model"
	"newExp/internal/repository/user"
	"time"
)

type Service struct {
	auth user.Authorization
	user user.User
}

type Claims struct {
	jwt.StandardClaims
	UserId uint64 `json:"userId"`
}

var (
	salt      = "asdad.3r92AJSDu123as"
	signedKey = "asd#knasdASD21.AS32o#kjnsdfasdjasd"
)

func NewService(auth user.Authorization, user user.User) *Service {
	return &Service{
		auth: auth,
		user: user,
	}
}

func (s *Service) GetUser(id uint64) (*model.User, error) {
	user, err := s.user.Get(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *Service) CreateUser(user *model.User) (string, error) {
	user.Password = s.hashPassword(user.Password)
	id, err := s.auth.CreateUser(user)
	if err != nil {
		return "", err
	}
	return s.generateSignedToken(id)
}

func (s *Service) SignIn(username, password string) (string, error) {
	user, err := s.auth.SignIn(username, s.hashPassword(password))
	if err != nil {
		return "", app.ErrWrongCredentials
	}

	return s.generateSignedToken(user.ID)
}

func (s *Service) ParseToke(bearerToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(bearerToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *Service) generateSignedToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	return token.SignedString([]byte(signedKey))
}

func (s *Service) hashPassword(password string) string {
	sha := sha1.New()
	sha.Write([]byte(password))
	sha.Write([]byte(salt))

	return fmt.Sprintf("%x", sha.Sum(nil))
}
