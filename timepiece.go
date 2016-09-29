// Package timepiece provides a utility to take a time.Time Type and provide
// a struct of time pieces
package timepiece

import (
  "strconv"
  "strings"
  "time"
)

type TimePiece struct {
  Year int64
	Month int64
	Day int64
	Hour int64
	Minute int64
	Second float64
}

// Given a time.Time type return a struct with the time broken into the
// following pieces:
//    Year (int64)
//    Month (int64)
//    Day (int64)
//    Hour (int64)
//    Minute (int64)
//    Second (float64)
func TimeToTimePiece(t time.Time) TimePiece {
	fields := strings.Fields(t.String())
	date_parts := strings.Split(fields[0], "-")
  time_parts := strings.Split(fields[1], ":")

  var pieces TimePiece
  var err error

  pieces.Year, err = strconv.ParseInt(date_parts[0], 10, 32)
	if err != nil {
		panic(err)
	}

  pieces.Month, err = strconv.ParseInt(date_parts[1], 10, 32)
	if err != nil {
		panic(err)
	}

  pieces.Day, err = strconv.ParseInt(date_parts[2], 10, 32)
  if err != nil {
    panic(err)
  }

  pieces.Hour, err = strconv.ParseInt(time_parts[0], 10, 32)
  if err != nil {
    panic(err)
  }

  pieces.Minute, err = strconv.ParseInt(time_parts[1], 10, 32)
  if err != nil {
    panic(err)
  }

  pieces.Second, err = strconv.ParseFloat(time_parts[2], 64)
	if err != nil {
		panic(err)
	}

  return pieces
}
