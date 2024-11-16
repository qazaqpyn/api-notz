package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/pkg/repository"
)

type AuthorizationService struct {
	repo *repository.Repository
}

type TokenClaim struct {
	jwt.StandardClaims
	UserId pgtype.UUID `json:"userId"`
}

const (
	tokenTTL   = 12 * time.Hour
	salt       = "gloryToKazakhstan"
	signingKey = "gloryToKazakhstan"
)

func UUIDToString(myUUID pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", myUUID.Bytes[0:4], myUUID.Bytes[4:6], myUUID.Bytes[6:8], myUUID.Bytes[8:10], myUUID.Bytes[10:16])
}

func StringToUUID(myString string) pgtype.UUID {
	var myUUID pgtype.UUID
	fmt.Sscanf(myString, "%x-%x-%x-%x-%x", myUUID.Bytes[0:4], myUUID.Bytes[4:6], myUUID.Bytes[6:8], myUUID.Bytes[8:10], myUUID.Bytes[10:16])

	return myUUID
}

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
	user.Id = pgtype.UUID{}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthorizationService) Login(ctx context.Context, body model.LoginRequest) (string, string, error) {
	password := generatePasswordHash(body.Password)

	user, err := s.repo.GetUser(ctx, body.Username, password)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := s.GenerateTokenPair(ctx, user.Id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthorizationService) GenerateTokenPair(ctx context.Context, userId pgtype.UUID) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   UUIDToString(userId),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
	})

	accessToken, err := t.SignedString([]byte(salt))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := t.SignedString([]byte(signingKey))
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
			return nil, fmt.Errorf("invalid token")
		}

		return []byte(salt), nil
	})

	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
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
