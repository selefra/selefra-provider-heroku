package heroku_client

import (
	"context"
	heroku "github.com/heroku/heroku-go/v5"
)

func Connect(_ context.Context, config *Config) (*heroku.Service, error) {
	heroku.DefaultTransport.BearerToken = config.Token
	client := heroku.DefaultClient
	client.Transport = Paginator{transport: client.Transport}
	return heroku.NewService(client), nil
}
