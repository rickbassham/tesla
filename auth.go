package tesla

import (
	"fmt"
	"net/http"
)

// Authenticate starts the initial authentication process via an OAuth 2.0 Password Grant with the
// same credentials used for tesla.com and the mobile apps.
//
// The current client ID and secret are available at https://pastebin.com/pS7Z6yyP.
//
// We will get back an access_token which is treated as an OAuth 2.0 Bearer Token. This token is
// passed along in an Authorization header with all future requests:
func (c *Conn) Authenticate(email, password string) error {
	type request struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Email        string `json:"email"`
		Password     string `json:"password"`
	}

	reqBody := request{
		GrantType:    "password",
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Email:        email,
		Password:     password,
	}

	type response struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		CreatedAt    int    `json:"created_at"`
	}

	var respBody response

	err := c.doRequest(http.MethodPost, "/oauth/token?grant_type=password", &reqBody, &respBody)
	if err != nil {
		return err
	}

	c.accessToken = respBody.AccessToken
	c.refreshToken = respBody.RefreshToken

	return nil
}

// UpdateRefreshToken will do an OAuth 2.0 Refresh Token Grant and obtain a new access token. Note:
// This will invalidate the previous access token.
func (c *Conn) UpdateRefreshToken() error {
	if c.refreshToken == "" {
		return fmt.Errorf("%w", ErrMissingRefreshToken)
	}

	type request struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
	}

	reqBody := request{
		GrantType:    "refresh_token",
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		RefreshToken: c.refreshToken,
	}

	type response struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		CreatedAt    int    `json:"created_at"`
	}

	var respBody response

	err := c.doRequest(http.MethodPost, "/oauth/token?grant_type=refresh_token", &reqBody, &respBody)
	if err != nil {
		return err
	}

	c.accessToken = respBody.AccessToken
	c.refreshToken = respBody.RefreshToken

	return nil
}
