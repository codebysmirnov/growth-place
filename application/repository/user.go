package repository

import (
	"github.com/jackc/pgconn"
	"gorm.io/gorm"

	"growth-place/application/domain"
)

// UserRepo presents domain.User db instance
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo return new UserRepo instance
func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

// Create write new record on users table
func (r UserRepo) Create(user domain.User) error {
	err := r.db.Create(&user).Error
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.ConstraintName {
		case "users__login_uniq_idx":
			return ErrUserWithPassedLoginIsExists
		case "users__email_uniq_idx":
			return ErrUserWithPassedEmailIsExists
		case "users__phone_uniq_idx":
			return ErrUserWithPassedPhoneIsExists
		}
	}
	return nil
}
