# GBFS

Go implementation of client, server for GBFS (General Bikeshare Feed Specification).

## Usage

### Client

```go
import "github.com/petoc/gbfs/v2"
```

```go
c, err := gbfs.NewClient(gbfs.ClientOptions{
    AutoDiscoveryURL: "http://127.0.0.1:8080/v2/system_id/gbfs.json",
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
    log.Printf("feed=%s system_id=%s", f.Name(), *f.Data.SystemID)
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
import "github.com/petoc/gbfs/v2"
```

```go
s, err := gbfs.NewServer(gbfs.ServerOptions{
    SystemID:     "system_id",
    RootDir:      "public",
    BaseURL:      "http://127.0.0.1:8080",
    BasePath:     "v2/system_id",
    Version:      gbfs.V21,
    DefaultTTL:   60,
    FeedHandlers: []*gbfs.FeedHandler{
        // see example for how to add feed handlers
    },
    UpdateHandler: func(s *gbfs.Server, feed gbfs.Feed, path string, err error) {
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("system=%s ttl=%d version=%s updated=%s", s.Options.SystemID, feed.GetTTL(), feed.GetVersion(), path)
    },
})
if err != nil {
    log.Fatal(err)
}
log.Fatal(s.Start())
```

Main autodiscovery feed `gbfs.json` will be constructed from all available feeds in `FeedHandlers`. After that, all `FeedHandlers` will be regularly executed after configured `TTL`. If `TTL` is not set for individual feeds, it will be inherited from `FeedHandler`. If `TTL` is not set even for FeedHandler, `DefaultTTL` from `ServerOptions` will be used for feed.

#### Serving feeds

Feeds can be served as static files with standard webservers (Nginx, Apache, ...) or with simple built-in static file server.

```go
fs, err := gbfs.NewFileServer("127.0.0.1:8080", "public")
if err != nil {
    log.Fatal(err)
}
log.Fatal(fs.ListenAndServe())
```

## License

Licensed under MIT license.
