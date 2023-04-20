package repositories

import (
	"github.com/laertkokona/crud-test/models"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepo interface {
	FindAll(pagination models.Pagination) ([]models.User, error)
	FindByID(int) (models.User, error)
	FindByUsername(string) (models.User, error)
	Save(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(models.User) error
	DeleteById(int) (models.User, error)
}

// NewUserRepo returns a new instance of userRepo
func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		DB: db,
	}
}

// FindAll returns all users
func (u userRepo) FindAll(pagination models.Pagination) ([]models.User, error) {
	// If pagination is not set, return all users
	// If pagination is set, return users based on pagination
	var users []models.User
	if pagination.Limit == 0 || pagination.Page == 0 {
		if err := u.DB.Find(&users).Error; err != nil {
			return nil, err
		}
		return users, u.DB.Find(&users).Error
	}

	if err := u.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, u.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&users).Error
}

// FindByID returns a user by id
func (u userRepo) FindByID(id int) (models.User, error) {
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, u.DB.First(&user, id).Error
}

// FindByUsername returns a user by username
func (u userRepo) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := u.DB.First(&user, "username=?", username).Error; err != nil {
		return user, err
	}
	return user, u.DB.First(&user, "username=?", username).Error
}

// Save saves a user
func (u userRepo) Save(user models.User) (models.User, error) {
	return user, u.DB.Create(&user).Error
}

// Update updates a user
func (u userRepo) Update(user models.User) (models.User, error) {
	return user, u.DB.Save(&user).Error
	//if err := u.DB.First(&user, user.ID).Error; err != nil {
	//	return user, err
	//}
	//return user, u.DB.Model(&user).Updates(&user).Error
}

// Delete deletes a user
func (u userRepo) Delete(user models.User) error {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return err
	}
	return u.DB.Delete(&user).Error
}

// DeleteById deletes a user by id
func (u userRepo) DeleteById(id int) (models.User, error) {
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, u.DB.Delete(&user).Error
}
