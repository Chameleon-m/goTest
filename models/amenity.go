package models

type Amenity struct {
	ID      int     `json:"id" gorm:"PRIMARY_KEY"`
	HotelID int     `json:"hotel_id" gorm:"not null"`
	Name    *string `json:"name"`
	Value   string  `json:"value"`
}

func (amenity Amenity) Equal(comparedAmenity Amenity) bool {
	n1, n2 := *amenity.Name, *comparedAmenity.Name
	return n1 == n2 && amenity.Value == comparedAmenity.Value;
}

type AmenityList []Amenity

func (list AmenityList) HasExist(amenity Amenity) bool {
	for i := 0; i < len(list); i++ {
		if list[i].Equal(amenity) {
			return true
		}
	}
	return false
}
