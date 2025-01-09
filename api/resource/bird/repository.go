package bird

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) ListAll() (Birds, error) {
	birds := make([]*Bird, 0)
	if err := r.db.Find(&birds).Error; err != nil {
		return nil, err
	}

	return birds, nil
}

// func (r *Repository) Create(bird *Bird) (*Bird, error) {
// 	if err := r.db.Create(bird).Error; err != nil {
// 		return nil, err
// 	}

// 	return bird, nil
// }

func (r *Repository) GetSingle(id uuid.UUID) (*Bird, error) {
	bird := &Bird{}
	if err := r.db.Where("id = ?", id).First(&bird).Error; err != nil {
		return nil, err
	}

	return bird, nil
}

// func (r *Repository) Update(bird *Bird) (int64, error) {
// 	result := r.db.Model(&Bird{}).
// 		Select("CommonName", "Confidence", "RecordingDate", "InputDevice", "Description").
// 		Where("id = ?", bird.ID).
// 		Updates(bird)

// 	return result.RowsAffected, result.Error
// }

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Bird{})

	return result.RowsAffected, result.Error

}
