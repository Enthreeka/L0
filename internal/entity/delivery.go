package entity

type Delivery struct {
	OrderUID string `json:"order_uid" db:"order_uid"`
	Phone    string `json:"phone" db:"phone"`
	Name     string `json:"name" db:"name"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Address  string `json:"address" db:"address"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}
