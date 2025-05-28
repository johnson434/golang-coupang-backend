package service

import (
	"context"

	"golang-coupang-backend.com/m/model"
	"golang-coupang-backend.com/m/repository"
)

type ParcelService struct {
	Repo repository.ParcelRepository
}

func NewParcelService(r repository.ParcelRepository) *ParcelService {
	return &ParcelService{Repo: r}
}

func (service *ParcelService) CreateParcel(ctx context.Context, p model.Parcel) error {
	return service.Repo.Create(ctx, p)
}

//func (service *ParcelService) UpdateParcel(ctx context.Context, p model.Parcel) error {
//	// Parcel 업데이트
//	return service.Repo.Update(ctx, p)
//}

func (service *ParcelService) DeleteParcel(ctx context.Context, id int) error {
	// Parcel 삭제
	return service.Repo.Delete(ctx, id)
}

func (service *ParcelService) GetAllParcels(ctx context.Context) ([]model.Parcel, error) {
	// 모든 Parcel 조회
	return service.Repo.GetAll(ctx)
}

func (service *ParcelService) GetParcelByID(ctx context.Context, id int) (model.Parcel, error) {
	// Parcel ID로 조회
	return service.Repo.GetByID(ctx, id)
}
