package entities

import (
	"time"
)

type ApiKey interface {
	GetUUID() string
	GetSecret() string
	SetUUID(uuid string)
	SetSecret(secret string)
	GetSlug() string
	SetSlug(slug string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
}

type apiKey struct {
	uuid      string
	secret    string
	slug      string
	createdAt time.Time
}

func NewApiKey(uuid string, secret string, slug string) ApiKey {
	return &apiKey{
		uuid:   uuid,
		secret: secret,
		slug:   slug,
	}
}

func NewApiKeyWithCreatedAt(uuid string, secret string, slug string, createdAt time.Time) ApiKey {
	return &apiKey{
		uuid:      uuid,
		secret:    secret,
		slug:      slug,
		createdAt: createdAt,
	}
}

func (r *apiKey) GetUUID() string {
	return r.uuid
}

func (r *apiKey) GetSecret() string {
	return r.secret
}

func (r *apiKey) SetUUID(uuid string) {
	r.uuid = uuid
}

func (r *apiKey) SetSecret(secret string) {
	r.secret = secret
}

func (r *apiKey) GetSlug() string {
	return r.slug
}

func (r *apiKey) SetSlug(slug string) {
	r.slug = slug
}

func (r *apiKey) GetCreatedAt() time.Time {
	return r.createdAt
}

func (r *apiKey) SetCreatedAt(createdAt time.Time) {
	r.createdAt = createdAt
}
