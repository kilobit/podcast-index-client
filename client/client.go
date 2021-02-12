/* Copyright 2021 Kilobit Labs Inc. */

// A client for using the Podcast Index API in Golang.
//
package client

import "fmt"
import _ "errors"

import "context"
import "crypto/sha1"
import "io"
import "io/ioutil"
import "encoding/json"
import "path"
import "strconv"
import "time"
import "net/url"
import "net/http"

const DefaultTimeout time.Duration = 10 * time.Second
const DefaultUserAgent string = "Podcast Gateway/0.0.1"
const DefaultScheme string = "https"
const DefaultHost string = "api.podcastindex.org"
const DefaultRoot string = "api"

const APIEntityAdd = "add"
const APIEntityEpisodes = "episodes"
const APIEntityPodcasts = "podcasts"
const APIEntityRecent = "recent"
const APIEntitySearch = "search"

const APISelectBatchByFeedURL = "batch/byfeedurl"
const APISelectByID = "byid"
const APISelectByItunesID = "byitunesid"
const APISelectByFeedURL = "byfeedurl"
const APISelectByFeedID = "byfeedid"
const APISelectRandom = "random"
const APISelectByTerm = "byterm"
const APISelectFeeds = "feeds"
const APISelectEpisodes = "episodes"
const APISelectNewFeeds = "newfeeds"
const APISelectSoundbites = "soundbites"

const API_VERSION = "1.0"

type cKey int

const (
	PICAPIKey cKey = iota
	PICAPISecret
)

type PodcastIndexClient struct {
	client  *http.Client
	ua      string
	scheme  string
	host    string
	root    string
	version string
	key     string
	secret  string
	ctx     context.Context
}

type Option func(*PodcastIndexClient)

func OptionClient(client *http.Client) Option {
	return Option(func(pic *PodcastIndexClient) { pic.client = client })
}

func OptionScheme(scheme string) Option {
	return Option(func(pic *PodcastIndexClient) { pic.scheme = scheme })
}

func OptionHost(host string) Option {
	return Option(func(pic *PodcastIndexClient) { pic.host = host })
}

func OptionRoot(root string) Option {
	return Option(func(pic *PodcastIndexClient) { pic.root = root })
}

func New(ctx context.Context, opts ...Option) *PodcastIndexClient {

	if ctx == nil {
		ctx = context.TODO()
	}

	client := &http.Client{
		Timeout: DefaultTimeout,
	}

	key := stringFromCtx(ctx, PICAPIKey)
	secret := stringFromCtx(ctx, PICAPISecret)

	pic := PodcastIndexClient{
		client,
		DefaultUserAgent,
		DefaultScheme,
		DefaultHost,
		DefaultRoot,
		API_VERSION,
		key,
		secret,
		ctx,
	}

	pic.Options(opts...)

	return &pic
}

func (pic *PodcastIndexClient) Options(opts ...Option) {
	for _, opt := range opts {
		opt(pic)
	}
}

func (pic *PodcastIndexClient) SetAPIKeys(key, secret string) {
	pic.key = key
	pic.secret = secret
}

func (pic *PodcastIndexClient) Search(ctx context.Context, query string) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntitySearch,
		APISelectByTerm,
		NewParam("q", query),
	)
}

func (pic *PodcastIndexClient) PodcastsByFeedURL(ctx context.Context, url string) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityPodcasts,
		APISelectByFeedURL,
		NewParam("url", url),
	)
}

func (pic *PodcastIndexClient) PodcastsByFeedID(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityPodcasts,
		APISelectByFeedID,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) PodcastsByItunesID(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityPodcasts,
		APISelectByItunesID,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) EpisodesByFeedID(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityEpisodes,
		APISelectByFeedID,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) EpisodesByFeedURL(ctx context.Context, url string) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityEpisodes,
		APISelectByFeedURL,
		NewParam("url", url),
	)
}

func (pic *PodcastIndexClient) EpisodesByItunesID(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityEpisodes,
		APISelectByItunesID,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) EpisodesByID(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityEpisodes,
		APISelectByID,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) EpisodesRandom(ctx context.Context, id int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityEpisodes,
		APISelectRandom,
		NewParam("id", strconv.Itoa(id)),
	)
}

func (pic *PodcastIndexClient) RecentEpisodes(ctx context.Context, max int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityRecent,
		APISelectEpisodes,
		NewParam("max", strconv.Itoa(max)),
	)
}

func (pic *PodcastIndexClient) RecentFeeds(ctx context.Context, max int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityRecent,
		APISelectFeeds,
		NewParam("max", strconv.Itoa(max)),
	)
}

func (pic *PodcastIndexClient) RecentNewFeeds(ctx context.Context, max int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityRecent,
		APISelectNewFeeds,
	)
}

func (pic *PodcastIndexClient) RecentSoundBites(ctx context.Context, max int) (Result, error) {

	return pic.CallReadAPI(
		ctx,
		http.MethodGet,
		APIEntityRecent,
		APISelectSoundbites,
	)
}

type Param struct {
	key   string
	value string
}

func NewParam(key, value string) *Param {
	return &Param{key, value}
}

func (pic *PodcastIndexClient) CallReadAPI(ctx context.Context, method, entity, selector string, params ...*Param) (Result, error) {

	path := path.Join(pic.root, pic.version, entity, selector)

	u := url.URL{Scheme: pic.scheme, Host: pic.host, Path: path}
	q := u.Query()
	for _, param := range params {
		q.Set(param.key, param.value)
	}
	u.RawQuery = q.Encode()

	result := Result{}

	resp, err := pic.request(method, u.String(), nil)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("Bad response code from the upstream API, %s.", resp.Status)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return result, fmt.Errorf("Badly formatted response, expected application/json.")
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("Error while reading the response body, %v.", err)
	}

	err = json.Unmarshal(bs, &result)
	if err != nil {
		return result, fmt.Errorf("Error decoding response JSON, %v", err)
	}

	return result, nil
}

func (pic *PodcastIndexClient) request(method, url string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	req = req.WithContext(ctx)

	now := strconv.FormatInt(time.Now().Unix(), 10)
	auth := calculateAuth(pic.key, pic.secret, now)

	req.Header.Set("User-Agent", pic.ua)
	req.Header.Set("X-Auth-Date", now)
	req.Header.Set("X-Auth-Key", pic.key)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/json")

	return pic.client.Do(req)
}

func calculateAuth(key, secret, datestr string) string {

	h := sha1.New()
	io.WriteString(h, key)
	io.WriteString(h, secret)
	io.WriteString(h, datestr)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func stringFromCtx(ctx context.Context, key interface{}) string {
	str, ok := ctx.Value(key).(string)
	if !ok {
		str = ""
	}

	return str
}
