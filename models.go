package main

type Product struct {
	ID         string   `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Name       string   `json:"name" gorm:"size:50;not null"`
	Price      float32  `json:"price" gorm:"not null;default:0"`
	Icon       string   `json:"icon"`
	CategoryID string   `json:"category_id" gorm:"index;type:varchar(255)"` // Foreign Key column
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`        // Defines the relationship
}

type Category struct {
	ID   string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Name string `json:"name" gorm:"size:50;not null"`
}
