package gbfs

type (
	FeedGbfsVersions struct {
		FeedCommon
		Data *FeedGbfsVersionsData `json:"data"`
	}
	FeedGbfsVersionsData struct {
		Versions []*FeedGbfsVersionsVersion `json:"versions"`
	}
	FeedGbfsVersionsVersion struct {
		Version *string `json:"version"`
		URL     *string `json:"url"`
	}
)

func (f *FeedGbfsVersions) Name() string {
	return FeedNameGbfsVersions
}
