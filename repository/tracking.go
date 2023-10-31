package repository

import (
	"context"
	"time"

	"github.com/erzaffasya/Go-Tracking-Mafindo/models"
	"gorm.io/gorm"
)

type TrackingRepo struct {
	db *gorm.DB
}

func NewTrackingRepo(db *gorm.DB) *TrackingRepo {
	return &TrackingRepo{db}
}

func (sr TrackingRepo) Fetch(c context.Context, m *[]models.Tracking, userID uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username", "ProfileImageUrl")
		}).
		Find(&m).Error
	if err != nil {
		return err
	}
	return
}

func (sr *TrackingRepo) Store(c context.Context, m *models.Tracking) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (sr TrackingRepo) GetByUserID(c context.Context, m *models.Tracking, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).Select("user_id").First(&m, id).Error
	if err != nil {
		return err
	}
	return
}

func (sr *TrackingRepo) Update(c context.Context, mu models.Tracking, id uint) (socialMedia models.Tracking, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	socialMedia = models.Tracking{}
	err = sr.db.Debug().WithContext(ctx).First(&socialMedia, id).Error
	if err != nil {
		return socialMedia, err
	}

	err = sr.db.Debug().WithContext(ctx).Model(&socialMedia).Where("id = ?", id).
		Updates(mu).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (sr TrackingRepo) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).First(&models.Tracking{}, id).Error
	if err != nil {
		return err
	}

	err = sr.db.Debug().WithContext(ctx).Delete(&models.Tracking{}, id).Error
	if err != nil {
		return err
	}
	return
}
