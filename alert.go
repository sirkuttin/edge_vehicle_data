package data

import "github.com/go-ozzo/ozzo-validation"

type Alert struct {
	VehicleId uint32 `json:"vehicle_id"`
	AlertId   uint32 `json:"alert_id"`
}

func (alert Alert) Validate() error {
	return validation.ValidateStruct(&alert,
		validation.Field(&alert.VehicleId, validation.Required),
		validation.Field(&alert.AlertId, validation.Required),
	)
}
