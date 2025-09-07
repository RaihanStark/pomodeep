package users

import (
	"context"
	"errors"

	"github.com/raihanstark/pomodeep/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// Service defines the interface for user business logic
type Service interface {
	SignUp(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error)
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

func (s *service) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate the token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		User:  *user,
		Token: token,
	}, nil
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
