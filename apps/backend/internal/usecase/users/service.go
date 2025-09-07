package users

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service defines the interface for user business logic
type Service interface {
	SignUp(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
}

// service implements Service interface
type service struct {
	repo Repository
}

// NewService creates a new user service
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// SignUp creates a new user account
func (s *service) SignUp(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	user, err := s.repo.Create(ctx, req.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		User: *user,
	}, nil
}
