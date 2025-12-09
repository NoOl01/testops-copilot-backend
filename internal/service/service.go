package service

import (
	"context"
	"net"
	"net/http"
	"testops_copilot/internal/dto"
	"time"
)

type Service interface {
	Generate(testCase dto.Case, ctx context.Context) (*dto.GenerateResult, error)
}

type service struct {
	httpClient *http.Client
}

func NewService() Service {
	return &service{
		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 3 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ResponseHeaderTimeout: 30 * time.Second,
			},
			Timeout: 120 * time.Second,
		},
	}
}
