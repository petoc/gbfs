package validator

import "github.com/petoc/gbfs"

// ValidateFeedGbfsVersions ...
func ValidateFeedGbfsVersions(f *gbfs.FeedGbfsVersions, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if nilOrZero(f.Data.Versions) {
		r.ErrorW("data.versions", ErrRequired)
		return r
	}
	for i, s := range f.Data.Versions {
		sliceIndexName := sliceIndexN("data.versions", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if !ValidateVersion(*s.Version) {
			r.ErrorW(sliceIndexName+".version", ErrInvalidValue)
		}
		if !validateURL(s.URL) {
			r.ErrorW(sliceIndexName+".url", ErrInvalidValue)
		}
	}
	return r
}
