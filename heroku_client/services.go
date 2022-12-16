package heroku_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/selefra/selefra-provider-heroku/constants"
	"net/http"
	"os"
	"strings"

	heroku "github.com/heroku/heroku-go/v5"
)

func Connect(ctx context.Context, herokuConfig *Config) (*heroku.Service, error) {
	email := os.Getenv(constants.HEROKUEMAIL)
	apiKey := os.Getenv(constants.HEROKUAPIKEY)
	if herokuConfig.Email != constants.Constants_0 {
		email = herokuConfig.Email
	}
	if herokuConfig.APIKey != constants.Constants_1 {
		apiKey = herokuConfig.APIKey
	}
	if email == constants.Constants_2 {
		return nil, errors.New(constants.Emailmustbeconfigured)
	}
	if apiKey == constants.Constants_3 {
		return nil, errors.New(constants.Apikeymustbeconfigured)
	}
	return heroku.NewService(&http.Client{
		Transport: &heroku.Transport{
			Username:  email,
			Password:  apiKey,
			UserAgent: fmt.Sprintf(constants.Sselefraproviderheroku, heroku.DefaultUserAgent),
			Transport: heroku.RoundTripWithRetryBackoff{},
		},
	}), nil
}

func IsCancelled(ctx context.Context) bool {
	err := ctx.Err()
	return err != nil && (errors.Is(err, context.Canceled) || strings.Contains(err.Error(), constants.Contextcanceled))
}
