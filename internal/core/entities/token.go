package entities

import "time"

type Token interface {
	UUID() string
	ApiKey() string
	CreatedAt() time.Time
}

type token struct {
	Uuid      string
	ApiKey    string
	createdAt time.Time
}

func NewToken(uuid string, apiKey string) *token {
	return &token{
		Uuid:      uuid,
		ApiKey:    apiKey,
		createdAt: time.Now(),
	}
}
