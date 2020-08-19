package validator

import (
	"github.com/petoc/gbfs"
	"golang.org/x/text/language"
)

// ValidateFeedGbfs ...
func ValidateFeedGbfs(f *gbfs.FeedGbfs, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrZero(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	for l, lang := range f.Data {
		var err error
		_, err = language.Parse(l)
		if err != nil {
			r.ErrorW("data: language key", ErrInvalidValue)
		}
		sliceLangName := "data[" + l + "]"
		if nilOrEmpty(lang) || nilOrZero(lang.Feeds) {
			r.ErrorW(sliceLangName, ErrRequired)
			continue
		}
		for i, lf := range lang.Feeds {
			sliceIndexName := sliceLangName + sliceIndexN(".feeds", i)
			if nilOrEmpty(lf) {
				r.ErrorW(sliceIndexName, ErrRequired)
				continue
			}
			if lf.Name == nil || *lf.Name == "" {
				r.ErrorW(sliceIndexName+".name", ErrRequired)
			}
			// else {
			// 	if !ValidateFeedName(*lf.Name) {
			// 		r.WarningW(sliceIndexName+".name", ErrInvalidValue)
			// 	}
			// }
			if lf.URL == nil || *lf.URL == "" {
				r.ErrorW(sliceIndexName+".url", ErrRequired)
			} else {
				if !validateURL(lf.URL) {
					r.ErrorW(sliceIndexName+".url", ErrInvalidValue)
				}
			}
		}
	}
	return r
}
