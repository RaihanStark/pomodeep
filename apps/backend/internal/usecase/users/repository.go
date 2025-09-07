package users

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/raihanstark/pomodeep/db"
)

// Repository defines the interface for user data operations
type Repository interface {
	Create(ctx context.Context, email, password string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int32) (*User, error)
}

// repository implements Repository interface
type repository struct {
	queries *db.Queries
}

// NewRepository creates a new user repository
func NewRepository(queries *db.Queries) Repository {
	return &repository{
		queries: queries,
	}
}

// Create creates a new user
func (r *repository) Create(ctx context.Context, email, password string) (*User, error) {
	dbUser, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	user := FromDBUser(dbUser)
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	user := FromDBUser(dbUser)
	return &user, nil
}

// GetByID retrieves a user by ID
func (r *repository) GetByID(ctx context.Context, id int32) (*User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	user := FromDBUser(dbUser)
	return &user, nil
}
