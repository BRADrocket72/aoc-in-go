package main

import (
	"image"
	"testing"
)

func Test_calcDist(t *testing.T) {
	val := calcDistance(image.Pt(0, 0), image.Pt(1, 1))
	if val != 1 {
		t.Error("test error")
	}

	val2 := calcDistance(image.Pt(2, 2), image.Pt(0, 4))
	if val2 != 2 {
		t.Error("test error")
	}

}

func Test_findDoubles(t *testing.T) {
	array1 := []int{1, 5, 2, 10}
	val := findDoubles(array1)
	if val[0] != 2 && val[1] != 3 {
		t.Error("test error")
	}

	array2 := []int{10, 2, 5, 1}
	val2 := findDoubles(array2)
	if val2[0] != 0 && val2[1] != 1 {
		t.Error("test error")
	}

}

//227268405 too high
