/* Copyright 2021 Kilobit Labs Inc. */

package client

import _ "fmt"
import _ "errors"
import "time"

import . "kilobit.ca/go/objected"

// An item encapsulates what the Podcast Index knows about a particular
// podcast episode.
//
type Item struct {
	Object
}

func (itm *Item) ID() int {
	n, err := itm.GetNumber("id")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) FeedID() int {
	n, err := itm.GetNumber("feedId")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) Title() string {
	return itm.GetString("title")
}

func (itm *Item) Link() string {
	return itm.GetString("link")
}

func (itm *Item) Description() string {
	return itm.GetString("description")
}

func (itm *Item) GUID() string {
	return itm.GetString("guid")
}

func (itm *Item) EnclosureURL() string {
	return itm.GetString("enclosureURL")
}

func (itm *Item) EnclosureType() string {
	return itm.GetString("enclosureType")
}

func (itm *Item) EnclosureLength() int {
	n, err := itm.GetNumber("enclosureLength")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) Episode() int {
	n, err := itm.GetNumber("episode")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) Season() int {
	n, err := itm.GetNumber("season")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) EpisodeType() string {
	return itm.GetString("episodeType")
}

func (itm *Item) Image() string {
	return itm.GetString("image")
}

func (itm *Item) DatePublished() time.Time {

	result := time.Time{}

	n, err := itm.GetNumber("datePublished")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (itm *Item) DatePublishedPretty() string {
	return itm.GetString("datePublishedPretty")
}

func (itm *Item) DateCrawled() time.Time {

	result := time.Time{}

	n, err := itm.GetNumber("dateCrawled")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (itm *Item) FeedItunesID() int {
	n, err := itm.GetNumber("feedItunesId")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (itm *Item) FeedTitle() string {
	return itm.GetString("feedTitle")
}

func (itm *Item) FeedLanguage() string {
	return itm.GetString("feedLanguage")
}
