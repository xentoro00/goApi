package auth

import (
	"context"
	// Using an alias is good practice to prevent name conflicts
	supaauth "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

type Service interface {
	Signup(ctx context.Context, req SignupRequest) error
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
}

// service is the implementation of the auth Service.
type service struct {
	// This field MUST be a concrete implementation of the client interface.
	supaClient supaauth.Client
}

// NewService creates a new authentication service.
func NewService(supaClient supaauth.Client) Service {
	return &service{supaClient: supaClient}
}

// Signup creates a new user in Supabase.
func (s *service) Signup(ctx context.Context, req SignupRequest) error {
	// This call is now correct because s.supaClient is a pointer.
	_, err := s.supaClient.Signup(types.SignupRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	return err
}

// Login authenticates a user and returns tokens.
func (s *service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	resp, err := s.supaClient.SignInWithEmailPassword(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		ExpiresIn:    resp.ExpiresIn,
	}, nil
}
