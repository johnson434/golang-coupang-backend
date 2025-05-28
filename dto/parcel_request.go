package dto

import (
	"golang-coupang-backend.com/m/model"
)

type CreateParcelRequest struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Address   string `json:"address"`
	Status    string `json:"status"`
	CreatedAt int    `json:"created_at"`
}

func (p CreateParcelRequest) ToModel() model.Parcel {
	return model.Parcel{
		Sender:    p.Sender,
		Receiver:  p.Receiver,
		Status:    p.Status,
		Address:   p.Address,
		CreatedAt: p.CreatedAt,
	}
}
