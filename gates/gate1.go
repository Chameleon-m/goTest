package gates

import (
	"strconv"
	"strings"
	"github.com/Chameleon-m/hotellook/models"
)

type Gate1 struct {
	CheckInTime  string `json:"check_in_time"`
	CheckOutTime string `json:"check_out_time"`
	CountryCode  string `json:"country_code"`
	Email        string `json:"email"`
	En struct {
		Address string `json:"address"`
		Amenities []struct {
			Amenities []string `json:"amenities"`
			GroupName string   `json:"group_name"`
			GroupSlug string   `json:"group_slug"`
		} `json:"amenities"`
		City               string `json:"city"`
		Country            string `json:"country"`
		Description        string `json:"description"`
		DescriptionShort   string `json:"description_short"`
		Name               string `json:"name"`
		PolicyDescription  string `json:"policy_description"`
		RatingTotalVerbose string `json:"rating_total_verbose"`
		RoomGroups []struct {
			Amenities struct {
			} `json:"amenities"`
			ImageListTmpl []interface{} `json:"image_list_tmpl"`
			Name          string        `json:"name"`
			RoomGroupID   int           `json:"room_group_id"`
		} `json:"room_groups"`
	} `json:"en"`
	Hotelpage string `json:"hotelpage"`
	ID        string `json:"id"`
	IDCrc64   string `json:"id_crc64"`
	Images []struct {
		Height     int    `json:"height"`
		OrigHeight int    `json:"orig_height"`
		OrigURL    string `json:"orig_url"`
		OrigWidth  int    `json:"orig_width"`
		URL        string `json:"url"`
		Width      int    `json:"width"`
	} `json:"images"`
	Kind      string  `json:"kind"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	LowRate   int     `json:"low_rate"`
	Matching struct {
		Booking   int `json:"Booking"`
		Carsolize int `json:"Carsolize"`
		Expedia   int `json:"Expedia"`
		HotelsPro int `json:"HotelsPro"`
		Laterooms int `json:"Laterooms"`
		Ostrovok  int `json:"Ostrovok"`
	} `json:"matching"`
	Phone           string      `json:"phone"`
	ProviderEngines interface{} `json:"provider_engines"`
	Rating struct {
		Count int `json:"count"`
		Detailed struct {
			Cleanness float64 `json:"cleanness"`
			Comfort   float64 `json:"comfort"`
			Location  float64 `json:"location"`
			Personnel float64 `json:"personnel"`
			Price     float64 `json:"price"`
			Services  float64 `json:"services"`
		} `json:"detailed"`
		Exists                   bool    `json:"exists"`
		OtherReviewsCount        int     `json:"other_reviews_count"`
		OurPublishedReviewsCount int     `json:"our_published_reviews_count"`
		OurReviewsCount          int     `json:"our_reviews_count"`
		ReviewsCount             int     `json:"reviews_count"`
		Total                    float64 `json:"total"`
	} `json:"rating"`
	RegionCategory string `json:"region_category"`
	Ru struct {
		Address string `json:"address"`
		Amenities []struct {
			Amenities []string `json:"amenities"`
			GroupName string   `json:"group_name"`
			GroupSlug string   `json:"group_slug"`
		} `json:"amenities"`
		City               string `json:"city"`
		Country            string `json:"country"`
		Description        string `json:"description"`
		DescriptionShort   string `json:"description_short"`
		Name               string `json:"name"`
		PolicyDescription  string `json:"policy_description"`
		RatingTotalVerbose string `json:"rating_total_verbose"`
		RoomGroups []struct {
			Amenities struct {
			} `json:"amenities"`
			ImageListTmpl []interface{} `json:"image_list_tmpl"`
			Name          string        `json:"name"`
			RoomGroupID   int           `json:"room_group_id"`
		} `json:"room_groups"`
	} `json:"ru"`
	RegionID   int                    `json:"region_id"`
	StarRating int                    `json:"star_rating"`
	Thumbnail  string                 `json:"thumbnail"`
	X          map[string]interface{} `json:"-" gorm:"-"`
}

func (gate Gate1) GetHotel(db models.Datastore) *models.Hotel {
	id, err := strconv.ParseUint(gate.ID, 10, 64)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}

	hotel, err := db.HotelFindFirstForGateById(1, id)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}
	hotel.ForeignID = id
	hotel.GateType = 1
	hotel.NameEn = gate.En.Name
	hotel.AddressEn = gate.En.Address
	hotel.CityEn = gate.En.City
	hotel.CountryEn = gate.En.Country
	hotel.DescriptionsEn = gate.En.Description

	if len(gate.En.Amenities) > 0 {
		for _, v := range gate.En.Amenities {
			gp := v.GroupName
			amenity := models.Amenity{Name: &gp, Value: strings.Join(v.Amenities, ",")}
			if db.NewRecord(hotel) || !hotel.Amenities.HasExist(amenity) {
				hotel.Amenities = append(hotel.Amenities, amenity)
			}
		}
	}

	if len(gate.Images) > 0 {
		for _, v := range gate.Images {
			h, w := v.Height, v.Width
			image := models.Image{Height: &h, Width: &w, URL: v.URL}
			if db.NewRecord(hotel) || !hotel.Images.HasExist(image) {
				hotel.Images = append(hotel.Images, image)
			}
		}
	}

	return hotel
}
