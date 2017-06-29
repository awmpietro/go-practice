package gigasecond

import (
	"fmt"
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(myTime time.Time) time.Time {
	gs := myTime.Add(time.Second * time.Duration(math.Pow(10, 9)))
	fmtTime, err := time.Parse("2006-01-02T15:04:05", gs.Format("2006-01-02T15:04:05"))
	if err != nil {
		fmt.Println(err)
	}
	return fmtTime
}
