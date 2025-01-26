package models

type Reservation struct {
	ID               uint   `gorm:"primaryKey"`
	RoomID           uint   `gorm:"not null"`
	GuestName        string `gorm:"size:255"`
	CheckInDate      string `gorm:"not null"`
	CheckOutDate     string `gorm:"not null"`
	Email            string
	Phone            string
	EmergencyContact EmergencyContact `gorm:"foreignKey:ReservationID"`
}

type EmergencyContact struct {
	ID            uint   `gorm:"primaryKey"`
	ReservationID uint   `gorm:"not null"`
	Name          string `gorm:"size:255"`
	Phone         string `gorm:"size:255"`
}
