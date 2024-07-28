package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func urlToFeed(url string) (RSSFFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFFeed{}, err
	}

	rssFeed := RSSFFeed{}
	err = xml.Unmarshal(dat, &rssFeed)

	if err != nil {
		return RSSFFeed{}, err
	}

	return rssFeed, nil
}
