# GBFS

Go implementation of client and server for GBFS (General Bikeshare Feed Specification).

## Usage

### Client

```go
c, err := gbfs.NewClient(&gbfs.ClientOptions{
    AutoDiscoveryURL: "http://127.0.0.1:8080/v3/system_id/gbfs.json",
    DefaultLanguage:  "en",
    Logger:           log.New(os.Stdout, "", 3),
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
err := c.Subscribe(&gbfs.ClientSubscribeOptions{
    Languages: []string{"en"},
    Handler: func(c *gbfs.Client, f gbfs.Feed) {
        j, _ := json.Marshal(f)
        c.Logger.Printf("feed=%s data=%s", f.Name(), j)
    },
})
if err != nil {
    log.Println(err)
}
```

### Server

```go
s, err := gbfs.NewServer(&gbfs.ServerOptions{
    SystemID:     systemID,
    RootDir:      "public",
    BaseURL:      "http://127.0.0.1:8080",
    BasePath:     "v3/system_id",
    Version:      gbfs.V30,
    DefaultTTL:   60,
    FeedHandlers: []*gbfs.FeedHandler{
        // see example for how to add feed handlers
    },
    Logger:       log.New(os.Stdout, "", 3),
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

Server will regularly update feeds. They can be served as static files with standard webservers (Nginx, Apache, ...) or with simple built-in static file server.

```go
fs, err := gbfs.NewFileServer(&gbfs.FileServerOptions{
    Addr:    "127.0.0.1:8080",
    RootDir: "public",
    Logger:  log.New(os.Stdout, "", 3),
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