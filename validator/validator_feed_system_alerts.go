package validator

import "github.com/petoc/gbfs"

// ValidateFeedSystemAlerts ...
func ValidateFeedSystemAlerts(f *gbfs.FeedSystemAlerts, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.Alerts == nil {
		r.ErrorW("data.alerts", ErrRequired)
		return r
	}
	if len(f.Data.Alerts) == 0 {
		return r
	}
	for i, s := range f.Data.Alerts {
		sliceIndexName := sliceIndexN("data.alerts", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.AlertID == nil {
			r.ErrorW(sliceIndexName+".alert_id", ErrRequired)
		} else if *s.AlertID == "" {
			r.ErrorW(sliceIndexName+".alert_id", ErrInvalidValue)
		}
		if s.Type == nil {
			r.ErrorW(sliceIndexName+".type", ErrRequired)
		} else if *s.Type == "" || !ValidateAlertType(s.Type) {
			r.ErrorW(sliceIndexName+".type", ErrInvalidValue)
		}
		if s.Times != nil && len(s.Times) > 0 {
			for si, sid := range s.Times {
				timeIndexName := sliceIndexN(sliceIndexName+".times", si)
				if nilOrEmpty(sid) {
					r.ErrorW(timeIndexName, ErrInvalidValue)
					continue
				}
				if sid.Start == nil {
					r.ErrorW(timeIndexName+".start", ErrRequired)
				} else if !validateTimestamp(*sid.Start) {
					r.ErrorW(timeIndexName+".start", ErrInvalidValue)
				}
				if sid.End != nil && !validateTimestamp(*sid.End) {
					r.ErrorW(timeIndexName+".end", ErrInvalidValue)
				}
			}

		}
		if s.StationIDs != nil && len(s.StationIDs) > 0 {
			for si, sid := range s.StationIDs {
				if sid == nil || *sid == "" {
					r.ErrorW(sliceIndexN(sliceIndexName+".station_ids", si), ErrInvalidValue)
				}
			}
		}
		if s.RegionIDs != nil && len(s.RegionIDs) > 0 {
			for si, sid := range s.RegionIDs {
				if sid == nil || *sid == "" {
					r.ErrorW(sliceIndexN(sliceIndexName+".region_ids", si), ErrInvalidValue)
				}
			}
		}
		if s.URL != nil && (*s.URL == "" || !validateURL(s.URL)) {
			r.ErrorW(sliceIndexName+".url", ErrInvalidValue)
		}
		if s.Summary == nil {
			r.ErrorW(sliceIndexName+".summary", ErrRequired)
		} else if *s.Summary == "" {
			r.ErrorW(sliceIndexName+".summary", ErrInvalidValue)
		}
		if s.Description != nil && *s.Description == "" {
			r.WarningW(sliceIndexName+".description", ErrEmptyValue)
		}
		if s.LastUpdated != nil && !validateTimestamp(*s.LastUpdated) {
			r.WarningW(sliceIndexName+".last_updated", ErrInvalidValue)
		}
	}
	return r
}
