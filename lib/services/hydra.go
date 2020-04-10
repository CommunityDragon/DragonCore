package services

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

type HydraService struct {
	baseURL string
}

// ORY Hydra Service
func Hydra() *HydraService {
	return &HydraService{"http://hydra:4445"}
}

// Name returns the service name
func (h *HydraService) Name() string {
	return "hydra"
}

// GetLoginEndpoint returns the hydra login request endpoint
func (h *HydraService) GetLoginEndpoint(challenge string) string {
	return fmt.Sprintf("%s/oauth2/auth/requests/login/%s", h.baseURL, challenge)
}

// GetLoginChallenge gets the login challenge from the params
func (h *HydraService) GetLoginChallenge(c echo.Context) (string, error) {
	challenge := c.QueryParam("login_challenge")
	if challenge == "" {
		return "", errors.New("login challenge not found")
	}
	return challenge, nil
}