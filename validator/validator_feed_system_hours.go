package validator

import "github.com/petoc/gbfs"

// ValidateFeedSystemHours ...
func ValidateFeedSystemHours(f *gbfs.FeedSystemHours, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if nilOrZero(f.Data.RentalHours) {
		r.ErrorW("data.rental_hours", ErrRequired)
		return r
	}
	for i, s := range f.Data.RentalHours {
		sliceIndexName := sliceIndexN("data.rental_hours", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if nilOrZero(s.UserTypes) {
			r.ErrorW(sliceIndexName+".user_types", ErrRequired)
		} else {
			for j, u := range s.UserTypes {
				if !ValidateUserType(u) {
					r.ErrorW(sliceIndexName+sliceIndexN(".user_types", j), ErrInvalidValue)
				}
			}
		}
		if nilOrZero(s.Days) || len(s.Days) > 7 {
			r.ErrorW(sliceIndexName+".days", ErrRequired)
		} else {
			for j, d := range s.Days {
				if !ValidateDay(d) {
					r.ErrorW(sliceIndexName+sliceIndexN(".days", j), ErrInvalidValue)
				}
			}
		}
		if s.StartTime == nil || *s.StartTime == "" {
			r.ErrorW(sliceIndexName+".start_time", ErrRequired)
		} else {
			if !validateTime(s.StartTime) {
				r.ErrorW(sliceIndexName+".start_time", ErrInvalidValue)
			}
		}
		if s.EndTime == nil || *s.EndTime == "" {
			r.ErrorW(sliceIndexName+".end_time", ErrRequired)
		} else {
			if !validateTime(s.EndTime) {
				r.ErrorW(sliceIndexName+".start_time", ErrInvalidValue)
			}
		}
	}
	return r
}
