package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/pkg/repository"
)

type AuthorizationService struct {
	repo *repository.Repository
}

type TokenClaim struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

const (
	tokenTTL   = 12 * time.Hour
	salt       = "gloryToKazakhstan"
	signingKey = "gloryToKazakhstan"
)

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthService(repo *repository.Repository) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(ctx context.Context, user model.User) error {
	user.Password = generatePasswordHash(user.Password)
	user.Id = uuid.New().String()

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthorizationService) Login(ctx context.Context, body model.LoginRequest) (string, string, error) {
	password := generatePasswordHash(body.Password)

	user, err := s.repo.GetUser(ctx, body.Email, password)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := s.GenerateTokenPair(ctx, user.Id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthorizationService) GenerateTokenPair(ctx context.Context, userId string) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
	})

	accessToken, err := t.SignedString([]byte(salt))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {

		return "", "", err
	}

	if err := s.repo.Token.Create(ctx, model.RefreshSession{
		UserId:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(tokenTTL),
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthorizationService) ParseToken(ctx context.Context, token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(salt), nil
	})

	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return id, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *AuthorizationService) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	session, err := s.repo.Get(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", errors.New("refresh token expired")
	}

	return s.GenerateTokenPair(ctx, session.UserId)
}
