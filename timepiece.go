// Package timepiece provides a utility to take a time.Time Type and provide
// a struct of time pieces
package timepiece

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// TimePiece structs are objects with each 'piece' of time represented in its own field
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

// ReplaceTime takes a string, replaces any TimePiece variables in them (tokens prefaeced
// with a '$') with the actual values of the TimePiece struct passed with the
// string
// Example: "$Month" should get replaced with the actual struct's Month field
//          value
func ReplaceTime(contents string, t TimePiece) string {
	// using reflection, try to replace any word that matches a field name with
	// the TimePiece struct as long as it also is prefaced with a $
	piecesOfTime := reflect.ValueOf(&t).Elem()

	for i := 0; i < piecesOfTime.NumField(); i++ {
		fieldName := piecesOfTime.Type().Field(i).Name
		fieldValue := piecesOfTime.Field(i)

		contents = strings.Replace(
			contents,
			"$"+fieldName,
			fmt.Sprintf("%v", fieldValue),
			-1)
	}

	return contents
}

// TimeToTimePiece takes a time.Time type and returns a struct with the time broken into the
// following pieces:
//    Year (int64)
//    Month (int64)
//    Day (int64)
//    Hour (int64)
//    Minute (int64)
//    Second (float64)
//
// Usage: t := TimePiece(time.Now())
func TimeToTimePiece(t time.Time) TimePiece {
	fields := strings.Fields(t.String())
	dateParts := strings.Split(fields[0], "-")
	timeParts := strings.Split(fields[1], ":")

	var pieces TimePiece
	var err error

	pieces.Year, err = strconv.ParseInt(dateParts[0], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.Month, err = strconv.ParseInt(dateParts[1], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.PaddedMonth = fmt.Sprintf("%02d", pieces.Month)

	pieces.Day, err = strconv.ParseInt(dateParts[2], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.PaddedDay = fmt.Sprintf("%02d", pieces.Day)

	pieces.Hour, err = strconv.ParseInt(timeParts[0], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.Minute, err = strconv.ParseInt(timeParts[1], 10, 32)
	if err != nil {
		panic(err)
	}

	pieces.Second, err = strconv.ParseFloat(timeParts[2], 64)
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
//  tp := (time.Now())
//  fmt.Println(tp.String) // prints string like 2015-01-01 00:00:00
func (t *TimePiece) String(formatString ...string) string {
	if formatString == nil {
		return strconv.FormatInt(t.Year, 10) +
			"-" + fmt.Sprintf("%02d", t.Month) +
			"-" + fmt.Sprintf("%02d", t.Day) +
			" " + fmt.Sprintf("%02d", t.Hour) +
			":" + fmt.Sprintf("%02d", t.Minute) +
			":" + strconv.FormatFloat(t.Second, 'f', -1, 64)
	}

	firstAndOnly := formatString[0]
	tokens := make(map[string]string)
	tokens["%Y"] = strconv.FormatInt(t.Year, 10)
	tokens["%m"] = fmt.Sprintf("%02d", t.Month)
	tokens["%d"] = fmt.Sprintf("%02d", t.Day)
	tokens["%H"] = fmt.Sprintf("%02d", t.Hour)
	tokens["%M"] = fmt.Sprintf("%02d", t.Minute)
	tokens["%S"] = strconv.FormatFloat(t.Second, 'f', -1, 64)

	for token, replacement := range tokens {
		firstAndOnly = strings.Replace(firstAndOnly, token, replacement, -1)
	}

	return firstAndOnly
}
