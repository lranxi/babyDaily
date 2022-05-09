package babies

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database/babies"
)

type babyService struct {
	babyRepository babies.BabyRepository
}

type BabyService interface {
	List(userId int) (model.Babies, error)
	Create(baby *model.Baby) (*model.Baby, error)
	Update(baby *model.Baby) (*model.Baby, error)
	Delete(baby *model.Baby) error
	GenerateDefault(userId int) (*model.Baby, error)
}

// NewBabyService 创建baby业务层接口
func NewBabyService(babyRepository babies.BabyRepository) BabyService {
	return &babyService{babyRepository: babyRepository}
}

func (b *babyService) List(userId int) (model.Babies, error) {
	return b.babyRepository.List(userId)
}

func (b *babyService) Create(baby *model.Baby) (*model.Baby, error) {
	return b.babyRepository.Create(baby)
}

func (b *babyService) Update(baby *model.Baby) (*model.Baby, error) {
	return b.babyRepository.Update(baby)
}

func (b *babyService) Delete(baby *model.Baby) error {
	return b.babyRepository.Delete(baby)
}

func (b *babyService) GenerateDefault(userId int) (*model.Baby, error) {
	return b.babyRepository.GenerateDefault(userId)
}
