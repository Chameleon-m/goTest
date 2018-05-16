package models

type Image struct {
	ID      int    `json:"id" gorm:"PRIMARY_KEY"`
	HotelID int    `json:"hotel_id" gorm:"not null"`
	Height  *int   `json:"height"`
	Width   *int   `json:"width"`
	URL     string `json:"url"`
}

func (img Image) Equal(comparedImg Image) bool {
	return img.URL == comparedImg.URL;
}

type ImageList []Image

func (list ImageList) HasExist(image Image) bool {
	for i := 0; i < len(list); i++ {
		if list[i].Equal(image) {
			return true
		}
	}
	return false
}
