package models

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var (
	TokenBodyNumeric      = []rune("0123456789")
	TokenBodyAlphaNumeric = []rune("ABCDEFGHIJKLNMOPQRSTUVWXYZ0123456789")
)

type Token struct {
	Id            uuid.UUID    `json:"id" gorm:"primaryKey"`
	TokenableType string       `json:"tokenable_type"`
	TokenableId   uuid.UUID    `json:"tokenable_id" gorm:"index"`
	Type          string       `json:"type"`
	Body          string       `json:"body"` // the token content/body/string
	UsedAt        sql.NullTime `json:"used_at"`
	ExpiredAt     sql.NullTime `json:"expired_at"`
	CreatedAt     time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`

	// Relations

	Tokenable any `json:"tokenable" gorm:"-"`
}

// GenerateBody generates a new token body and assigns it to models.Token.Body
func (token *Token) GenerateBody(chars []rune, length int) {
	token.Body = ""
	for i := 0; i < length; i++ {
		char := chars[rand.Intn(len(chars))]
		token.Body = token.Body + string(char)
	}
}

// IsExpired returns bool that indicates token already expired or not
func (token Token) IsExpired() bool {
	return time.Now().After(token.ExpiredAt.Time) // if curremt time is after the expiration time it means token is expired
}

// IsUsed returns bool that indicates token already used or not used
func (token Token) IsUsed() bool {
	return token.UsedAt.Valid && (token.UsedAt.Time != time.Time{})
}
