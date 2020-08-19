package validator

import "github.com/petoc/gbfs"

// ValidateFeedSystemRegions ...
func ValidateFeedSystemRegions(f *gbfs.FeedSystemRegions, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if nilOrZero(f.Data.Regions) {
		r.ErrorW("data.regions", ErrRequired)
		return r
	}
	for i, s := range f.Data.Regions {
		sliceIndexName := sliceIndexN("data.regions", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.RegionID == nil {
			r.ErrorW(sliceIndexName+".region_id", ErrRequired)
		} else if *s.RegionID == "" {
			r.ErrorW(sliceIndexName+".region_id", ErrInvalidValue)
		}
		if s.Name == nil {
			r.ErrorW(sliceIndexName+".name", ErrRequired)
		} else if *s.Name == "" {
			r.ErrorW(sliceIndexName+".name", ErrInvalidValue)
		}
	}
	return r
}
