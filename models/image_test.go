package models

import (
	"testing"
)

func TestImageEqual(t *testing.T) {
	var value1, value2 = "value1", "value2"
	image1 := Image{URL: value1}
	image2 := Image{URL: value1}
	image3 := Image{URL: value2}

	if image1.Equal(image2) == false {
		t.Error("1")
	}

	if image1.Equal(image3) == true {
		t.Error("2")
	}
}

func TestImageHasExist(t *testing.T) {
	var value1, value2 = "value1", "value2"
	image1 := Image{URL: value1}
	image2 := Image{URL: value1}
	image3 := Image{URL: value2}

	list := ImageList{image1, image2}
	if list.HasExist(image3) == true {
		t.Error("3")
	}

	if list.HasExist(image2) == false {
		t.Error("4")
	}
}