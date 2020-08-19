package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemInformation(t *testing.T) {
	f := &gbfs.FeedSystemInformation{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemInformationData{}
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.SystemID = gbfs.NewID("system_id")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 3 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Name = gbfs.NewString("System Name")
	f.Data.Language = gbfs.NewString("invalid")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) > 2 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.Data.Language = gbfs.NewString("en")
	r = ValidateFeedSystemInformation(f, "")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.Timezone = gbfs.NewString("*")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) > 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.Data.Timezone = gbfs.NewString("Europe/Amsterdam")
	r = ValidateFeedSystemInformation(f, "")
	if errorCount(r.Errors, ErrRequired) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.URL = gbfs.NewString("http://")
	f.Data.PurchaseURL = gbfs.NewString("http://")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 2 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.Data.URL = gbfs.NewString("http://localhost/en")
	f.Data.PurchaseURL = gbfs.NewString("http://localhost/en")
	r = ValidateFeedSystemInformation(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.StartDate = gbfs.NewString("2020-01")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.Data.StartDate = gbfs.NewString("2020-01-01")
	r = ValidateFeedSystemInformation(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.Email = gbfs.NewString("@example.com")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.Data.Email = gbfs.NewString("example@example.com")
	r = ValidateFeedSystemInformation(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.FeedContactEmail = gbfs.NewString("@example.com")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if !hasError(r.Infos, ErrAvailableFromVersion) {
		t.Errorf("expected info [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	f.Data.FeedContactEmail = gbfs.NewString("example@example.com")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Infos, ErrAvailableFromVersion) {
		t.Errorf("expected info [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	f.Version = gbfs.NewString("1.1")
	f.Data.FeedContactEmail = gbfs.NewString("example@example.com")
	r = ValidateFeedSystemInformation(f, "1.1")
	if hasError(r.Errors, ErrAvailableFromVersion) {
		t.Errorf("unexpected infos %v", r.Infos)
		return
	}
	f.Data.FeedContactEmail = nil
	f.Data.LicenseURL = gbfs.NewString("http://")
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.LicenseURL = gbfs.NewString("http://license")
	r = ValidateFeedSystemInformation(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.LicenseID = gbfs.NewString("MIT")
	f.Version = gbfs.NewString("1.0")
	r = ValidateFeedSystemInformation(f, "1.0")
	if !hasError(r.Infos, ErrAvailableFromVersion) {
		t.Errorf("unexpected info %v", r.Infos)
		return
	}
	f.Version = gbfs.NewString("3.0")
	r = ValidateFeedSystemInformation(f, "3.0")
	if !hasError(r.Errors, ErrConflict) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.LicenseURL = nil
	r = ValidateFeedSystemInformation(f, "3.0")
	if hasError(r.Errors, ErrConflict) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.LicenseID = nil
	f.Data.RentalApps = &gbfs.RentalApps{}
	f.Version = nil
	r = ValidateFeedSystemInformation(f, "")
	if !hasError(r.Infos, ErrAvailableFromVersion) {
		t.Errorf("unexpected infos %v", r.Infos)
		return
	}
	f.Data.RentalApps.Android = &gbfs.RentalApp{
		StoreURI:     gbfs.NewString("https://"),
		DiscoveryURI: gbfs.NewString("com.example.android"),
	}
	f.Data.RentalApps.IOS = &gbfs.RentalApp{
		StoreURI:     gbfs.NewString("https://"),
		DiscoveryURI: gbfs.NewString("com.example.ios"),
	}
	f.Version = gbfs.NewString("1.1")
	r = ValidateFeedSystemInformation(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) != 4 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.RentalApps.Android = &gbfs.RentalApp{
		StoreURI:     gbfs.NewString("https://play.google.com/store/apps/details?id=com.example.android"),
		DiscoveryURI: gbfs.NewString("com.example.android://"),
	}
	f.Data.RentalApps.IOS = &gbfs.RentalApp{
		StoreURI:     gbfs.NewString("https://apps.apple.com/app/application/id123456789"),
		DiscoveryURI: gbfs.NewString("com.example.ios://"),
	}
	f.Version = gbfs.NewString("1.1")
	r = ValidateFeedSystemInformation(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.Data.RentalApps.IOS = &gbfs.RentalApp{
		StoreURI:     gbfs.NewString("https://apps.apple.com/en/app/application-name/id123456789"),
		DiscoveryURI: gbfs.NewString("com.example.ios://"),
	}
	f.Version = gbfs.NewString("1.1")
	r = ValidateFeedSystemInformation(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
