package main

import "testing"
var expected [15]string = [15]string{"up", "left", "down", "left", "down", "right", "down", "left", "down", "right", "up", "right", "down", "right", "down"}

func TestFirstFifteen(t *testing.T) {
	fractal := new(DragonFractal)

	for i:=0; i<15; i+=1 {
		result := fractal.Next()
		if result != expected[i] {
			t.Errorf("Expected `%s` but got `%s` for step `%d`",expected[i], result, i); 
		}
	}
}
func TestFifty(t *testing.T) {
	fractal := new(DragonFractal)
	
	for i:=0; i<49; i+=1 {
		fractal.Next()
	}
	result := fractal.Next()
	if result != "right" {
		t.Errorf("Expected `%s` but got `%s` for step `%d`", "right", result, 50); 
	}
}
