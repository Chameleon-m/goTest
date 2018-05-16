package models

import (
	"time"
)

type HotelInterface interface {
	getHotel() Hotel
}

func NewHotel() *Hotel {
	return &Hotel{}
}

type Hotel struct {
	ID                  uint     `json:"id" gorm:"not null;AUTO_INCREMENT;PRIMARY_KEY"`
	GateType            uint     `json:"gate_type" gorm:"not null"`
	ForeignID           uint64   `json:"foreign_id" gorm:"not null"`
	NameEn              string   `json:"name_en" gorm:"type:varchar(255);not null"`
	NameRu              *string  `json:"name_ru" gorm:"type:varchar(255)"`
	AddressEn           string   `json:"address_en" gorm:"type:varchar(255);not null"`
	AddressRu           *string  `json:"address_ru" gorm:"type:varchar(255)"`
	CityEn              string   `json:"city_en" gorm:"type:varchar(255);not null"`
	CityRu              *string  `json:"city_ru" gorm:"type:varchar(255)"`
	CountryEn           string   `json:"country_en" gorm:"type:varchar(255);not null"`
	CountryRu           *string  `json:"country_ru" gorm:"type:varchar(255)"`
	CountryCode         *string  `json:"country_code"`
	Postcode            *string  `json:"postcode"`
	Longitude           *float64 `json:"longitude"`
	Latitude            *float64 `json:"latitude"`
	Stars               *int64   `json:"stars"`
	Phone               *string  `json:"phone"`
	Thumbnail           *string  `json:"thumbnail"`
	DescriptionsEn      string   `json:"descriptions_en" gorm:"not null"`
	DescriptionsRu      *string  `json:"descriptions_ru"`
	DescriptionsShortEn *string  `json:"descriptions_short_en"`
	DescriptionsShortRu *string  `json:"descriptions_short_ru"`
	CheckIn             *string  `json:"check_in" gorm:"type: time"`
	CheckOut            *string  `json:"check_out" gorm:"type: time"`
	RatingTotal         *float64 `json:"rating_total"`
	ReviewsCount        *int64   `json:"reviews_count" gorm:"DEFAULT:'0''"`
	ForeignURL          *string  `json:"foreign_url"`

	Images    ImageList   `json:"images" gorm:"foreignkey:HotelId"`
	Amenities AmenityList `json:"amenities" gorm:"foreignkey:HotelId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (db *DB) HotelFindFirstForGateById(gate int, id uint64) (*Hotel, error) {
	h := NewHotel()
	//.Set("gorm:query_option", "FOR UPDATE")
	db.Preload("Images").Preload("Amenities").First(h, "gate_type = ? AND foreign_id = ?", gate, id)
	if err := db.Error; err != nil {
		return nil, err
	}
	return h, nil
}

// вернёт gorm.ErrRecordNotFound если записи не буду найдены
func (db *DB) HotelsFind() ([]*Hotel, error) {
	var hotels []*Hotel
	db.Preload("Images").Preload("Amenities").Find(&hotels)
	if err := db.Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func (db *DB) HotelSave(hotel *Hotel) (error) {
	db.Save(hotel)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
