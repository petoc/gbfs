package gbfs

type (
	FeedSystemRegions struct {
		FeedCommon
		Data *FeedSystemRegionsData `json:"data"`
	}
	FeedSystemRegionsData struct {
		Regions []*FeedSystemRegionsRegion `json:"regions"`
	}
	FeedSystemRegionsRegion struct {
		RegionID *ID                `json:"region_id"`
		Name     []*LocalizedString `json:"name"`
	}
)

func (f *FeedSystemRegions) Name() string {
	return FeedNameSystemRegions
}
