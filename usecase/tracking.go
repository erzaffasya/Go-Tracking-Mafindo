package usecase

import (
	"context"

	"github.com/erzaffasya/Go-Tracking-Mafindo/models"
)

type trackingUsecase struct {
	sr models.TrackingRepo
}

func NewTrackingUsecase(sr models.TrackingRepo) *trackingUsecase {
	return &trackingUsecase{sr}
}

func (suc *trackingUsecase) Fetch(c context.Context, m *[]models.Tracking, userID uint) (err error) {
	if err = suc.sr.Fetch(c, m, userID); err != nil {
		return err
	}
	return
}

func (suc *trackingUsecase) Store(c context.Context, m *models.Tracking) (err error) {
	if err = suc.sr.Store(c, m); err != nil {
		return err
	}
	return
}

func (suc *trackingUsecase) GetByUserID(c context.Context, m *models.Tracking, id uint) (err error) {
	if err = suc.sr.GetByUserID(c, m, id); err != nil {
		return err
	}
	return
}

func (suc *trackingUsecase) Update(c context.Context, mu models.Tracking, id uint) (p models.Tracking, err error) {
	p, err = suc.sr.Update(c, mu, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (suc *trackingUsecase) Delete(c context.Context, id uint) (err error) {
	if err = suc.sr.Delete(c, id); err != nil {
		return err
	}
	return
}
