package models

import "database/sql"

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
}

// for creating session
func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: create a session token and and implement this func
	return nil, nil
}

// for looking up a User with given token in our db
func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
