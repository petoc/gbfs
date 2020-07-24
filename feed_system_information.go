package gbfs

type (
	// FeedSystemInformation ...
	FeedSystemInformation struct {
		FeedCommon
		Data *FeedSystemInformationData `json:"data"`
	}
	// FeedSystemInformationData ...
	FeedSystemInformationData struct {
		SystemID                    string      `json:"system_id"`
		Language                    string      `json:"language"`
		Name                        string      `json:"name"`
		ShortName                   string      `json:"short_name,omitempty"`
		Operator                    string      `json:"operator,omitempty"`
		URL                         string      `json:"url,omitempty"`
		PurchaseURL                 string      `json:"purchase_url,omitempty"`
		StartDate                   Date        `json:"start_date,omitempty"`
		PhoneNumber                 string      `json:"phone_number,omitempty"`
		Email                       string      `json:"email,omitempty"`
		FeedContactEmail            string      `json:"feed_contact_email,omitempty"` // (v1.1)
		Timezone                    string      `json:"timezone"`
		LicenseID                   string      `json:"license_id,omitempty"`                    // (v3.0-RC)
		LicenseURL                  string      `json:"license_url,omitempty"`                   // (v3.0-RC)
		AttributionOrganizationName string      `json:"attribution_organization_name,omitempty"` // (v3.0-RC)
		AttributionURL              string      `json:"attribution_url,omitempty"`               // (v3.0-RC)
		RentalApps                  *RentalApps `json:"rental_apps,omitempty"`                   // (v1.1)
	}
)

// Name ...
func (f *FeedSystemInformation) Name() string {
	return FeedNameSystemInformation
}
