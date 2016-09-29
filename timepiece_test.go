package timepiece

import (
  "reflect"
  "testing"
  "time"
)

func TestTimeToPieces(t *testing.T) {
  test_time := time.Date(2016, 12, 25, 0, 13, 46, 0, time.UTC)
  pieces := TimeToTimePiece(test_time)
  expected := TimePiece{2016, 12, 25, 0, 13, 46}

  pieces_reflection := reflect.ValueOf(&pieces).Elem()
  expected_reflection := reflect.ValueOf(&expected).Elem()

  // loop through structs and compare fields to each other
  for i := 0; i < pieces_reflection.NumField(); i++ {
    pieces_field := pieces_reflection.Field(i)
    expeced_field := expected_reflection.Field(i)

    if pieces_field.Interface() != expeced_field.Interface() {
      t.Error("expected",
              expeced_field.Interface(),
              "got",
              pieces_field.Interface(),
      )
    }
  }
}
