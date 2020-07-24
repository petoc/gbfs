package gbfs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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
	// Server ...
	Server struct {
		Options *ServerOptions
		Logger  *log.Logger
	}
	// ServerOptions ...
	ServerOptions struct {
		SystemID     string
		RootDir      string
		BaseURL      string
		BasePath     string
		Version      string
		DefaultTTL   int
		FeedHandlers []*FeedHandler
		Logger       *log.Logger
	}
	// FeedHandler ...
	FeedHandler struct {
		Language string
		TTL      int
		Path     string
		Handler  func(s *Server) ([]Feed, error)
	}
	// FileServer ...
	FileServer struct {
		httpServer *http.Server
		Options    *FileServerOptions
		Logger     *log.Logger
	}
	// FileServerOptions ...
	FileServerOptions struct {
		Addr    string
		RootDir string
		Logger  *log.Logger
	}
)

func writeFeed(filePath string, feed Feed) error {
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
	ioutil.WriteFile(filePath, b, 0644)
	return nil
}

// NewServer ...
func NewServer(options *ServerOptions) (*Server, error) {
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
		Options: options,
		Logger:  options.Logger,
	}
	if s.Logger == nil {
		s.Logger = log.New(os.Stdout, "", 0)
	}
	return s, nil
}

// Start ...
func (s *Server) Start() error {
	if s.Options.FeedHandlers == nil || len(s.Options.FeedHandlers) == 0 {
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
					s.Logger.Println(err)
					if !gbfsGenerated {
						wgGbfsFeed.Done()
					}
					break
				}
				for _, feed := range feeds {
					feed.SetLastUpdated(Timestamp(time.Now().Unix()))
					if feed.GetTTL() == 0 {
						feed.SetTTL(feedHandler.TTL)
					}
					feed.SetVersion(s.Options.Version)
					pathSegments := []string{}
					if s.Options.BasePath != "" {
						pathSegments = append(pathSegments, strings.Trim(s.Options.BasePath, "/"))
					}
					if string(feed.GetLanguage()) != "" {
						pathSegments = append(pathSegments, string(feed.GetLanguage()))
					}
					path := strings.Trim(feedHandler.Path, "/")
					if path == "" {
						path = string(feed.Name()) + ".json"
					}
					pathSegments = append(pathSegments, path)
					filePath := strings.Join(append([]string{s.Options.RootDir}, pathSegments...), "/")
					err := writeFeed(filePath, feed)
					if err != nil {
						s.Logger.Println(err)
						continue
					}
					s.Logger.Printf("system=%s ttl=%d updated=%s", s.Options.SystemID, feed.GetTTL(), "/"+strings.Join(pathSegments, "/"))
					if !gbfsGenerated && feed.Name() != FeedNameGbfs {
						if gbfsFeed.Data == nil {
							gbfsFeed.Data = make(map[string]*FeedGbfsLanguage)
						}
						gbfsFeedLanguage, ok := gbfsFeed.Data[feed.GetLanguage()]
						if !ok {
							gbfsFeed.Data[feed.GetLanguage()] = &FeedGbfsLanguage{
								Feeds: []*FeedGbfsFeed{},
							}
						}
						gbfsFeedLanguage, ok = gbfsFeed.Data[feed.GetLanguage()]
						if ok {
							gbfsFeedLanguage.Feeds = append(gbfsFeedLanguage.Feeds, &FeedGbfsFeed{
								Name: feed.Name(),
								URL:  strings.Join(append([]string{strings.Trim(s.Options.BaseURL, "/")}, pathSegments...), "/"),
							})
						}
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
	go (func(gbfsFeed *FeedGbfs) {
		wgGbfsFeed.Wait()
		gbfsGenerated = true
		feedNames := FeedNameAll()
		for _, langData := range gbfsFeed.Data {
			sort.Slice(langData.Feeds, func(i, j int) bool {
				if indexOfStr(string(langData.Feeds[i].Name), feedNames) > indexOfStr(string(langData.Feeds[j].Name), feedNames) {
					return false
				}
				return true
			})
		}
		for {
			gbfsFeed.SetLastUpdated(Timestamp(time.Now().Unix()))
			pathSegments := []string{}
			if s.Options.BasePath != "" {
				pathSegments = append(pathSegments, strings.Trim(s.Options.BasePath, "/"))
			}
			pathSegments = append(pathSegments, string(gbfsFeed.Name())+".json")
			filePath := strings.Join(append([]string{s.Options.RootDir}, pathSegments...), "/")
			err := writeFeed(filePath, gbfsFeed)
			if err != nil {
				s.Logger.Println(err)
				continue
			}
			s.Logger.Printf("system=%s ttl=%d updated=%s", s.Options.SystemID, gbfsFeed.GetTTL(), "/"+strings.Join(pathSegments, "/"))
			if gbfsFeed.TTL == 0 {
				break
			}
			time.Sleep(time.Duration(gbfsFeed.GetTTL()) * time.Second)
		}
	})(gbfsFeed)
	return nil
}

// StartAndWait ...
func (s *Server) StartAndWait() error {
	err := s.Start()
	if err != nil {
		return err
	}
	w := make(chan struct{})
	<-w
	return nil
}

// NewFileServer ...
func NewFileServer(options *FileServerOptions) (*FileServer, error) {
	if options.Addr == "" {
		return nil, ErrMissingServerAddress
	}
	if options.RootDir == "" {
		return nil, ErrMissingRootDir
	}
	s := &FileServer{
		Options: options,
		Logger:  options.Logger,
	}
	if s.Logger == nil {
		s.Logger = log.New(os.Stdout, "", 0)
	}
	if s.httpServer == nil {
		s.httpServer = &http.Server{
			Addr:         options.Addr,
			Handler:      http.FileServer(http.Dir(options.RootDir)),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
	}
	return s, nil
}

// ListenAndServe ...
func (s *FileServer) ListenAndServe() error {
	s.Logger.Printf("listening=%s", s.Options.Addr)
	return s.httpServer.ListenAndServe()
}
