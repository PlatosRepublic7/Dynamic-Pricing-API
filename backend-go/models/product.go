package models

type Product struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:255;not null"`
	BasePrice float64 `gorm:"not null"`
	Demand    float64 `gorm:"default:1.0"`
	Price     float64 `gorm:"not null"`
}

// Simple Calculation for "dynamic" price based on demand
func (p *Product) CalculatePrice() {
	p.Price = p.BasePrice * p.Demand
}
