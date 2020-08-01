# GBFS

Go implementation of client and server for GBFS (General Bikeshare Feed Specification).

## Usage

### Client

```go
c, err := gbfs.NewClient(gbfs.ClientOptions{
    AutoDiscoveryURL: "http://127.0.0.1:8080/v3/system_id/gbfs.json",
    DefaultLanguage:  "en",
})
if err != nil {
    log.Fatal(err)
}
```

#### Get single feed

```go
f := &gbfs.FeedSystemInformation{}
err := c.Get(f)
if err != nil {
    log.Println(err)
}
if f.Data != nil {
    log.Printf("feed=%s system_id=%s", f.Name(), f.Data.SystemID)
}
```

#### Subscribe

Client provides built-in function to handle feed updates.

```go
err := c.Subscribe(gbfs.ClientSubscribeOptions{
    // Languages: []string{"en"},
    // FeedNames: []string{gbfs.FeedNameStationInformation, gbfs.FeedNameFreeBikeStatus},
    Handler: func(c *gbfs.Client, f gbfs.Feed, err error) {
        if err != nil {
            log.Println(err)
            return
        }
        j, _ := json.Marshal(f)
        log.Printf("feed=%s data=%s", f.Name(), j)
    },
})
if err != nil {
    log.Println(err)
}
```

Subscription options `Languages` and `FeedNames` restrict subscription only to selected languages and feeds.

### Server

```go
s, err := gbfs.NewServer(gbfs.ServerOptions{
    SystemID:     systemID,
    RootDir:      "public",
    BaseURL:      "http://127.0.0.1:8080",
    BasePath:     "v3/system_id",
    Version:      gbfs.V30,
    DefaultTTL:   60,
    FeedHandlers: []*gbfs.FeedHandler{
        // see example for how to add feed handlers
    },
    UpdateHandler: func(s *gbfs.Server, feed gbfs.Feed, path string, err error) {
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("system=%s ttl=%d updated=%s", s.Options.SystemID, feed.GetTTL(), path)
    },
})
if err != nil {
    log.Fatal(err)
}
err = s.Start()
if err != nil {
    log.Fatal(err)
}
```

Main autodiscovery feed `gbfs.json` will be constructed from all available feeds in `FeedHandlers`. After that, all `FeedHandlers` will be regularly executed after configured `TTL`. If `TTL` is not set for individual feeds, it will be inherited from `FeedHandler`. If `TTL` is not set even for FeedHandler, `DefaultTTL` from `ServerOptions` will be used for feed.

#### Serving feeds

Feeds can be served as static files with standard webservers (Nginx, Apache, ...) or with simple built-in static file server.

```go
fs, err := gbfs.NewFileServer(gbfs.FileServerOptions{
    Addr:    "127.0.0.1:8080",
    RootDir: "public",
})
if err != nil {
    log.Fatal(err)
}
log.Fatal(fs.ListenAndServe())
```

#### Single system

For single system deployment combined with standard webserver, function `Start` provided by `Server` can be replaced with blocking function `StartAndWait`.

```go
log.Fatal(s.StartAndWait())
```

## License

Licensed under MIT license.
