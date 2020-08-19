package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedGbfs(t *testing.T) {
	f := &gbfs.FeedGbfs{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = make(map[string]*gbfs.FeedGbfsLanguage)
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feedLanguage := &gbfs.FeedGbfsLanguage{}
	f.Data["invalid"] = feedLanguage
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	delete(f.Data, "invalid")
	f.Data["en"] = nil
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feedLanguage.Feeds = []*gbfs.FeedGbfsFeed{}
	f.Data["en-US"] = feedLanguage
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	delete(f.Data, "en-US")
	f.Data["en"] = feedLanguage
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feedLanguage.Feeds = append(feedLanguage.Feeds, nil)
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feedLanguageFeed := &gbfs.FeedGbfsFeed{}
	feedLanguage.Feeds = []*gbfs.FeedGbfsFeed{
		feedLanguageFeed,
	}
	r = ValidateFeedGbfs(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feedLanguageFeed.Name = gbfs.NewString("invalid")
	r = ValidateFeedGbfs(f, "")
	ee := []error{ErrRequired}
	if !hasErrors(r.Errors, ee) {
		t.Errorf("expected errors %v, got %v", ee, r.Errors)
		return
	}
	feedLanguageFeed.Name = gbfs.NewString("system_information")
	ps := []string{"ftp://", "http://", "http://.*", "//127.0.0.1:8080", "127.0.0.1:8080"}
	for _, p := range ps {
		feedLanguageFeed.URL = gbfs.NewString(p)
		r = ValidateFeedGbfs(f, "")
		if !hasError(r.Errors, ErrInvalidValue) {
			t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
			return
		}
	}
	feedLanguageFeed.URL = gbfs.NewString("http://127.0.0.1:8080/v3/system_id/en/system_information.json")
	r = ValidateFeedGbfs(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
