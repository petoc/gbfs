package gbfs

type (
	FeedSystemInformation struct {
		FeedCommon
		Data *FeedSystemInformationData `json:"data"`
	}
	FeedSystemInformationData struct {
		SystemID                    *ID                `json:"system_id"`
		Languages                   []string           `json:"languages"`
		Name                        []*LocalizedString `json:"name"`
		OpeningHours                *string            `json:"opening_hours"`
		ShortName                   []*LocalizedString `json:"short_name,omitempty"`
		Operator                    []*LocalizedString `json:"operator,omitempty"`
		URL                         *string            `json:"url,omitempty"`
		PurchaseURL                 *string            `json:"purchase_url,omitempty"`
		StartDate                   *string            `json:"start_date,omitempty"`
		TerminationDate             *string            `json:"termination_date,omitempty"`
		PhoneNumber                 *string            `json:"phone_number,omitempty"`
		Email                       *string            `json:"email,omitempty"`
		FeedContactEmail            *string            `json:"feed_contact_email,omitempty"`
		ManifestURL                 *string            `json:"manifest_url,omitempty"`
		Timezone                    *string            `json:"timezone"`
		LicenseID                   *string            `json:"license_id,omitempty"`
		LicenseURL                  *string            `json:"license_url,omitempty"`
		AttributionOrganizationName []*LocalizedString `json:"attribution_organization_name,omitempty"`
		AttributionURL              *string            `json:"attribution_url,omitempty"`
		BrandAssets                 *BrandAssets       `json:"brand_assets,omitempty"`
		TermsURL                    []*LocalizedString `json:"terms_url,omitempty"`
		TermsLastUpdated            *string            `json:"terms_last_updated"`
		PrivacyURL                  []*LocalizedString `json:"privacy_url,omitempty"`
		PrivacyLastUpdated          *string            `json:"privacy_last_updated"`
		RentalApps                  *RentalApps        `json:"rental_apps,omitempty"`
	}
	BrandAssets struct {
		BrandLastModified *string `json:"brand_last_modified"`
		BrandTermsURL     *string `json:"brand_terms_url,omitempty"`
		BrandImageURL     *string `json:"brand_image_url"`
		BrandImageURLDark *string `json:"brand_image_url_dark,omitempty"`
		Color             *string `json:"color,omitempty"`
	}
)

func (f *FeedSystemInformation) Name() string {
	return FeedNameSystemInformation
}
