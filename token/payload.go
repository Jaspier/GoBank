package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different error types returned by the VerifyToken function
var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
	ErrPayloadEmpty = errors.New("payload is empty")
)

// Payload contains the paylaod data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	if payload != nil {
		return jwt.NewNumericDate(payload.ExpiredAt), nil
	}
	return nil, ErrPayloadEmpty
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	if payload != nil {
		return jwt.NewNumericDate(payload.IssuedAt), nil
	}
	return nil, ErrPayloadEmpty
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	if payload != nil {
		return jwt.NewNumericDate(payload.IssuedAt), nil
	}
	return nil, ErrPayloadEmpty
}
func (payload *Payload) GetIssuer() (string, error) {
	if payload != nil {
		return payload.Username, nil
	}
	return "", ErrPayloadEmpty
}
func (payload *Payload) GetSubject() (string, error) {
	return "JWT", nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return []string{"gobank_user"}, nil
}
