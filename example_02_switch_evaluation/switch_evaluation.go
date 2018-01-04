package main

import (
	"fmt"
	"time"
)

func howManyDays(day time.Weekday) {
	fmt.Println("When is", day, "?")
	today := time.Now().Weekday()
	switch day {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		var days int
		if today < day {
			days = (int)(day - today)
		} else {
			days = (int)(time.Saturday - today + day + 1)
		}
		fmt.Println("In ", days, "days.")
	}
}

func main() {
	var Weekdays = []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}
	for _, day := range Weekdays {
		howManyDays(day)
	}
}
