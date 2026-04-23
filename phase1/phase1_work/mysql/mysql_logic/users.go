package mysql_logic

import (
	"context"
	"errors"
	"phase1/phase1_work/mysql"
	"phase1/phase1_work/mysql/mysql_model"

	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user *mysql_model.Users) (int64, error) {
	if err := mysql.DB.Create(user).Error; err != nil {
		return 0, err
	}

	if user.Id <= 0 {
		return 0, errors.New("create user failed, invalid user id")
	}

	return user.Id, nil
}

func GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*mysql_model.Users, error) {
	var user mysql_model.Users
	if err := mysql.DB.Where("username = ? AND password = ? AND deleted_at = 0", username, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if user.Id <= 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUserById(ctx context.Context, userId int64) (*mysql_model.Users, error) {
	var user mysql_model.Users
	if err := mysql.DB.Where("id = ? AND deleted_at = 0", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if user.Id <= 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUsersByIds(ctx context.Context, userIds []int64) (map[int64]*mysql_model.Users, error) {
	var users []mysql_model.Users
	if err := mysql.DB.Where("id IN ? AND deleted_at = 0", userIds).Find(&users).Error; err != nil {
		return nil, err
	}

	userMap := make(map[int64]*mysql_model.Users)
	for i := range users {
		userMap[users[i].Id] = &users[i]
	}

	return userMap, nil
}
