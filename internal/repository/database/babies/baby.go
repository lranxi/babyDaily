package babies

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database"
	"gorm.io/gorm"
	"time"
)

var createBabyFields = []string{"user_id", "name", "height", "weight", "age", "age_unit", "photo", "gender"}

type babyRepository struct {
	db *gorm.DB
}

type BabyRepository interface {
	Migrate() error
	Create(baby *model.Baby) (*model.Baby, error)
	Update(baby *model.Baby) (*model.Baby, error)
	Delete(baby *model.Baby) error
	List(userId int) (model.Babies, error)
	GenerateDefault(userId int) (*model.Baby, error)
}

// NewBabyRepository 创建baby持久层
func NewBabyRepository() BabyRepository {
	return &babyRepository{
		db: database.Get(),
	}
}

func (b *babyRepository) Migrate() error {
	return b.db.AutoMigrate(&model.Baby{})
}

// List 查询用户的baby信息
func (b *babyRepository) List(userId int) (model.Babies, error) {
	babies := make(model.Babies, 0)
	if err := b.db.Order("created_at").Where("user_id = ?", userId).Find(&babies).Error; err != nil {
		return nil, err
	}
	return babies, nil
}

// Create 创建
func (b *babyRepository) Create(baby *model.Baby) (*model.Baby, error) {
	if err := b.db.Create(baby).Error; err != nil {
		return nil, err
	}

	return baby, nil
}

// Update 更新信息
func (b *babyRepository) Update(baby *model.Baby) (*model.Baby, error) {
	if err := b.db.Model(&model.Baby{}).Where("id = ?", baby.ID).Updates(baby).Error; err != nil {
		return nil, err
	}

	return baby, nil
}

// Delete 删除数据
func (b *babyRepository) Delete(baby *model.Baby) error {
	err := b.db.Where("user_id = ?", baby.ID).Delete(baby).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *babyRepository) GenerateDefault(userId int) (*model.Baby, error) {
	baby := &model.Baby{
		UserId:    userId,
		Name:      "",
		Height:    0,
		Weight:    0,
		Brith:     "",
		Gender:    0,
		Photo:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return b.Create(baby)
}
