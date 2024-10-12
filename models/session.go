package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/boaltl/lenslocked/rand"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID     int
	UserID int
	// here Token is only set when creating a new session. When looking up it is empty
	// since we only store TokenHash in our db and it can't be reversed to raw token
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
	// if it is less than MinBytesPerToken, then second is used
	BytesPerToken int
}

// for creating session
func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		// skipping ID so that it's set by DB itself
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	// TODO: store the session in DB
	return &session, nil
}

// for looking up a User with given token in our db
func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
