package gbfs

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	ErrMissingSystemID      = errors.New("missing system id")
	ErrMissingRootDir       = errors.New("missing root directory")
	ErrMissingBaseURL       = errors.New("missing base url")
	ErrInvalidDefaultTTL    = errors.New("invalid default ttl")
	ErrMissingFeedHandlers  = errors.New("missing feed handlers")
	ErrMissingServerAddress = errors.New("missing server address")
)

type (
	Server struct {
		Options *ServerOptions
	}
	ServerOptions struct {
		SystemID      string
		RootDir       string
		BaseURL       string
		BasePath      string
		Version       string
		DefaultTTL    int
		FeedHandlers  []*FeedHandler
		UpdateHandler func(server *Server, feed Feed, path string, err error)
	}
	FeedHandler struct {
		TTL     int
		Path    string
		Handler func(*Server) ([]Feed, error)
	}
)

func WriteFeed(filePath string, feed Feed) error {
	b, err := json.Marshal(feed)
	if err != nil {
		return err
	}
	filePath = filepath.FromSlash(filePath)
	fileDir := filepath.Dir(filePath)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(filePath, b, 0644)
}

func NewServer(options ServerOptions) (*Server, error) {
	if options.SystemID == "" {
		return nil, ErrMissingSystemID
	}
	if options.RootDir == "" {
		return nil, ErrMissingRootDir
	}
	if options.BaseURL == "" {
		return nil, ErrMissingBaseURL
	}
	if options.DefaultTTL <= 0 {
		return nil, ErrInvalidDefaultTTL
	}
	s := &Server{
		Options: &options,
	}
	return s, nil
}

func (s *Server) Start() error {
	if len(s.Options.FeedHandlers) == 0 {
		return ErrMissingFeedHandlers
	}
	gbfsGenerated := false
	gbfsFeed := &FeedGbfs{}
	gbfsFeed.SetTTL(s.Options.DefaultTTL)
	gbfsFeed.SetVersion(s.Options.Version)
	var wgGbfsFeed sync.WaitGroup
	wgGbfsFeed.Add(len(s.Options.FeedHandlers))
	for _, feedHandler := range s.Options.FeedHandlers {
		go (func(feedHandler *FeedHandler) {
			if feedHandler.TTL == 0 {
				feedHandler.TTL = s.Options.DefaultTTL
			}
			for {
				feeds, err := feedHandler.Handler(s)
				if err != nil {
					s.Options.UpdateHandler(s, nil, "", err)
					if !gbfsGenerated {
						wgGbfsFeed.Done()
					}
					break
				}
				for _, feed := range feeds {
					feed.SetLastUpdated(Timestamp(time.Now().Format(time.RFC3339)))
					if feed.GetTTL() == 0 {
						feed.SetTTL(feedHandler.TTL)
					}
					feed.SetVersion(s.Options.Version)
					pathSegments := []string{}
					if s.Options.BasePath != "" {
						pathSegments = append(pathSegments, strings.Trim(s.Options.BasePath, "/"))
					}
					path := strings.Trim(feedHandler.Path, "/")
					if path == "" {
						path = feed.Name() + ".json"
					}
					pathSegments = append(pathSegments, path)
					filePath := strings.Join(append([]string{s.Options.RootDir}, pathSegments...), "/")
					err := WriteFeed(filePath, feed)
					s.Options.UpdateHandler(s, feed, strings.Join(pathSegments, "/"), err)
					if err != nil {
						continue
					}
					if !gbfsGenerated && feed.Name() != FeedNameGbfs {
						gbfsFeed.Lock()
						if gbfsFeed.Data == nil {
							gbfsFeed.Data = &FeedGbfsData{
								Feeds: []*FeedGbfsFeed{},
							}
						}
						if gbfsFeed.Data.Feeds != nil {
							gbfsFeed.Data.Feeds = append(gbfsFeed.Data.Feeds, &FeedGbfsFeed{
								Name: NewString(feed.Name()),
								URL:  NewString(strings.Join(append([]string{strings.Trim(s.Options.BaseURL, "/")}, pathSegments...), "/")),
							})
						}
						gbfsFeed.Unlock()
					}
				}
				if !gbfsGenerated {
					wgGbfsFeed.Done()
				}
				if feedHandler.TTL == 0 {
					break
				}
				time.Sleep(time.Duration(feedHandler.TTL) * time.Second)
			}
		})(feedHandler)
	}
	wgGbfsFeed.Wait()
	gbfsGenerated = true
	feedNames := FeedNameAll()
	if gbfsFeed.Data.Feeds != nil {
		sort.Slice(gbfsFeed.Data.Feeds, func(i, j int) bool {
			if gbfsFeed.Data.Feeds[i].Name == nil || gbfsFeed.Data.Feeds[j].Name == nil {
				return false
			}
			if IndexInSlice(*gbfsFeed.Data.Feeds[i].Name, feedNames) > IndexInSlice(*gbfsFeed.Data.Feeds[j].Name, feedNames) {
				return false
			}
			return true
		})
	}
	for {
		gbfsFeed.SetLastUpdated(Timestamp(time.Now().Format(time.RFC3339)))
		pathSegments := []string{}
		if s.Options.BasePath != "" {
			pathSegments = append(pathSegments, strings.Trim(s.Options.BasePath, "/"))
		}
		pathSegments = append(pathSegments, gbfsFeed.Name()+".json")
		filePath := strings.Join(append([]string{s.Options.RootDir}, pathSegments...), "/")
		err := WriteFeed(filePath, gbfsFeed)
		s.Options.UpdateHandler(s, gbfsFeed, strings.Join(pathSegments, "/"), err)
		if gbfsFeed.GetTTL() == 0 {
			break
		}
		time.Sleep(time.Duration(gbfsFeed.GetTTL()) * time.Second)
	}
	return nil
}

func NewFileServer(addr, rootDir string) (*http.Server, error) {
	if addr == "" {
		return nil, ErrMissingServerAddress
	}
	if rootDir == "" {
		return nil, ErrMissingRootDir
	}
	s := &http.Server{
		Addr:         addr,
		Handler:      http.FileServer(http.Dir(rootDir)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return s, nil
}
