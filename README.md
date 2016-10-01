## timepiece
Golang package for returning time in pieces.  It conquers time, one piece at a
time!  There is probably a better way to do this and it may already exist.
Still this is good practice for me and helpful for me learning Go.

## Install
`go get github.com/pladdy/timepiece`

## Test
`go test`

## Docs
`godoc github.com/pladdy/timepiece`

## Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/pladdy/timepiece"
)

func main() {
	// it all starts once upon a time...
	someTime := time.Date(2016, 12, 25, 0, 13, 46, 0, time.UTC)

	// now break that bear up into some pieces!
	pieces := timepiece.TimeToTimePiece(someTime)

	// and now you can do things like
	fmt.Println(pieces.Year)     // nothing special
	fmt.Println(pieces.String()) // i know, not that great either but i was proud
	fmt.Println(pieces.String("The year %Y was cool, on %m-%d-%Y at %H:%M:%S I got some presents!"))
}
```
