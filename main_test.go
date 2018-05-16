package main

import (
	"github.com/Chameleon-m/hotellook/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type mockDB struct{}

func (mdb *mockDB) HotelFindFirstForGateById(gate int, id uint64) (*models.Hotel, error) {
	h := models.NewHotel()
	// todo
	return h, nil
}

func (mdb *mockDB) HotelsFind() ([]*models.Hotel, error) {

	amenities := make([]models.Amenity, 0)
	name := "Name"
	amenities = append(amenities, models.Amenity{1, 11, &name, "Value"})

	images := make([]models.Image, 0)
	w, h := 100, 100
	images = append(images, models.Image{1, 11, &w, &h, "http://url.ru"})

	hotels := make([]*models.Hotel, 0)
	hotels = append(hotels, &models.Hotel{ID: 11, GateType: 1, ForeignID: 1555, NameEn: "hotel1", AddressEn: "address1", CityEn: "city1", CountryEn: "country1", Amenities: amenities, Images: images})
	hotels = append(hotels, &models.Hotel{ID: 12, GateType: 2, ForeignID: 999, NameEn: "hotel2", AddressEn: "address2", CityEn: "city2", CountryEn: "country2"})
	return hotels, nil
}

func (mdb *mockDB) HotelSave(hotel *models.Hotel) error {
	// todo
	return nil
}

func (mdb *mockDB) NewRecord(value interface{}) bool {
	// todo
	return mdb.NewRecord(value)
}

func TestFetchAllHotels(t *testing.T) {
	router := gin.New()

	env := &Env{&mockDB{}}
	v1 := router.Group("/api/v1/hotels")
	{
		v1.GET("/", env.fetchAllHotels)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/hotels/", nil)
	router.ServeHTTP(w, req)

	expected := `{"data":[{"id":11,"gate_type":1,"foreign_id":1555,"name_en":"hotel1","name_ru":null,"address_en":"address1","address_ru":null,"city_en":"city1","city_ru":null,"country_en":"country1","country_ru":null,"country_code":null,"postcode":null,"longitude":null,"latitude":null,"stars":null,"phone":null,"thumbnail":null,"descriptions_en":"","descriptions_ru":null,"descriptions_short_en":null,"descriptions_short_ru":null,"check_in":null,"check_out":null,"rating_total":null,"reviews_count":null,"foreign_url":null,"images":[{"id":1,"hotel_id":11,"height":100,"width":100,"url":"http://url.ru"}],"amenities":[{"id":1,"hotel_id":11,"name":"Name","value":"Value"}],"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},{"id":12,"gate_type":2,"foreign_id":999,"name_en":"hotel2","name_ru":null,"address_en":"address2","address_ru":null,"city_en":"city2","city_ru":null,"country_en":"country2","country_ru":null,"country_code":null,"postcode":null,"longitude":null,"latitude":null,"stars":null,"phone":null,"thumbnail":null,"descriptions_en":"","descriptions_ru":null,"descriptions_short_en":null,"descriptions_short_ru":null,"check_in":null,"check_out":null,"rating_total":null,"reviews_count":null,"foreign_url":null,"images":null,"amenities":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}],"status":200}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
