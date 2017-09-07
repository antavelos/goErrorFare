package main

import (
	"fmt"
)

var URLS = []string{
	"http://www.secretflying.com/feed",
	"http://cheapflightslab.com/feed",
	"http://travelpirates.com/feed",
	"http://www.flynous.com/feed",
	"http://smartdealtrip.com/feed",
	"http://www.rushflights.com/feed",
	"http://www.vliegdeals.be/feed",
	"http://www.myholidayguru.co.uk/feed",
	"http://www.fly4free.com/feed",
}

func main() {
	ch := make(chan ErrorFare)
	for _, url := range URLS {
		go GetErrorFare(url, ch)
	}
	for ef := range ch {
		fmt.Printf("%s\n", ef.Title)
	}
}
