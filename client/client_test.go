/* Copyright 2021 Kilobit Labs Inc. */

package client_test

import _ "fmt"
import _ "errors"

import "context"
import "io"
import "time"
import "os"
import "net/url"
import "net/http"
import "net/http/httptest"

import "kilobit.ca/go/podcastindex/client"

import "kilobit.ca/go/tested/assert"
import "testing"

func TestClientTest(t *testing.T) {
	assert.Expect(t, true, true, "Failed Sanity Check.")
}

func TestClientSearchLive(t *testing.T) {

	t.Skip("This test requires remote resources and should normally be skipped.")

	key := os.Getenv("PODCAST_INDEX_API_KEY")
	secret := os.Getenv("PODCAST_INDEX_API_SECRET")

	if key == "" || secret == "" {
		t.Skip("Missing API Key and/or Secret. Skipping.")
	}

	ctx := context.TODO()
	ctx = context.WithValue(ctx, client.PICAPIKey, key)
	ctx = context.WithValue(ctx, client.PICAPISecret, secret)

	tests := []struct {
		query string
	}{
		{"batman university"},
	}

	pic := client.New(ctx)

	for _, test := range tests {

		resp, err := pic.Search(context.TODO(), test.query)
		assert.Ok(t, err, resp)

		t.Logf("%#v", resp)
	}
}

// From the developer documentation on podcastindex.org.
var testBody1 = `{
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

func TestClientSearchMock(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {

			assert.Expect(t, req.URL.Path, "/api/1.0/search/byterm", req.URL)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, testBody1)
		}),
	)
	defer srv.Close()

	hc := srv.Client()

	u, err := url.Parse(srv.URL)
	assert.Ok(t, err, u)

	pic := client.New(
		context.TODO(),
		client.OptionClient(hc),
		client.OptionScheme(u.Scheme),
		client.OptionHost(u.Host),
	)

	result, err := pic.Search(context.TODO(), "foo")
	assert.Ok(t, err, result, pic)

	//t.Log(result)

	feeds := result.Feeds()
	assert.Expect(t, 1, len(feeds), feeds)

	feed := feeds[0]

	assert.Expect(t, 75075, feed.ID(), feed)
	assert.Expect(t, "Batman University", feed.Title(), feed)
	assert.Expect(t, "Tony Sindelar", feed.Author(), feed)
	assert.Expect(t, time.Unix(1610854330, 0), feed.LastUpdated(), feed)
	assert.Expect(t, false, feed.Locked(), feed)
	assert.Expect(t, false, feed.Dead(), feed)
}
