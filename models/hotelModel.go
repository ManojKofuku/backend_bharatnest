package models

import (
	"time"

	"gorm.io/datatypes"
)



type Hotel struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name" validate:"required"`
    LocationShort string  `json:"location_short"`
    FullAddress string    `json:"full_address"`
    PinCode     string    `json:"pin_code" validate:"numeric,len=6"`
    Landmark    string    `json:"landmark"`
    Price       float64   `json:"price" validate:"gte=0"`
    Rating      float64   `json:"rating" validate:"gte=0,lte=5"`
    Image       string    `json:"image"`
    Availability bool     `json:"availability"`
    Description  string   `json:"description"`
    LocationLink string   `json:"location_link"`
    Superhost    bool     `json:"superhost"`
    HostStory    string   `json:"host_story"`
    CancellationPolicy string `json:"cancellation_policy"`
    CheckinTime time.Time `json:"checkin_time"` // stored as proper time.Time

    ManagerID  uint      `json:"manager_id"`
    Manager    Manager   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manager"`

    Rooms     []Room         `gorm:"foreignKey:HotelID" json:"rooms,omitempty"`
    Amenities datatypes.JSON `gorm:"type:json" json:"amenities"`

    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
}
// Manager Table
type Manager struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name" validate:"required"`
    About     string    `json:"about"`
    Phone     string    `json:"phone" validate:"e164"` // E.164 format validation e.g., +1234567890
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    Hotels []Hotel `gorm:"foreignKey:ManagerID" json:"hotels,omitempty"`
}

// Room Table
type Room struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    HotelID     uint           `json:"hotel_id"`
    RoomType    string         `json:"room_type"`  // e.g., Bedroom, Living Room
    BedType     string         `json:"bed_type"`   // e.g., King, Twin
    Capacity    int            `json:"capacity" validate:"gte=1"` // number of people
    Price       float64        `json:"price" validate:"gte=0"`
    Available   bool           `json:"available"`
    Description string         `json:"description"`
    Images      datatypes.JSON `gorm:"type:json" json:"images"` // Store multiple image URLs as JSON array
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
}

