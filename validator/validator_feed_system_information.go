package validator

import (
	"github.com/petoc/gbfs"
	"github.com/petoc/vago/tz"
	"golang.org/x/text/language"
)

// ValidateFeedSystemInformation ...
func ValidateFeedSystemInformation(f *gbfs.FeedSystemInformation, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.SystemID == nil || *f.Data.SystemID == "" {
		r.ErrorW("data.system_id", ErrRequired)
	}
	if f.Data.Language == nil || *f.Data.Language == "" {
		r.ErrorW("data.language", ErrRequired)
	} else {
		_, err := language.Parse(*f.Data.Language)
		if err != nil {
			r.ErrorW("data.language", ErrInvalidValue)
		}
	}
	if f.Data.Name == nil || *f.Data.Name == "" {
		r.ErrorW("data.name", ErrRequired)
	}
	if f.Data.ShortName != nil && *f.Data.ShortName == "" {
		r.ErrorW("data.short_name", ErrInvalidValue)
	}
	if f.Data.URL != nil && (*f.Data.URL == "" || !validateURL(f.Data.URL)) {
		r.ErrorW("data.url", ErrInvalidValue)
	}
	if f.Data.PurchaseURL != nil && (*f.Data.PurchaseURL == "" || !validateURL(f.Data.PurchaseURL)) {
		r.ErrorW("data.purchase_url", ErrInvalidValue)
	}
	if f.Data.StartDate != nil && (*f.Data.StartDate == "" || !validateDate(f.Data.StartDate)) {
		r.ErrorW("data.start_date", ErrInvalidValue)
	}
	if f.Data.Email != nil && (*f.Data.Email == "" || !validateEmail(f.Data.Email)) {
		r.ErrorW("data.email", ErrInvalidValue)
	}
	if (version == "" || verLT(version, gbfs.V11)) && f.Data.FeedContactEmail != nil {
		r.InfoWV("data.feed_contact_email", ErrAvailableFromVersion, gbfs.V11)
	}
	if f.Data.FeedContactEmail != nil && (*f.Data.FeedContactEmail == "" || !validateEmail(f.Data.FeedContactEmail)) {
		r.ErrorW("data.feed_contact_email", ErrInvalidValue)
	}
	if f.Data.Timezone == nil || *f.Data.Timezone == "" {
		r.ErrorW("data.timezone", ErrRequired)
	} else {
		// var err error
		// _, err = time.LoadLocation(*f.Data.Timezone)
		// if err != nil {
		if !tz.Is(*f.Data.Timezone) {
			r.ErrorW("data.timezone", ErrInvalidValue)
		}
	}
	if (version == "" || verLT(version, gbfs.V30)) && f.Data.LicenseID != nil {
		r.InfoWV("data.license_id", ErrAvailableFromVersion, gbfs.V30)
	}
	if f.Data.LicenseID != nil && *f.Data.LicenseID == "" {
		r.ErrorW("data.license_id", ErrInvalidValue)
	}
	if f.Data.LicenseURL != nil && (*f.Data.LicenseURL == "" || !validateURL(f.Data.LicenseURL)) {
		r.ErrorW("data.license_url", ErrInvalidValue)
	}
	if verGE(version, gbfs.V30) && f.Data.LicenseID != nil && f.Data.LicenseURL != nil {
		r.ErrorW("data.license_id, data.license_url", ErrConflict)
	}
	if (version == "" || verLT(version, gbfs.V30)) && f.Data.AttributionOrganizationName != nil {
		r.InfoWV("data.attribution_organization_name", ErrAvailableFromVersion, gbfs.V30)
	}
	if (version == "" || verLT(version, gbfs.V30)) && f.Data.AttributionURL != nil {
		r.InfoWV("data.attribution_url", ErrAvailableFromVersion, gbfs.V30)
	}
	if f.Data.AttributionURL != nil && (*f.Data.AttributionURL == "" || !validateURL(f.Data.AttributionURL)) {
		r.ErrorW("data.attribution_url", ErrInvalidValue)
	}
	if f.Data.RentalApps != nil {
		if version == "" || verLT(version, gbfs.V11) {
			r.InfoWV("data.rental_apps", ErrAvailableFromVersion, gbfs.V11)
		} else {
			// TODO: check rental_uris constraint
			if f.Data.RentalApps.Android != nil {
				if f.Data.RentalApps.Android.StoreURI == nil || *f.Data.RentalApps.Android.StoreURI == "" || !validateStoreURIAndroid(f.Data.RentalApps.Android.StoreURI) {
					r.ErrorW("data.rental_apps.android.store_uri", ErrInvalidValue)
				}
				if f.Data.RentalApps.Android.DiscoveryURI == nil || *f.Data.RentalApps.Android.DiscoveryURI == "" || !validateDiscoveryURI(f.Data.RentalApps.Android.DiscoveryURI) {
					r.ErrorW("data.rental_apps.android.discovery_uri", ErrInvalidValue)
				}
			}
			if f.Data.RentalApps.IOS != nil {
				if f.Data.RentalApps.IOS.StoreURI == nil || *f.Data.RentalApps.IOS.StoreURI == "" || !validateStoreURIIOS(f.Data.RentalApps.IOS.StoreURI) {
					r.ErrorW("data.rental_apps.ios.store_uri", ErrInvalidValue)
				}
				if f.Data.RentalApps.IOS.DiscoveryURI == nil || *f.Data.RentalApps.IOS.DiscoveryURI == "" || !validateDiscoveryURI(f.Data.RentalApps.IOS.DiscoveryURI) {
					r.ErrorW("data.rental_apps.ios.discovery_uri", ErrInvalidValue)
				}
			}
		}
	}
	return r
}
