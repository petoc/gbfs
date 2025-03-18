package gbfs

type (
	FeedGbfs struct {
		FeedCommon
		Data *FeedGbfsData `json:"data"`
	}
	FeedGbfsData struct {
		Feeds []*FeedGbfsFeed `json:"feeds"`
	}
	FeedGbfsFeed struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	}
)

func (f *FeedGbfs) Name() string {
	return FeedNameGbfs
}
