package repository

import "gorm.io/gorm"

type userSqlliteRepository struct {
	db *gorm.DB
}

func NewUserSqlliteRepository(db *gorm.DB) UserRepository {
	return &userSqlliteRepository{db: db}
}

func (r userSqlliteRepository) Create(request User) (*User, error) {
	user := User{
		Username: request.Username,
		Email:    request.Email,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
