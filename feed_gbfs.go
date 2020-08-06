package gbfs

type (
	// FeedGbfs ...
	FeedGbfs struct {
		FeedCommon
		Data map[string]*FeedGbfsLanguage `json:"data"`
	}
	// FeedGbfsLanguage ...
	FeedGbfsLanguage struct {
		Feeds []*FeedGbfsFeed `json:"feeds"`
	}
	// FeedGbfsFeed ...
	FeedGbfsFeed struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	}
)

// Name ...
func (f *FeedGbfs) Name() string {
	return FeedNameGbfs
}
