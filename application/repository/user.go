package repository

import (
	"errors"

	"github.com/google/uuid"
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

// Read get user by id
func (r UserRepo) Read(id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := r.db.Model(user).Where("id=?", id).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

// ReadByMail get user by passed email
func (r UserRepo) ReadByMail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Model(user).Where("email=?", email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

// Update modify users record (non-zero fields)
func (r UserRepo) Update(user domain.User) error {
	err := r.db.Updates(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
