package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IUser interface {
	FindAll(db *gorm.DB) (*[]models.User, error)
	FindbyId(db *gorm.DB, uid uint) (*models.User, error)
	FindbyUsername(db *gorm.DB, username string) (*models.User, error)
	Save(db *gorm.DB, user *models.User) (*models.User, error)
	Update(db *gorm.DB, user *models.User, uid uint) (*models.User, error)
	Delete(db *gorm.DB, uid uint) (int64, error)
}

type UserRepo struct{}

func NewUserRepo() IUser {
	return &UserRepo{}
}

func NewUser(user models.User) *models.User {
	return &models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Active:    user.Active,
		IsAdmin:   user.IsAdmin,
	}
}

func (cr *UserRepo) FindAll(db *gorm.DB) (*[]models.User, error) {
	user := &[]models.User{}
	err := db.Model(&models.User{}).Find(&user).Error
	if err != nil {
		log.Error().AnErr("course save error ::", err)
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) FindbyId(db *gorm.DB, uid uint) (*models.User, error) {
	user := &models.User{}
	err := db.Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) FindbyUsername(db *gorm.DB, username string) (*models.User, error) {
	user := &models.User{}
	err := db.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) Save(db *gorm.DB, user *models.User) (*models.User, error) {
	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Error().AnErr("course save error ::", err)
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) Update(db *gorm.DB, user *models.User, uid uint) (*models.User, error) {
	data := &models.User{}
	err := db.Model(&models.User{}).Where("id = ?", uid).Updates(user).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *UserRepo) Delete(db *gorm.DB, uid uint) (int64, error) {
	result := db.Model(&models.User{}).Where("id = ?", uid).Delete(&models.User{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
