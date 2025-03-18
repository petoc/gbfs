package gbfs

type (
	// FeedSystemRegions ...
	FeedSystemRegions struct {
		FeedCommon
		Data *FeedSystemRegionsData `json:"data"`
	}
	// FeedSystemRegionsData ...
	FeedSystemRegionsData struct {
		Regions []*FeedSystemRegionsRegion `json:"regions"`
	}
	// FeedSystemRegionsRegion ...
	FeedSystemRegionsRegion struct {
		RegionID *ID     `json:"region_id"`
		Name     *string `json:"name"`
	}
)

// Name ...
func (f *FeedSystemRegions) Name() string {
	return FeedNameSystemRegions
}
