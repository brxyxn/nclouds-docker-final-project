package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Config cache
var ctx = context.Background()
var rdb *redis.Client

// Config handlers
type Handlers struct {
	client *redis.Client
}

func NewHandlers(client *redis.Client) *Handlers {
	return &Handlers{client}
}
