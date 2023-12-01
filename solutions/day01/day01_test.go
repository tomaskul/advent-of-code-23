package day01

import (
	"reflect"
	"testing"
)

func Test_getAllCalibrationValues(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := []int{12, 38, 15, 77}

	sut := &Day01{}
	actual, err := sut.getAllCalibrationValues(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}
