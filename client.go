package gbfs

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var (
	ErrMissingAutodiscoveryURL = errors.New("missing auto discovery url")
	ErrFeedNotFound            = errors.New("feed not found")
	ErrFailedAutodiscoveryURL  = errors.New("failed to get auto discovery url")
	ErrInvalidLanguage         = errors.New("invalid language")
	ErrInvalidSubscribeHandler = errors.New("invalid subscribe handler")
)

type (
	// Client ...
	Client struct {
		httpClient *http.Client
		cache      Cache
		Options    *ClientOptions
	}
	// ClientOptions ...
	ClientOptions struct {
		AutoDiscoveryURL string
		DefaultLanguage  string
		UserAgent        string
		HTTPClient       *http.Client
		Cache            Cache
	}
	// ClientSubscribeOptions ...
	ClientSubscribeOptions struct {
		Languages []string
		FeedNames []string
		Handler   func(*Client, Feed, error)
	}
)

// NewClient ...
func NewClient(options ClientOptions) (*Client, error) {
	if options.AutoDiscoveryURL == "" {
		return nil, ErrMissingAutodiscoveryURL
	}
	c := &Client{
		httpClient: options.HTTPClient,
		cache:      options.Cache,
		Options:    &options,
	}
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	if c.cache == nil {
		c.cache = NewInMemoryCache()
	}
	return c, nil
}

func cacheGet(c *Client, feed Feed) (Feed, error) {
	if c.cache != nil {
		cacheKey := feed.Name()
		if feed.GetLanguage() != "" && feed.Name() != FeedNameGbfs {
			cacheKey = cacheKey + ":" + feed.GetLanguage()
		}
		feedCached, ok := c.cache.Get(cacheKey)
		if ok && !feedCached.Expired() {
			return feedCached, nil
		}
	}
	return nil, nil
}

func cacheSet(c *Client, feed Feed) {
	if c.cache != nil {
		cacheKey := feed.Name()
		if feed.GetLanguage() != "" && feed.Name() != FeedNameGbfs {
			cacheKey = cacheKey + ":" + feed.GetLanguage()
		}
		c.cache.Set(cacheKey, feed)
	}
}

// GetURL ...
func (c *Client) GetURL(url string, feed Feed) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	userAgent := c.Options.UserAgent
	if userAgent == "" {
		userAgent = "gbfs-client/1.0"
	}
	req.Header.Add("User-Agent", userAgent)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ErrFeedNotFound
		}
		return errors.New("invalid response status: " + strconv.Itoa(res.StatusCode))
	}
	err = json.NewDecoder(res.Body).Decode(feed)
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (c *Client) Get(feed Feed) error {
	cached, err := cacheGet(c, feed)
	if cached != nil {
		cloneValue(cached, feed)
		return nil
	}
	var gbfsFeed *FeedGbfs
	tmp, ok := c.cache.Get(FeedNameGbfs)
	if ok {
		gbfsFeed = tmp.(*FeedGbfs)
	}
	if !ok && feed.Name() != FeedNameGbfs {
		gbfsFeed = &FeedGbfs{}
		err = c.Get(gbfsFeed)
		if err != nil {
			return ErrFailedAutodiscoveryURL
		}
	}
	language := feed.GetLanguage()
	if language == "" {
		language = c.Options.DefaultLanguage
	}
	if language == "" {
		return NewError(feed.Name()+": ", ErrInvalidLanguage)
	}
	var url string
	if feed.Name() != FeedNameGbfs {
		l, ok := gbfsFeed.Data[language]
		if !ok {
			l, ok = gbfsFeed.Data[c.Options.DefaultLanguage]
			if !ok {
				return NewError(feed.Name()+": ", ErrInvalidLanguage)
			}
		}
		for _, f := range l.Feeds {
			if f.Name != nil && *f.Name == feed.Name() {
				url = *f.URL
				break
			}
		}
	} else {
		url = c.Options.AutoDiscoveryURL
	}
	if url == "" {
		return NewError(feed.Name()+": ", ErrFeedNotFound)
	}
	err = c.GetURL(url, feed)
	if err != nil {
		return NewError(feed.Name()+": ", err)
	}
	cacheSet(c, feed)
	return nil
}

func cloneValue(src, dst interface{}) {
	x := reflect.ValueOf(src)
	if x.Kind() == reflect.Ptr {
		x2 := x.Elem()
		y := reflect.New(x2.Type())
		y2 := y.Elem()
		y2.Set(x2)
		reflect.ValueOf(dst).Elem().Set(y.Elem())
	} else {
		dst = x.Interface()
	}
}

// Subscribe ...
func (c *Client) Subscribe(options ClientSubscribeOptions) error {
	if options.Handler == nil {
		return ErrInvalidSubscribeHandler
	}
	channel := make(chan interface{})
	go (func() {
		loops := []Feed{}
		g := &FeedGbfs{}
		err := c.Get(g)
		if err != nil {
			channel <- errors.New(g.Name() + ": " + err.Error())
			return
		}
		channel <- g
		if options.FeedNames == nil || inSlice(g.Name(), options.FeedNames) {
			loops = append(loops, g)
		}
		for language, languageData := range g.Data {
			if options.Languages != nil && !inSlice(language, options.Languages) {
				continue
			}
			for _, feed := range languageData.Feeds {
				if feed.Name == nil || options.FeedNames != nil && !inSlice(*feed.Name, options.FeedNames) {
					continue
				}
				f := FeedStruct(*feed.Name)
				if f == nil {
					continue
				}
				f.SetLanguage(language)
				err = c.Get(f)
				if err != nil {
					f.SetTTL(g.GetTTL())
					loops = append(loops, f)
					channel <- errors.New(*feed.Name + ": " + err.Error())
					continue
				}
				loops = append(loops, f)
				channel <- f
			}
		}
		for _, loop := range loops {
			go (func(loop Feed) {
				for {
					time.Sleep(time.Duration(loop.GetTTL()) * time.Second)
					f := FeedStruct(loop.Name())
					f.SetLanguage(loop.GetLanguage())
					err = c.Get(f)
					if err != nil {
						channel <- errors.New(loop.Name() + ": " + err.Error())
						continue
					}
					if loop.GetTTL() == 0 {
						break
					}
					if f.Expired() {
						continue
					}
					channel <- f
				}
			})(loop)
		}
	})()
	for {
		msg := <-channel
		switch v := msg.(type) {
		case Feed:
			options.Handler(c, v, nil)
		case error:
			options.Handler(c, nil, v)
		default:
			options.Handler(c, nil, errors.New("channel: unknown type"))
		}
	}
	return nil
}
