package repositories

import (
	"gorm.io/gorm"
)

type FindAllOptions struct {
	limit     *int
	skip      *int
	relations []string
}

type Repository[T any] interface {
	Create(entity *T) (*T, error)
	FindById(id uint, preload *[]string) (*T, error)
	FindOne(query *T, preload *[]string) (*T, error)
	FindAll(options *FindAllOptions) *[]T
	Update(entity *T) (*T, error)
	Delete(id uint) (bool, error)
}

type repository[T any] struct {
	db *gorm.DB
}

func (r repository[T]) FindOne(query *T, preload *[]string) (*T, error) {
	var entity T
	tx := r.db.Where(&query)
	tx = preloadColumns(tx, preload)
	err := tx.Take(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r repository[T]) FindById(id uint, preload *[]string) (*T, error) {
	var entity T
	tx := r.db.Where("id = ?", id)
	tx = preloadColumns(tx, preload)
	err := tx.Take(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r repository[T]) FindAll(options *FindAllOptions) *[]T {
	var entities []T
	var tx *gorm.DB
	if options.limit != nil {
		tx = r.db.Limit(*options.limit).Offset(*options.skip)
	}
	tx = preloadColumns(tx, &options.relations)
	err := tx.Find(&entities).Error
	if err != nil {
		return &[]T{}
	}
	return &[]T{}
}
func (r repository[T]) FindAllWhere(query *T, limit int, skip int, preload *[]string) *[]T {
	var entities []T
	tx := r.db.Where(query).Limit(limit).Offset(skip)
	tx = preloadColumns(tx, preload)
	err := tx.Find(&entities).Error
	if err != nil {
		return &[]T{}
	}
	return &entities
}

func (r repository[T]) Create(entity *T) (*T, error) {
	err := r.db.Debug().Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r repository[T]) CreateMany(entity *[]T) (*[]T, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r repository[T]) Update(entity *T) (*T, error) {
	err := r.db.Model(&entity).Updates(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (r repository[T]) UpdateMany(entity *[]T) (*[]T, error) {
	err := r.db.Model(&entity).Updates(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (r repository[T]) Delete(id uint) (bool, error) {
	var entity T
	err := r.db.Where("id = ?", id).Delete(&entity).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r repository[T]) FindOneOrCreate(entity *T) (*T, error) {
	err := r.db.Where(&entity).FirstOrCreate(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r repository[T]) Count(query *T) int64 {
	var count int64
	tx := r.db.Where(&query)
	err := tx.Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{
		db: db,
	}
}

func preloadColumns(tx *gorm.DB, preload *[]string) *gorm.DB {
	if preload != nil {
		for _, s := range *preload {
			tx.Preload(s)
		}
	}
	return tx
}
