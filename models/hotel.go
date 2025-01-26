package models

type Hotel struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Location string `gorm:"size:255"`
	Enabled  bool   `gorm:"default:true"`
	Rooms    []Room `gorm:"foreignKey:HotelID"`
}
