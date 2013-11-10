package main

import (
	"errors"
	"testing"
)

func assertColor(pixel Pixel, rgb ...byte) (error, byte) {
	if pixel.Color().Red != rgb[0] {
		return errors.New("Red pixel should not be"), pixel.Color().Red
	}

	if pixel.Color().Green != rgb[1] {
		return errors.New("Green pixel should not be"), pixel.Color().Green
	}

	if pixel.Color().Blue != rgb[2] {
		return errors.New("Blue pixel should not be"), pixel.Color().Blue
	}

	return nil, 0
}

func TestBasicRGBCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 13, 26, 52, 31, 33, 41,
	}
	header := Header{"RGB", 3}
	picture := ParseImage(data, header)

	if err, value := assertColor(picture.InspectPixel(0, 0), 0, 12, 244); err != nil {
		t.Error(err, value)
	}
	if err, value := assertColor(picture.InspectPixel(1, 0), 13, 26, 52); err != nil {
		t.Error(err, value)
	}
}
func TestBasicGRBCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 13, 26, 52, 31, 33, 41,
	}
	header := Header{"GRB", 3}
	picture := ParseImage(data, header)

	if err, value := assertColor(picture.InspectPixel(0, 0), 12, 0, 244); err != nil {
		t.Error(err, value)
	}
	if err, value := assertColor(picture.InspectPixel(2, 0), 33, 31, 41); err != nil {
		t.Error(err, value)
	}
}
func TestBasicRGBACall(t *testing.T) {
	data := []byte{
		0, 12, 244, 128, 14, 26, 52, 127, 31, 33, 41, 255, 36, 133, 241, 255,
	}
	header := Header{"RGBA", 4}
	picture := ParseImage(data, header)

	first_pixel := picture.InspectPixel(0, 0)
	if err, value := assertColor(first_pixel, 0, 6, 122); err != nil {
		t.Error(err, value)
	}

	second_pixel := picture.InspectPixel(3, 0)
	if err, value := assertColor(second_pixel, 36, 133, 241); err != nil {
		t.Error(err, value)
	}

	third_pixel := picture.InspectPixel(2, 0)
	if err, value := assertColor(third_pixel, 31, 33, 41); err != nil {
		t.Error(err, value)
	}
}

func TestBasicRGBARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}
	header := Header{"RGBA", 2}
	picture := ParseImage(data, header)

	pixel := picture.InspectPixel(1, 1)
	if err, value := assertColor(pixel, 36, 133, 241); err != nil {
		t.Error(err, value)
	}
	pixel1 := picture.InspectPixel(0, 1)
	if err, value := assertColor(pixel1, 31, 33, 41); err != nil {
		t.Error(err, value)
	}
	pixel2 := picture.InspectPixel(1, 0)
	if err, value := assertColor(pixel2, 6, 12, 25); err != nil {
		t.Error(err, value)
	}
}

func TestBasicBGRARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}
	header := Header{"BGRA", 2}
	picture := ParseImage(data, header)

	pixel := picture.InspectPixel(1, 1)
	if err, value := assertColor(pixel, 241, 133, 36); err != nil {
		t.Error(err, value)
	}
	pixel1 := picture.InspectPixel(1, 0)
	if err, value := assertColor(pixel1, 25, 12, 6); err != nil {
		t.Error(err, value)
	}
}
