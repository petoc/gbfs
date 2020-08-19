package validator

import "github.com/petoc/gbfs"

// Validator ...
type Validator struct{}

// Validate ...
func (v *Validator) Validate(feed gbfs.Feed, version string) *Result {
	r := ValidateFeedCommon(feed, version)
	if r.HasErrors() {
		return r
	}
	switch feed.Name() {
	case gbfs.FeedNameGbfs:
		if f, ok := feed.(*gbfs.FeedGbfs); ok {
			r = ValidateFeedGbfs(f, version)
			return r
		}
	case gbfs.FeedNameGbfsVersions:
		if f, ok := feed.(*gbfs.FeedGbfsVersions); ok {
			r = ValidateFeedGbfsVersions(f, version)
			return r
		}
	case gbfs.FeedNameVehicleTypes:
		if f, ok := feed.(*gbfs.FeedVehicleTypes); ok {
			r = ValidateFeedVehicleTypes(f, version)
			return r
		}
	case gbfs.FeedNameSystemInformation:
		if f, ok := feed.(*gbfs.FeedSystemInformation); ok {
			r = ValidateFeedSystemInformation(f, version)
			return r
		}
	case gbfs.FeedNameStationInformation:
		if f, ok := feed.(*gbfs.FeedStationInformation); ok {
			r = ValidateFeedStationInformation(f, version)
			return r
		}
	case gbfs.FeedNameStationStatus:
		if f, ok := feed.(*gbfs.FeedStationStatus); ok {
			r = ValidateFeedStationStatus(f, version)
			return r
		}
	case gbfs.FeedNameFreeBikeStatus:
		if f, ok := feed.(*gbfs.FeedFreeBikeStatus); ok {
			r = ValidateFeedFreeBikeStatus(f, version)
			return r
		}
	case gbfs.FeedNameSystemHours:
		if f, ok := feed.(*gbfs.FeedSystemHours); ok {
			r = ValidateFeedSystemHours(f, version)
			return r
		}
	case gbfs.FeedNameSystemCalendar:
		if f, ok := feed.(*gbfs.FeedSystemCalendar); ok {
			r = ValidateFeedSystemCalendar(f, version)
			return r
		}
	case gbfs.FeedNameSystemRegions:
		if f, ok := feed.(*gbfs.FeedSystemRegions); ok {
			r = ValidateFeedSystemRegions(f, version)
			return r
		}
	case gbfs.FeedNameSystemPricingPlans:
		if f, ok := feed.(*gbfs.FeedSystemPricingPlans); ok {
			r = ValidateFeedSystemPricingPlans(f, version)
			return r
		}
	case gbfs.FeedNameSystemAlerts:
		if f, ok := feed.(*gbfs.FeedSystemAlerts); ok {
			r = ValidateFeedSystemAlerts(f, version)
			return r
		}
	case gbfs.FeedNameGeofencingZones:
		if f, ok := feed.(*gbfs.FeedGeofencingZones); ok {
			r = ValidateFeedGeofencingZones(f, version)
			return r
		}
	}
	r.Error(ErrInvalidInput)
	return r
}

// ValidateAll ...
func (v *Validator) ValidateAll(feeds []gbfs.Feed, version string) *Result {
	result := &Result{}
	for _, feed := range feeds {
		r := v.Validate(feed, version)
		result.Infos = append(result.Infos, r.Infos...)
		result.Warnings = append(result.Warnings, r.Warnings...)
		result.Errors = append(result.Errors, r.Errors...)
	}
	return result
}

// New ...
func New() *Validator {
	return &Validator{}
}
