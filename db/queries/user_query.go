package queries

import (
	"unittest/models"

	"gorm.io/gorm"
)

type IUserQueries interface {
	GetUsers() (*[]models.User, error)
	GetUser(id int64) (*models.User, error)
	CreateUser(user models.CreateUserData) (*models.User, error)
}

type UserQueries struct {
	db *gorm.DB
}

func NewUserQueries(db *gorm.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (q *UserQueries) GetUsers() (*[]models.User, error) {
	var users []models.User
	result := q.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (q *UserQueries) GetUser(id int64) (*models.User, error) {
	var user models.User
	result := q.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (q *UserQueries) CreateUser(user models.CreateUserData) (*models.User, error) {
	createdUser := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}

	result := q.db.Create(&createdUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &createdUser, nil
}
