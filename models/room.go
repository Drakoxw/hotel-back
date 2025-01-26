package models

type Room struct {
	ID          uint    `gorm:"primaryKey"`
	HotelID     uint    `gorm:"not null"`
	RoomType    string  `gorm:"size:255"`
	BaseCost    float64 `gorm:"not null"`
	Taxes       float64
	Enabled     bool `gorm:"default:true"`
	Description string
}
