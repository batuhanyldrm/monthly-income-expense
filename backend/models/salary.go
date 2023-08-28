package models

type Salary struct {
	ID        string `json:"id" bson:"id"`
	Salary    string `json:"salary" bson:"salary"`
	Debit     string `json:"debit" bson:"debit"`
	MoneyGain string `json:"moneyGain" bson:"moneyGain"`
}

type User struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
