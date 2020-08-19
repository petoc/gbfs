package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemRegions(t *testing.T) {
	f := &gbfs.FeedSystemRegions{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemRegions(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemRegionsData{}
	r = ValidateFeedSystemRegions(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Regions = []*gbfs.FeedSystemRegionsRegion{}
	r = ValidateFeedSystemRegions(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	region := &gbfs.FeedSystemRegionsRegion{}
	f.Data.Regions = []*gbfs.FeedSystemRegionsRegion{region}
	r = ValidateFeedSystemRegions(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	region.RegionID = gbfs.NewID("123")
	region.Name = gbfs.NewString("")
	r = ValidateFeedSystemRegions(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	region.Name = gbfs.NewString("Region")
	r = ValidateFeedSystemRegions(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
