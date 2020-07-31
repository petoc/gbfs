package gbfs

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"sync"
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
		cache      *ClientCache
		Options    *ClientOptions
	}
	// ClientOptions ...
	ClientOptions struct {
		AutoDiscoveryURL string
		DefaultLanguage  string
		UserAgent        string
		HTTPClient       *http.Client
	}
	// ClientCache ...
	ClientCache struct {
		sync.RWMutex
		feedGbfs *FeedGbfs
		feeds    map[string]Feed
	}
	// ClientSubscribeOptions ...
	ClientSubscribeOptions struct {
		Languages []string
		FeedNames []string
		Handler   func(*Client, Feed, error)
	}
)

// NewClient ...
func NewClient(options *ClientOptions) (*Client, error) {
	if options.AutoDiscoveryURL == "" {
		return nil, ErrMissingAutodiscoveryURL
	}
	c := &Client{
		httpClient: options.HTTPClient,
		cache: &ClientCache{
			feeds: make(map[string]Feed),
		},
		Options: options,
	}
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	return c, nil
}

func cacheGet(c *Client, feed Feed) (Feed, error) {
	if c.cache != nil {
		if feed.Name() == FeedNameGbfs {
			if c.cache.feedGbfs != nil && !c.cache.feedGbfs.Expired() {
				return c.cache.feedGbfs, nil
			}
		} else {
			cacheKey := string(feed.Name())
			if feed.GetLanguage() != "" {
				cacheKey = cacheKey + ":" + feed.GetLanguage()
			}
			c.cache.RLock()
			feedCached, ok := c.cache.feeds[cacheKey]
			c.cache.RUnlock()
			if ok && !feedCached.Expired() {
				return feedCached, nil
			}
		}
	}
	return nil, nil
}

func cacheSet(c *Client, feed Feed) {
	if c.cache != nil {
		if feed.Name() == FeedNameGbfs {
			c.cache.feedGbfs, _ = feed.(*FeedGbfs)
		} else {
			cacheKey := string(feed.Name())
			if feed.GetLanguage() != "" {
				cacheKey = cacheKey + ":" + feed.GetLanguage()
			}
			c.cache.Lock()
			c.cache.feeds[cacheKey] = feed
			c.cache.Unlock()
		}
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
		return errors.New("invalid response status")
	}
	err = json.NewDecoder(res.Body).Decode(feed)
	if err != nil {
		return errors.New("invalid json")
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
	if c.cache.feedGbfs == nil && feed.Name() != FeedNameGbfs {
		g := &FeedGbfs{}
		err = c.Get(g)
		if err != nil {
			return ErrFailedAutodiscoveryURL
		}
	}
	language := feed.GetLanguage()
	if language == "" {
		language = c.Options.DefaultLanguage
	}
	url := c.Options.AutoDiscoveryURL
	if feed.Name() != FeedNameGbfs {
		l, ok := c.cache.feedGbfs.Data[language]
		if !ok {
			l, ok = c.cache.feedGbfs.Data[c.Options.DefaultLanguage]
			if !ok {
				return ErrInvalidLanguage
			}
		}
		for _, f := range l.Feeds {
			if f.Name == feed.Name() {
				url = f.URL
				break
			}
		}
	}
	err = c.GetURL(url, feed)
	if err != nil {
		return err
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
func (c *Client) Subscribe(options *ClientSubscribeOptions) error {
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
		if options.FeedNames == nil || indexOfStr(g.Name(), options.FeedNames) > -1 {
			loops = append(loops, g)
		}
		for language, languageData := range g.Data {
			if options.Languages != nil {
				match := false
				for _, languageSub := range options.Languages {
					if language == languageSub {
						match = true
						break
					}
				}
				if !match {
					continue
				}
			}
			for _, feed := range languageData.Feeds {
				if options.FeedNames != nil && indexOfStr(feed.Name, options.FeedNames) == -1 {
					continue
				}
				f := FeedStruct(feed.Name)
				f.SetLanguage(language)
				err = c.Get(f)
				if err != nil {
					f.SetTTL(g.GetTTL())
					loops = append(loops, f)
					channel <- errors.New(feed.Name + ": " + err.Error())
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
