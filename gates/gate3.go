package gates

import (
	"encoding/xml"
	"strconv"
	"github.com/Chameleon-m/hotellook/models"
)

type Gate3 struct {
	XMLName            xml.Name       `xml:"hotel,omitempty"`
	Address            string         `xml:"address,omitempty"`
	Amenities          *Amenities     `xml:"amenities,omitempty"`
	City               *City          `xml:"city,omitempty"`
	CityID             string         `xml:"cityid,omitempty"`
	Country            *Country       `xml:"country,omitempty"`
	CountryID          string         `xml:"countryid,omitempty"`
	CountryTwoCharCode string         `xml:"countrytwocharcode,omitempty"`
	Descriptions       *Descriptions  `xml:"descriptions,omitempty"`
	Fax                string         `xml:"fax,omitempty"`
	ID                 string         `xml:"id,omitempty"`
	Latitude           string         `xml:"latitude,omitempty"`
	Longitude          string         `xml:"longitude,omitempty"`
	Name               string         `xml:"name,omitempty"`
	Phone              string         `xml:"phone,omitempty"`
	Photos             *Photos        `xml:"photos,omitempty"`
	Postcode           string         `xml:"postcode,omitempty"`
	Stars              string         `xml:"stars,omitempty"`
	Statename          string         `xml:"statename,omitempty"`
	Thumbnail          string         `xml:"thumbnail,omitempty"`
	X                  []UnknownField `xml:",any"`
}

func (gate Gate3) GetHotel(db models.Datastore) *models.Hotel {
	id, err := strconv.ParseUint(gate.ID, 10, 64)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}

	hotel, err := db.HotelFindFirstForGateById(3, id)
	if err != nil {
		// todo добавить возвращение ошибок или панику.
	}
	hotel.ForeignID = id
	hotel.GateType = 3
	hotel.NameEn = gate.Name
	hotel.AddressEn = gate.Address
	hotel.CityEn = gate.City.En
	hotel.CountryEn = gate.Country.En
	hotel.DescriptionsEn = gate.Descriptions.En

	if gate.Amenities.En != "" {
		amenity := models.Amenity{Value: gate.Amenities.En}
		if db.NewRecord(hotel) || !hotel.Amenities.HasExist(amenity) {
			hotel.Amenities = append(hotel.Amenities, amenity)
		}
	}

	if len(gate.Photos.Photo) > 0 {
		for _, v := range gate.Photos.Photo {
			image := models.Image{URL: v.Url}
			if db.NewRecord(hotel) || !hotel.Images.HasExist(image) {
				hotel.Images = append(hotel.Images, image)
			}
		}
	}

	return hotel
}

type Amenities struct {
	XMLName xml.Name `xml:"amenities,omitempty" json:"amenities,omitempty"`
	De      string   `xml:"de,omitempty" json:"de,omitempty"`
	En      string   `xml:"en,omitempty" json:"en,omitempty"`
	Es      string   `xml:"es,omitempty" json:"es,omitempty"`
	Fr      string   `xml:"fr,omitempty" json:"fr,omitempty"`
	It      string   `xml:"it,omitempty" json:"it,omitempty"`
	Nl      string   `xml:"nl,omitempty" json:"nl,omitempty"`
}

type City struct {
	XMLName xml.Name `xml:"city,omitempty" json:"city,omitempty"`
	De      string   `xml:"de,omitempty" json:"de,omitempty"`
	En      string   `xml:"en,omitempty" json:"en,omitempty"`
	Es      string   `xml:"es,omitempty" json:"es,omitempty"`
	Fr      string   `xml:"fr,omitempty" json:"fr,omitempty"`
	It      string   `xml:"it,omitempty" json:"it,omitempty"`
	Nl      string   `xml:"nl,omitempty" json:"nl,omitempty"`
}

type Country struct {
	XMLName xml.Name `xml:"country,omitempty" json:"country,omitempty"`
	De      string   `xml:"de,omitempty" json:"de,omitempty"`
	En      string   `xml:"en,omitempty" json:"en,omitempty"`
	Es      string   `xml:"es,omitempty" json:"es,omitempty"`
	Fr      string   `xml:"fr,omitempty" json:"fr,omitempty"`
	It      string   `xml:"it,omitempty" json:"it,omitempty"`
	Nl      string   `xml:"nl,omitempty" json:"nl,omitempty"`
}

type Descriptions struct {
	XMLName xml.Name `xml:"descriptions,omitempty" json:"descriptions,omitempty"`
	De      string   `xml:"de,omitempty" json:"de,omitempty"`
	En      string   `xml:"en,omitempty" json:"en,omitempty"`
	Es      string   `xml:"es,omitempty" json:"es,omitempty"`
	Fr      string   `xml:"fr,omitempty" json:"fr,omitempty"`
	It      string   `xml:"it,omitempty" json:"it,omitempty"`
	Nl      string   `xml:"nl,omitempty" json:"nl,omitempty"`
}

type UnknownField struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
	Attr    string   `xml:",any"`
}

type Photo struct {
	XMLName xml.Name `xml:"photo,omitempty" json:"photo,omitempty"`
	Title   string   `xml:"title,omitempty" json:"title,omitempty"`
	Url     string   `xml:"url,omitempty" json:"url,omitempty"`
}

type Photos struct {
	XMLName xml.Name `xml:"photos,omitempty" json:"photos,omitempty"`
	Photo   []*Photo `xml:"photo,omitempty" json:"photo,omitempty"`
}
