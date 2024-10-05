package models

import "time"

type LoginStatus string

const (
	LoginStatusSuccess LoginStatus = "success"
	LoginStatusFailure LoginStatus = "failure"
)

type LoginHistory struct {
	ID         uint
	CustomerID uint
	Status     LoginStatus `gorm:"type:login_status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
