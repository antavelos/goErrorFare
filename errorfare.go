package main

import (
	"github.com/mmcdole/gofeed"
)

type Site struct {
	Name string
	Url  string
}

type ErrorFare struct {
	Title  string
	Url    string
	From   string
	To     string
	Amount string
}

func GetErrorFare(url string, ch chan ErrorFare) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}
	var errorFeeds []ErrorFare
	for _, item := range feed.Items {
		ch <- ErrorFare{
			Title: item.Title,
			Url:   item.Link,
		}
	}
	return nil
}
