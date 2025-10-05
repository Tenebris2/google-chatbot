package chat

import (
	"context"

	gc "google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

// chat client

type Client struct {
	Service *gc.Service
}

func NewClient(ctx context.Context, credsPath string) *Client {
	svc, err := gc.NewService(ctx, option.WithCredentialsFile(credsPath))

	if err != nil {
		panic(err)
	}
	return &Client{
		Service: svc,
	}
}
