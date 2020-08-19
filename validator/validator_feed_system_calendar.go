package validator

import "github.com/petoc/gbfs"

// ValidateFeedSystemCalendar ...
func ValidateFeedSystemCalendar(f *gbfs.FeedSystemCalendar, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if nilOrZero(f.Data.Calendars) {
		r.ErrorW("data.calendars", ErrRequired)
		return r
	}
	for i, s := range f.Data.Calendars {
		sliceIndexName := sliceIndexN("data.calendars", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if !validateIntMonth(s.StartMonth) {
			r.ErrorW(sliceIndexName+".start_month", ErrInvalidValue)
		}
		if !validateIntDay(s.StartDay) {
			r.ErrorW(sliceIndexName+".start_day", ErrInvalidValue)
		}
		if s.StartYear != nil && (*s.StartYear < 0 || *s.StartYear > 0 && !validateIntYear(s.StartYear)) {
			r.ErrorW(sliceIndexName+".start_year", ErrInvalidValue)
		}
		if !validateIntMonth(s.EndMonth) {
			r.ErrorW(sliceIndexName+".end_month", ErrInvalidValue)
		}
		if !validateIntDay(s.EndDay) {
			r.ErrorW(sliceIndexName+".end_day", ErrInvalidValue)
		}
		if s.EndYear != nil && (*s.EndYear < 0 || *s.EndYear > 0 && !validateIntYear(s.EndYear)) {
			r.ErrorW(sliceIndexName+".end_year", ErrInvalidValue)
		}
	}
	return r
}
