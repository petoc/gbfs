package gbfs

type (
	FeedManifest struct {
		FeedCommon
		Data *FeedManifestData `json:"data"`
	}
	FeedManifestData struct {
		Datasets []*FeedManifestDataset `json:"datasets"`
	}
	FeedManifestDataset struct {
		SystemID *string                       `json:"system_id"`
		Versions []*FeedManifestDatasetVersion `json:"versions"`
	}
	FeedManifestDatasetVersion struct {
		Version *string `json:"version"`
		URL     *string `json:"url"`
	}
)

func (f *FeedManifest) Name() string {
	return FeedNameManifest
}
