/* Copyright 2021 Kilobit Labs Inc. */

package client

import _ "fmt"
import _ "errors"
import "time"

import . "kilobit.ca/go/objected"

// A feed encapsulates what the Podcast Index knows about a particular
// podcast.
//
type Feed struct {
	Object
}

func (f *Feed) ID() int {
	n, err := f.GetNumber("id")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (f *Feed) Title() string {
	return f.GetString("title")
}

func (f *Feed) URL() string {
	return f.GetString("url")
}

func (f *Feed) OriginalURL() string {
	return f.GetString("originalUrl")
}

func (f *Feed) Link() string {
	return f.GetString("link")
}

func (f *Feed) Description() string {
	return f.GetString("description")
}

func (f *Feed) Author() string {
	return f.GetString("author")
}

func (f *Feed) ownerName() string {
	return f.GetString("ownerName")
}

func (f *Feed) ImageURL() string {
	return f.GetString("image")
}

func (f *Feed) ArtworkURL() string {
	return f.GetString("artwork")
}

func (f *Feed) LastUpdated() time.Time {

	result := time.Time{}

	n, err := f.GetNumber("lastUpdateTime")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (f *Feed) LastCrawled() time.Time {

	result := time.Time{}

	n, err := f.GetNumber("lastCrawlTime")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (f *Feed) LastParsed() time.Time {

	result := time.Time{}

	n, err := f.GetNumber("lastParseTime")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (f *Feed) LastGoodHTTPStatusTime() time.Time {

	result := time.Time{}

	n, err := f.GetNumber("lastGoodHttpStatusTime")
	if err == nil {
		result = time.Unix(int64(n), 0)
	}

	return result
}

func (f *Feed) LastHTTPStatus() int {
	n, err := f.GetNumber("lastHttpStatus")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (f *Feed) ContentType() string {
	return f.GetString("contentType")
}

func (f *Feed) ItunesID() int {
	n, err := f.GetNumber("itunesId")
	if err != nil {
		n = float64(-1)
	}

	return int(n)
}

func (f *Feed) Language() string {
	return f.GetString("language")
}

func (f *Feed) Locked() bool {
	n, err := f.GetNumber("locked")
	if err != nil {
		n = float64(-1)
	}

	return n != 0
}

func (f *Feed) Dead() bool {
	n, err := f.GetNumber("dead")
	if err != nil {
		n = float64(-1)
	}

	return n != 0
}
