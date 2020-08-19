package validator

import (
	"github.com/petoc/gbfs"
	"github.com/petoc/vago/iso4217"
)

// ValidateFeedSystemPricingPlans ...
func ValidateFeedSystemPricingPlans(f *gbfs.FeedSystemPricingPlans, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.Plans == nil {
		r.ErrorW("data.plans", ErrRequired)
		return r
	}
	if len(f.Data.Plans) == 0 {
		return r
	}
	for i, s := range f.Data.Plans {
		sliceIndexName := sliceIndexN("data.plans", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.PlanID == nil {
			r.ErrorW(sliceIndexName+".plan_id", ErrRequired)
		} else if *s.PlanID == "" {
			r.ErrorW(sliceIndexName+".plan_id", ErrInvalidValue)
		}
		if s.URL != nil && (*s.Name == "" || !validateURL(s.URL)) {
			r.ErrorW(sliceIndexName+".url", ErrInvalidValue)
		}
		if s.Name == nil {
			r.ErrorW(sliceIndexName+".name", ErrRequired)
		} else if *s.Name == "" {
			r.ErrorW(sliceIndexName+".name", ErrInvalidValue)
		}
		if s.Currency == nil {
			r.ErrorW(sliceIndexName+".currency", ErrRequired)
		} else if *s.Currency == "" || !iso4217.Is(*s.Currency) {
			r.ErrorWSP(sliceIndexName+".currency", ErrInvalidValue, "ISO-4217")
		}
		if s.Price == nil {
			r.ErrorW(sliceIndexName+".price", ErrRequired)
		} else if s.Price.OldType != "" {
			r.WarningWSP(sliceIndexName+".price", ErrInvalidType, "non-negative float")
		} else if s.Price.Float64 < 0 {
			r.ErrorWSP(sliceIndexName+".price", ErrInvalidValue, "non-negative float")
		}
		if s.IsTaxable == nil {
			r.ErrorW(sliceIndexName+".is_taxable", ErrRequired)
		}
		if s.Description == nil {
			r.ErrorW(sliceIndexName+".description", ErrRequired)
		} else if *s.Description == "" {
			r.ErrorW(sliceIndexName+".description", ErrInvalidValue)
		}
	}
	return r
}
