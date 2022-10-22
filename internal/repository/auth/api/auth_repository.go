package authRepository

import "context"

type AuthRepository interface {
	ValidateCredentials(ctx context.Context, email string, password string) error
}
