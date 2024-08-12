package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In time file")

	var currentTime = time.Now()

	fmt.Println("Current time without any formatting is: ", currentTime) // Prints 2024-08-13 00:05:46.017169 +0530 IST m=+0.000193043

	formattedCurrentTime := currentTime.Format("01-02-2006 15:04:05 Monday") // A little weird but must be followed like this to get mm dd yyyy time and day
	fmt.Println("Formatted current time: ", formattedCurrentTime)

	createdDate := time.Date(
		2020,
		time.September,
		19,
		21,
		30,
		30,
		0,
		time.Local,
	)

	fmt.Println("Created Date: ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
