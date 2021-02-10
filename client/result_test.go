/* Copyright 2021 Kilobit Labs Inc. */

package client_test

import _ "fmt"
import _ "errors"

import _ "context"
import _ "io"
import _ "os"
import _ "net/url"
import _ "net/http"
import _ "net/http/httptest"
import "encoding/json"

import "kilobit.ca/go/podcastindex/client"

import "kilobit.ca/go/tested/assert"
import "testing"

func TestResultTest(t *testing.T) {
	assert.Expect(t, true, true, "Failed Sanity Check.")
}

// From the developer documentation on podcastindex.org.
var testBody2 = `{
    "status": "true",
    "feeds": [
        {
            "id": 75075,
            "title": "Batman University",
            "url": "https:\/\/feeds.theincomparable.com\/batmanuniversity",
            "originalUrl": "https:\/\/feeds.theincomparable.com\/batmanuniversity",
            "link": "https:\/\/www.theincomparable.com\/batmanuniversity\/",
            "description": "Batman University is a seasonal podcast about you know who. It began with an analysis of episodes of \u201cBatman: The Animated Series\u201d but has now expanded to cover other series, movies, and media. Your professor is Tony Sindelar.",
            "author": "Tony Sindelar",
            "ownerName": "The Incomparable",
            "image": "https:\/\/www.theincomparable.com\/imgs\/logos\/logo-batmanuniversity-3x.jpg?cache-buster=2019-06-11",
            "artwork": "https:\/\/www.theincomparable.com\/imgs\/logos\/logo-batmanuniversity-3x.jpg?cache-buster=2019-06-11",
            "lastUpdateTime": 1610854330,
            "lastCrawlTime": 1612184257,
            "lastParseTime": 1610854342,
            "lastGoodHttpStatusTime": 1612184257,
            "lastHttpStatus": 200,
            "contentType": "application\/rss+xml",
            "itunesId": 1441923632,
            "generator": null,
            "language": "en-us",
            "type": 0,
            "dead": 0,
            "crawlErrors": 0,
            "parseErrors": 0,
            "categories": {
                "104": "Tv",
                "105": "Film",
                "107": "Reviews"
            },
            "locked": 0,
            "imageUrlHash": 1639321931
        }
    ],
    "count": 1,
    "query": "batman university",
    "description": "Found matching feeds."
}`

func TestResult(t *testing.T) {

	result := client.Result{}
	err := json.Unmarshal([]byte(testBody2), &result)
	assert.Ok(t, err, result)

	//t.Log(result.Keys())
	//t.Log(result.Description())

	query, ok := result.Get("query")
	assert.Expect(t, true, ok, query)
	assert.Expect(t, "batman university", query)
	//t.Log(query, ok)

	//fs, ok := result.Get("feeds")
	//t.Log(fs, ok)

	feeds := result.Feeds()
	assert.Expect(t, 1, len(feeds), feeds)
	//t.Log(feeds[0].Keys())

	authors := feeds[0].GetString("author")
	assert.Expect(t, "Tony Sindelar", authors)
	//t.Log(authors)
}
