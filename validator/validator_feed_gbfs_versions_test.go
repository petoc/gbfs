package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedGbfsVersions(t *testing.T) {
	f := &gbfs.FeedGbfsVersions{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedGbfsVersions(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedGbfsVersionsData{}
	r = ValidateFeedGbfsVersions(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Versions = []*gbfs.FeedGbfsVersionsVersion{}
	r = ValidateFeedGbfsVersions(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	version := &gbfs.FeedGbfsVersionsVersion{}
	f.Data.Versions = []*gbfs.FeedGbfsVersionsVersion{version}
	r = ValidateFeedGbfsVersions(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	version.Version = gbfs.NewString("123")
	r = ValidateFeedGbfsVersions(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 2 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	version.Version = gbfs.NewString("3.0")
	version.URL = gbfs.NewString("http://127.0.0.1:8080/v3/system_id/json")
	r = ValidateFeedGbfsVersions(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
