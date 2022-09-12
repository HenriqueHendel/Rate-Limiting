package client

import (
	"errors"

	tokenBucket "github.com/henriquehendel/rateLimiting/rateLimiter"
)

type Client struct {
	Name      string `json:"name"`
	MaxTokens int64  `json:"maxTokens"`
	FillRate  int64  `json:"fillRate"`
}

var clientBucketMap = make(map[string]*tokenBucket.TokenBucket)

type Rule struct {
	MaxTokens int64
	Rate      int64
}

func SetNewClient(c Client) {
	clientBucketMap[c.Name] = tokenBucket.NewTokenBucket(c.FillRate, c.MaxTokens)
}

func GetBucket(identifier string) (*tokenBucket.TokenBucket, error) {
	clientBucket := clientBucketMap[identifier]

	if clientBucket == nil {
		return nil, errors.New("Client not found")
	}

	return clientBucket, nil
}
