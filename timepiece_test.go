package timepiece

import (
	"reflect"
	"testing"
	"time"
)

var (
	test_time = time.Date(2016, 12, 25, 0, 13, 46, 0, time.UTC)
	pieces    = TimeToTimePiece(test_time)
)

func TestTimeToPieces(t *testing.T) {
	expected := TimePiece{2016, 12, "12", 25, "25", 0, 13, 46}

	pieces_reflection := reflect.ValueOf(&pieces).Elem()
	expected_reflection := reflect.ValueOf(&expected).Elem()

	// loop through structs and compare fields to each other
	for i := 0; i < pieces_reflection.NumField(); i++ {
		pieces_field := pieces_reflection.Field(i)
		expeced_field := expected_reflection.Field(i)

		if pieces_field.Interface() != expeced_field.Interface() {
			t.Error(
				"expected",
				expeced_field.Interface(),
				"got",
				pieces_field.Interface(),
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
