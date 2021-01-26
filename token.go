package omnimlib

import (
	"bytes"
	"fmt"
	"github.com/google/jsonapi"
)

type Token struct {
	ID    int    `jsonapi:"primary,tokens"`
	Value string `jsonapi:"attr,value"`
}

type TokenService struct {
	client *Client
}

func (s *TokenService) GetNew() (string, error) {
	payload, err := s.client.doRequest(
		"POST",
		"/tokens/",
		&ReqOptions{
			ContentType: "application/x-www-form-urlencoded",
			Args:        map[string]string{"username": s.client.authData.Username, "password": s.client.authData.Password},
		},
	)
	if err != nil {
		return "", err
	}
	token := new(Token)
	err = jsonapi.UnmarshalPayload(bytes.NewReader(payload), token)
	if err != nil {
		s.client.log.Errorf("Cant unmarshal payload: %s", err)
		return "", fmt.Errorf("Cant unmarshal payload")
	}

	return token.Value, nil
}
