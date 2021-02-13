/* Copyright 2021 Kilobit Labs Inc. */

package client_test

import _ "fmt"
import _ "errors"
import "time"

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

// From the developer documentation on podcastindex.org.
var testBodyWithFeed = `{
    "status": "true",
    "query": {
        "url": "https:\/\/feeds.theincomparable.com\/batmanuniversity"
    },
    "feed": {
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
        "lastCrawlTime": 1612844700,
        "lastParseTime": 1610854342,
        "lastGoodHttpStatusTime": 1612844700,
        "lastHttpStatus": 200,
        "contentType": "application\/rss+xml",
        "itunesId": 1441923632,
        "generator": null,
        "language": "en-us",
        "type": 0,
        "dead": 0,
        "chash": "ad651c60eaaf3344595c0dd0bd787993",
        "episodeCount": 19,
        "crawlErrors": 0,
        "parseErrors": 0,
        "categories": {
            "104": "Tv",
            "105": "Film",
            "107": "Reviews"
        },
        "locked": 0,
        "imageUrlHash": 1639321931
    },
    "description": "Found matching feed."
}`

func TestResultFeed(t *testing.T) {

	result := client.Result{}
	err := json.Unmarshal([]byte(testBodyWithFeed), &result)
	assert.Ok(t, err, result)

	//t.Log(result.Keys())
	//t.Log(result.Description())

	//fs, ok := result.Get("feeds")
	//t.Log(fs, ok)

	feed := result.Feed()
	assert.Expect(t, false, feed == nil, result)
	//t.Log(feed.Keys())

	assert.Expect(t, "Tony Sindelar", feed.Author())
	assert.Expect(t, "Batman University", feed.Title())
	//t.Log(authors)
}

var testBodyWithItems string = `{
    "status": "true",
    "items": [
        {
            "id": 1670933447,
            "title": "-RESUMEN CONFERENCIA MA\u00d1ANERA 12 DE FEBRERO",
            "link": "https:\/\/anchor.fm\/bibiu00e1n-reyes\/episodes\/-RESUMEN-CONFERENCIA-MAANERA-12-DE-FEBRERO-eqb2ra",
            "description": "<p>Informaci\u00f3n Nacional<\/p>",
            "guid": "6299506d-52e8-4ebe-9f11-c157d1609210",
            "datePublished": 1613165754,
            "datePublishedPretty": "February 12, 2021 3:35pm",
            "dateCrawled": 1613174300,
            "enclosureUrl": "https:\/\/anchor.fm\/s\/22d7a678\/podcast\/play\/26626346\/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-1-12%2F154715110-44100-2-d12ecbdab2996.m4a",
            "enclosureType": "audio\/x-m4a",
            "enclosureLength": 1041768,
            "explicit": 0,
            "episode": null,
            "episodeType": "full",
            "season": 0,
            "image": "https:\/\/d3t3ozftmdmh3i.cloudfront.net\/production\/podcast_uploaded_episode\/5745582\/5745582-1613165762229-69926b06db573.jpg",
            "feedItunesId": 1516125350,
            "feedImage": "https:\/\/d3t3ozftmdmh3i.cloudfront.net\/production\/podcast_uploaded_nologo\/5745582\/5745582-1590358064595-b1e1fe0ca6175.jpg",
            "feedId": 291242,
            "feedTitle": "Luces del Siglo",
            "feedLanguage": "es"
        },
        {
            "id": 1670932947,
            "title": "Defense & Aerospace Podcast [Washington Roundtable Feb 12, 21]",
            "link": "https:\/\/soundcloud.com\/defaeroreport\/defense-aerospace-podcast-washington-roundtable-feb-12-21",
            "description": "On this Washington Roundtable episode of the Defense & Aerospace Report Podcast, sponsored by Bell, our guests in segment one are Michael Bayer, former chairman of the Defense Business Board and the president of the Dumbarton Strategies consultancy, and Arnold Punaro, the chairman of the National Defense Industrial Association and CEO of the Punaro Group consultancy to discuss the ongoing presidential transition and how to best accomplish the monumental task of a peaceful transition of power.\n\nIn segment two, our roundtable guests are Dov Zakheim, PhD, former DoD comptroller, now with the Center for Strategic and International Studies, Dr Gordon Adams...",
            "guid": "tag:soundcloud,2010:tracks\/984402556",
            "datePublished": 1613164626,
            "datePublishedPretty": "February 12, 2021 3:17pm",
            "dateCrawled": 1613174300,
            "enclosureUrl": "http:\/\/feeds.soundcloud.com\/stream\/984402556-defaeroreport-defense-aerospace-podcast-washington-roundtable-feb-12-21.mp3",
            "enclosureType": "audio\/mpeg",
            "enclosureLength": 78411148,
            "explicit": 0,
            "episode": null,
            "episodeType": null,
            "season": 0,
            "image": "http:\/\/i1.sndcdn.com\/avatars-000494055963-8sk2v4-original.jpg",
            "feedItunesId": 1228868129,
            "feedImage": "http:\/\/i1.sndcdn.com\/avatars-000494055963-8sk2v4-original.jpg",
            "feedId": 809504,
            "feedTitle": "Defense & Aerospace Report",
            "feedLanguage": "en"
        },
        {
            "id": 1670932507,
            "title": "LifeMessage - Menerima Perbedaan",
            "link": "https:\/\/anchor.fm\/lifehouse-jakarta\/episodes\/LifeMessage---Menerima-Perbedaan-epos53",
            "description": "Ps. Andreas Agus - Menerima Perbedaan",
            "guid": "9ca2a8c1-a001-4aff-913a-cff6302e90ce",
            "datePublished": 1613163600,
            "datePublishedPretty": "February 12, 2021 3:00pm",
            "dateCrawled": 1613174299,
            "enclosureUrl": "https:\/\/anchor.fm\/s\/16bd4104\/podcast\/play\/26029667\/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-1-1%2F150830660-44100-2-b0a95ff741124.m4a",
            "enclosureType": "audio\/x-m4a",
            "enclosureLength": 6193529,
            "explicit": 0,
            "episode": null,
            "episodeType": "full",
            "season": 0,
            "image": "https:\/\/d3t3ozftmdmh3i.cloudfront.net\/production\/podcast_uploaded_episode\/3715017\/3715017-1612156942461-8b9f0c8ce73ee.jpg",
            "feedItunesId": 1503273955,
            "feedImage": "https:\/\/d3t3ozftmdmh3i.cloudfront.net\/production\/podcast_uploaded\/3715017\/3715017-1584437013771-e35d8c3ef8fd6.jpg",
            "feedId": 720963,
            "feedTitle": "Lifehouse Jakarta",
            "feedLanguage": "in"
        },
        {
            "id": 1670932125,
            "title": "En La Jugada - Febrero 12 de 2021",
            "link": "https:\/\/www.spreaker.com\/user\/rcnradiocolombia\/en-la-jugada-febrero-12-de-2021",
            "description": "En la Jugada conversamos sobre la victoria del DIM sobre el Tolima en la Copa Betplay, la llegada de  Diber Cambindo al Am\u00e9rica de Cali, la victoria de Novak Djokovic en el Abierto de Australia y mucho m\u00e1s.",
            "guid": "https:\/\/api.spreaker.com\/episode\/43436872",
            "datePublished": 1613166467,
            "datePublishedPretty": "February 12, 2021 3:47pm",
            "dateCrawled": 1613174299,
            "enclosureUrl": "https:\/\/api.spreaker.com\/download\/episode\/43436872\/prog_en_la_jugada_feb_12.mp3",
            "enclosureType": "audio\/mpeg",
            "enclosureLength": 21294398,
            "explicit": 0,
            "episode": null,
            "episodeType": "full",
            "season": 0,
            "image": "https:\/\/d3wo5wojvuv7l.cloudfront.net\/t_rss_itunes_square_1400\/images.spreaker.com\/original\/46ba84a1d2346d4c9d51d6ecfc946644.jpg",
            "feedItunesId": 1129843813,
            "feedImage": "https:\/\/d3wo5wojvuv7l.cloudfront.net\/t_rss_itunes_square_1400\/images.spreaker.com\/original\/46ba84a1d2346d4c9d51d6ecfc946644.jpg",
            "feedId": 51782,
            "feedTitle": "En La Jugada RCN",
            "feedLanguage": "es"
        }
    ],
    "count": 4,
    "max": "7",
    "description": "Found matching items."
}`

func TestResultItems(t *testing.T) {

	result := client.Result{}
	err := json.Unmarshal([]byte(testBodyWithItems), &result)
	assert.Ok(t, err, result)

	items := result.Items()
	assert.Expect(t, result.Count(), len(items), result)

	item := items[3]

	assert.Expect(t, "En La Jugada - Febrero 12 de 2021", item.Title(), item)
	assert.Expect(t, time.Unix(1613166467, 0), item.DatePublished(), item)
	assert.Expect(t, "audio/mpeg", item.EnclosureType(), item)
	assert.Expect(t, -1, item.Episode(), item)
	assert.Expect(t, 0, item.Season(), item)
	assert.Expect(t, "En La Jugada RCN", item.FeedTitle(), item)
	assert.Expect(t, "es", item.FeedLanguage(), item)
}
