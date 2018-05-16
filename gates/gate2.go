package gates

import (
	"reflect"
	"strconv"
	"github.com/Chameleon-m/hotellook/models"
)

type Gate2 struct {
	HotelID             string `csv:"hotel_id"`
	ChainID             string `csv:"chain_id"`
	ChainName           string `csv:"chain_name"`
	BrandID             string `csv:"brand_id"`
	BrandName           string `csv:"brand_name"`
	HotelName           string `csv:"hotel_name"`
	HotelFormerlyName   string `csv:"hotel_formerly_name"`
	HotelTranslatedName string `csv:"hotel_translated_name"`
	Addressline1        string `csv:"addressline1"`
	Addressline2        string `csv:"addressline2"`
	Zipcode             string `csv:"zipcode"`
	City                string `csv:"city"`
	State               string `csv:"state"`
	Country             string `csv:"country"`
	Countryisocode      string `csv:"countryisocode"`
	StarRating          string `csv:"star_rating"`
	Longitude           string `csv:"longitude"`
	Latitude            string `csv:"latitude"`
	URL                 string `csv:"url"`
	Checkin             string `csv:"checkin"`
	Checkout            string `csv:"checkout"`
	Numberrooms         string `csv:"numberrooms"`
	Numberfloors        string `csv:"numberfloors"`
	Yearopened          string `csv:"yearopened"`
	Yearrenovated       string `csv:"yearrenovated"`
	Photo1              string `csv:"photo1"`
	Photo2              string `csv:"photo2"`
	Photo3              string `csv:"photo3"`
	Photo4              string `csv:"photo4"`
	Photo5              string `csv:"photo5"`
	Overview            string `csv:"overview"`
	RatesFrom           string `csv:"rates_from"`
	ContinentID         string `csv:"continent_id"`
	ContinentName       string `csv:"continent_name"`
	CityID              string `csv:"city_id"`
	CountryID           string `csv:"country_id"`
	NumberOfReviews     string `csv:"number_of_reviews"`
	RatingAverage       string `csv:"rating_average"`
	RatesCurrency       string `csv:"rates_currency"`
}

func (gate Gate2) GetHotel(db models.Datastore) *models.Hotel {
	id, err := strconv.ParseUint(gate.HotelID, 10, 64)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}

	hotel, err := db.HotelFindFirstForGateById(2, id)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}
	hotel.ForeignID = id
	hotel.GateType = 2
	hotel.NameEn = gate.HotelName
	hotel.AddressEn = gate.Addressline1
	hotel.CityEn = gate.City
	hotel.CountryEn = gate.Country
	hotel.DescriptionsEn = gate.Overview

	var images models.ImageList
	if gate.Photo1 != "" {
		images = append(images, models.Image{URL: gate.Photo1})
	}
	if gate.Photo2 != "" {
		images = append(images, models.Image{URL: gate.Photo2})
	}
	if gate.Photo3 != "" {
		images = append(images, models.Image{URL: gate.Photo3})
	}
	if gate.Photo4 != "" {
		images = append(images, models.Image{URL: gate.Photo4})
	}
	if gate.Photo5 != "" {
		images = append(images, models.Image{URL: gate.Photo5})
	}

	if db.NewRecord(hotel) {
		hotel.Images = images
	} else if len(images) > 0 {
		for _, image := range images {
			if !hotel.Images.HasExist(image) {
				hotel.Images = append(hotel.Images, image)
			}
		}
	}

	return hotel
}

func (gate Gate2) GetDiffFields(fields []string) []string {
	return gate.difference(gate.getTags("csv"), fields)
}

func (gate Gate2) getTags(key string) []string {
	var tags []string
	t := reflect.TypeOf(gate)
	tags = make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		tags[i] = t.Field(i).Tag.Get(key)
	}

	return tags
}

func (gate Gate2) difference(a, b []string) []string {
	lenB := len(b)

	if lenB == 0 {
		return a
	}

	for i := range a {
		if lenB < i {
			return a[i:]
		}
		if a[i] != b[i] {
			return b[i:i+1]
		}
	}

	if len(a) < lenB {
		return b[len(a):]
	}

	return []string{}
}

func (gate *Gate2) Assign(fields []string) {
	// :)) todo use reflection
	gate.HotelID = fields[0]
	gate.ChainID = fields[1]
	gate.ChainName = fields[2]
	gate.BrandID = fields[3]
	gate.BrandName = fields[4]
	gate.HotelName = fields[5]
	gate.HotelFormerlyName = fields[6]
	gate.HotelTranslatedName = fields[7]
	gate.Addressline1 = fields[8]
	gate.Addressline2 = fields[9]
	gate.Zipcode = fields[10]
	gate.City = fields[11]
	gate.State = fields[12]
	gate.Country = fields[13]
	gate.Countryisocode = fields[14]
	gate.StarRating = fields[15]
	gate.Longitude = fields[16]
	gate.Latitude = fields[17]
	gate.URL = fields[18]
	gate.Checkin = fields[19]
	gate.Checkout = fields[20]
	gate.Numberrooms = fields[21]
	gate.Numberfloors = fields[22]
	gate.Yearopened = fields[23]
	gate.Yearrenovated = fields[24]
	gate.Photo1 = fields[25]
	gate.Photo2 = fields[26]
	gate.Photo3 = fields[27]
	gate.Photo4 = fields[28]
	gate.Photo5 = fields[29]
	gate.Overview = fields[30]
	gate.RatesFrom = fields[31]
	gate.ContinentID = fields[32]
	gate.ContinentName = fields[33]
	gate.CityID = fields[34]
	gate.CountryID = fields[35]
	gate.NumberOfReviews = fields[36]
	gate.RatingAverage = fields[37]
	gate.RatesCurrency = fields[38]
}