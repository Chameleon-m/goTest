package models

import (
	"testing"
)

func TestAmenityEqual(t *testing.T) {
	name1, name2 := "name1", "name2"
	value1, value2 := "value1", "value2"
	amenity1 := Amenity{Name: &name1, Value: value1}
	amenity2 := Amenity{Name: &name1, Value: value1}
	amenity3 := Amenity{Name: &name2, Value: value2}

	if amenity1.Equal(amenity2) == false {
		t.Error("1")
	}

	if amenity1.Equal(amenity3) == true {
		t.Error("2")
	}
}

func TestAmenityHasExist(t *testing.T) {
	name1, name2 := "name1", "name2"
	value1, value2 := "value1", "value2"
	amenity1 := Amenity{Name: &name1, Value: value1}
	amenity2 := Amenity{Name: &name1, Value: value1}
	amenity3 := Amenity{Name: &name2, Value: value2}

	list := AmenityList{amenity1, amenity2}
	if list.HasExist(amenity3) == true {
		t.Error("3")
	}

	if list.HasExist(amenity2) == false {
		t.Error("4")
	}
}