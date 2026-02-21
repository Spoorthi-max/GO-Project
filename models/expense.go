package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	GroupID uint  `json:"group_id"`
	PaidBy  uint  `json:"paid_by"`
	Amount  int64 `json:"amount"` // stored in cents
}

type ExpenseSplit struct {
	gorm.Model
	ExpenseID uint  `json:"expense_id"`
	UserID    uint  `json:"user_id"`
	Share     int64 `json:"share"` // stored in cents
}