package records

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database/records"
	"baby-daily-api/pkg/timeutil"
	"errors"
)

type recordService struct {
	recordRepository records.RecordRepository
}

type RecordService interface {
	Create(record *model.Record) (*model.Record, error)
	Update(record *model.Record) (*model.Record, error)
	List(userId, babyId int, day string) ([]*model.RecordDetail, error)
	Delete(record *model.Record) error
	Details(userId int, babyId int, rType int, curTime string) (*model.Records, error)
}

func NewRecordService(recordRepository records.RecordRepository) RecordService {
	return &recordService{recordRepository: recordRepository}
}

// Create 创建记录
func (r *recordService) Create(record *model.Record) (*model.Record, error) {
	if err := validate(record); err != nil {
		return nil, err
	}
	return r.recordRepository.Create(record)
}

// Update 更新记录
func (r *recordService) Update(record *model.Record) (*model.Record, error) {
	if err := validate(record); err != nil {
		return nil, err
	}
	return r.recordRepository.Update(record)
}

// Delete 删除记录
func (r *recordService) Delete(record *model.Record) error {
	return r.recordRepository.Delete(record)
}

func (r *recordService) Details(userId int, babyId int, rType int, curTime string) (*model.Records, error) {
	return r.recordRepository.Details(userId, babyId, rType, curTime)
}

// 数据检查
func validate(record *model.Record) error {
	if record.UserId <= 0 || record.BabyId <= 0 {
		return errors.New("用户或宝贝id为空")
	}
	if record.Type <= 0 && record.SubType <= 0 {
		return errors.New("记录类型信息为空")
	}
	// 时间检查，开始时间必须小于结束时间
	if before, _ := timeutil.Before(record.Start, record.End); !before {
		return errors.New("开始时间必须小于结束时间")
	}
	return nil
}

// List 记录列表
func (r *recordService) List(userId, babyId int, day string) ([]*model.RecordDetail, error) {
	records, err := r.recordRepository.List(userId, babyId, day)
	if err != nil {
		return nil, err
	}
	remarkMap := make(map[int]*model.RecordDetail)
	for _, record := range *records {
		if detail, ok := remarkMap[record.Type]; ok {
			// 将当前记录放在record的subRecord中
			subDetails := detail.SubRecord
			sub := &model.SubRecordDetail{
				SubType:  record.SubType,
				Quantity: record.Quantity,
			}
			// 重新赋值
			subDetails = append(subDetails, sub)
			detail.SubRecord = subDetails
			remarkMap[record.Type] = detail
		} else {
			// 创建一个新的detail
			recordDetail := &model.RecordDetail{
				BabyId: record.BabyId,
				UserId: record.UserId,
				Type:   record.Type,
			}
			sub := &model.SubRecordDetail{
				SubType:  record.SubType,
				Quantity: record.Quantity,
			}
			subs := []*model.SubRecordDetail{sub}
			recordDetail.SubRecord = subs

			remarkMap[record.Type] = recordDetail
		}
	}

	var result []*model.RecordDetail
	for _, value := range remarkMap {
		result = append(result, value)
	}

	return result, nil
}
