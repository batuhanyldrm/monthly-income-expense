package models

import "time"

type Salary struct {
	ID        string    `json:"id" bson:"id"`
	Salary    string    `json:"salary" bson:"salary"`
	Debit     string    `json:"debit" bson:"debit"`
	MoneyGain string    `json:"moneyGain" bson:"moneyGain"`
	Users     []User    `json:"users" bson:"users"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type SalaryDTO struct {
	Salary    string    `json:"salary" bson:"salary"`
	Debit     string    `json:"debit" bson:"debit"`
	MoneyGain string    `json:"moneyGain" bson:"moneyGain"`
	Users     []User    `json:"users" bson:"users"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type User struct {
	ID        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
