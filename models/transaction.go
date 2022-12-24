package models

import "time"

type Transaction struct {
	ID           int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID       int                 `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Remaining    int                 `json:"remaining_active"`
	Status       string              `json:"status"`
	Amount       int                 `json:"ammount"`
	Subscription string              `json:"subscription"`
	UpdatedAt    time.Time           `json:"-"`
	User         UserProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type TransactionResponse struct {
// 	Remaining int `json:"remaining" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }

// func (TransactionResponse) TableName() string {
// 	return "transactions"
// }
