package gbfs

type (
	// FeedGbfsVersions ...
	FeedGbfsVersions struct {
		FeedCommon
		Data *FeedGbfsVersionsData `json:"data"`
	}
	// FeedGbfsVersionsData ...
	FeedGbfsVersionsData struct {
		Versions []*FeedGbfsVersionsVersion `json:"versions"`
	}
	// FeedGbfsVersionsVersion ...
	FeedGbfsVersionsVersion struct {
		Version string `json:"version"`
		URL     string `json:"url"`
	}
)

// Name ...
func (f *FeedGbfsVersions) Name() string {
	return FeedNameGbfsVersions
}
