// Package timepiece provides a utility to take a time.Time Type and provide
// a struct of time pieces
package timepiece

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TimePiece struct {
	Year        int64
	Month       int64
	PaddedMonth string
	Day         int64
	PaddedDay   string
	Hour        int64
	Minute      int64
	Second      float64
}

// Given a time.Time type return a struct with the time broken into the
// following pieces:
//    Year (int64)
//    Month (int64)
//    Day (int64)
//    Hour (int64)
//    Minute (int64)
//    Second (float64)
//
// Usage: timePiece = TimePiece(time.Now())
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

	pieces.PaddedMonth = fmt.Sprintf("%02d", pieces.Month)
	if err != nil {
		panic(err)
	}

	pieces.Day, err = strconv.ParseInt(date_parts[2], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.PaddedDay = fmt.Sprintf("%02d", pieces.Day)
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

// String takes an optional format string; if none is provided a default
// string version of the time pieces is returned.  The string format can be
// set using the below tokens:
//    %Y year
//    %m month
//    %d day
//    %H hour
//    %M minute
//    %S seconds
//
// Example:
//  tp := TimePiece(time.Now())
//  fmt.Println(tp.String) // prints string like 2015-01-01 00:00:00
func (timePiece *TimePiece) String(formatString ...string) string {
	if formatString == nil {
		return strconv.FormatInt(timePiece.Year, 10) +
			"-" + padSingleInt(timePiece.Month) +
			"-" + padSingleInt(timePiece.Day) +
			" " + padSingleInt(timePiece.Hour) +
			":" + padSingleInt(timePiece.Minute) +
			":" + strconv.FormatFloat(timePiece.Second, 'f', -1, 64)
	} else {
		firstAndOnly := formatString[0]
		tokens := make(map[string]string)
		tokens["%Y"] = strconv.FormatInt(timePiece.Year, 10)
		tokens["%m"] = padSingleInt(timePiece.Month)
		tokens["%d"] = padSingleInt(timePiece.Day)
		tokens["%H"] = padSingleInt(timePiece.Hour)
		tokens["%M"] = padSingleInt(timePiece.Minute)
		tokens["%S"] = strconv.FormatFloat(timePiece.Second, 'f', -1, 64)

		for token, replacement := range tokens {
			firstAndOnly = strings.Replace(firstAndOnly, token, replacement, -1)
		}

		return firstAndOnly
	}
}

/* Private */

// helpfer function to pad a number with a leading 0
func padSingleInt(number int64) string {
	numberString := strconv.FormatInt(number, 10)
	if len(numberString) == 1 {
		return "0" + numberString
	}
	return numberString
}
