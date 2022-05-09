package records

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database"
	"gorm.io/gorm"
)

var createRecordFields = []string{}

type recordRepository struct {
	db *gorm.DB
}

type RecordRepository interface {
	Migrate() error
	Create(record *model.Record) (*model.Record, error)
	Update(record *model.Record) (*model.Record, error)
	List(userId, babyId int, day string) (*model.Records, error)
	Delete(record *model.Record) error
	Details(userId int, babyId int, rType int, curTime string) (*model.Records, error)
}

// NewRecordRepository 创建record持久化接口
func NewRecordRepository() RecordRepository {
	return &recordRepository{
		db: database.Get(),
	}
}

func (r *recordRepository) Migrate() error {
	return r.db.AutoMigrate(&model.Record{})
}

func (r *recordRepository) List(userId, babyId int, day string) (*model.Records, error) {
	start := day + " 00:00:00"
	end := day + " 23:59:59"
	records := new(model.Records)
	if err := r.db.Raw("SELECT user_id,baby_id,type,sub_type,SUM(CASE quantity WHEN 0 THEN 1 ELSE quantity END) AS quantity FROM t_record WHERE user_id = ? AND baby_id = ? AND `created_at` >= ? AND `created_at` <= ? GROUP BY type,sub_type", userId, babyId, start, end).Scan(records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *recordRepository) Create(record *model.Record) (*model.Record, error) {
	if err := r.db.Create(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *recordRepository) Update(record *model.Record) (*model.Record, error) {
	if err := r.db.Model(&model.Record{}).Where("id = ?", record.ID).Updates(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *recordRepository) Delete(record *model.Record) error {
	return r.db.Where("user_id = ?", record.UserId).Delete(record).Error
}

func (r *recordRepository) Details(userId int, babyId int, rType int, curTime string) (*model.Records, error) {
	start := curTime + " 00:00:00"
	end := curTime + " 23:59:59"
	records := new(model.Records)
	if err := r.db.Order("created_at").Where("user_id = ? and baby_id = ? and `type` = ? and `created_at` >= ? and `created_at` <= ?", userId, babyId, rType, start, end).Find(records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
