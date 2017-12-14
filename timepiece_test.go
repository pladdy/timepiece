package timepiece

import (
	"reflect"
	"testing"
	"time"
)

var (
	testTime = time.Date(2016, 12, 25, 0, 13, 46, 0, time.UTC)
	pieces   = TimeToTimePiece(testTime)
)

func TestReplaceTime(t *testing.T) {
	result := string(ReplaceTime("The Year is $Year", pieces))
	expectedString := "The Year is 2016"

	if result != expectedString {
		t.Error("expected", expectedString, "got", result)
	}
}

func TestTimeToPieces(t *testing.T) {
	expected := TimePiece{2016, 12, "12", 25, "25", 0, 13, 46}

	piecesReflection := reflect.ValueOf(&pieces).Elem()
	expectedReflection := reflect.ValueOf(&expected).Elem()

	// loop through structs and compare fields to each other
	for i := 0; i < piecesReflection.NumField(); i++ {
		piecesField := piecesReflection.Field(i)
		expectedField := expectedReflection.Field(i)

		if piecesField.Interface() != expectedField.Interface() {
			t.Error(
				"expected",
				expectedField.Interface(),
				"got",
				piecesField.Interface(),
			)
		}
	}
}

func TestTimePieceStringDefault(t *testing.T) {
	expected := "2016-12-25 00:13:46"
	result := pieces.String()

	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}

func TestTimePieceString(t *testing.T) {
	expected := "The year is 2016 and the month is 12"
	result := pieces.String("The year is %Y and the month is %m")

	if result != expected {
		t.Error("expected", expected, "got", result)
	}

	expected = "H:M:S -> 00:13:46"
	result = pieces.String("H:M:S -> %H:%M:%S")

	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}
