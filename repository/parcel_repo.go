package repository

import (
	"context"

	"golang-coupang-backend.com/m/model"
)

type ParcelRepository interface {
	Create(ctx context.Context, p model.Parcel) error
	// Update(ctx context.Context, parcel model.Parcel) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Parcel, error)
	GetByID(ctx context.Context, id int) (model.Parcel, error)
}
