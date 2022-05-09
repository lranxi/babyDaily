package users

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database"
	"gorm.io/gorm"
)

// 创建用户需要的字段
var (
	userCreateField = []string{"nickname", "open_id", "union_id", "avatar"}
)

type userRepository struct {
	db *gorm.DB
}

// UserRepository 用户信息 持久层接口
type UserRepository interface {
	GetUserByOpenId(openId string) (*model.User, error)
	GetUserByUnionId(unionId string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Migrate() error
}

func New() UserRepository {
	return &userRepository{db: database.Get()}
}

func (u *userRepository) GetUserByOpenId(openId string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Where("open_id = ?", openId).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetUserByUnionId(unionId string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Where("union_id = ?", unionId).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Update(user *model.User) (*model.User, error) {
	if err := u.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
