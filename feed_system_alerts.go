package gbfs

type (
	// FeedSystemAlerts ...
	FeedSystemAlerts struct {
		FeedCommon
		Data *FeedSystemAlertsData `json:"data"`
	}
	// FeedSystemAlertsData ...
	FeedSystemAlertsData struct {
		Alerts []*FeedSystemAlertsAlert `json:"alerts"`
	}
	// FeedSystemAlertsAlert ...
	FeedSystemAlertsAlert struct {
		AlertID     *string                      `json:"alert_id"`
		Type        *string                      `json:"type"`
		Times       []*FeedSystemAlertsAlertTime `json:"times,omitempty"`
		StationIDs  []string                     `json:"station_ids,omitempty"`
		RegionIDs   []string                     `json:"region_ids,omitempty"`
		URL         *string                      `json:"url,omitempty"`
		Summary     *string                      `json:"summary"`
		Description *string                      `json:"description,omitempty"`
		LastUpdated *Timestamp                   `json:"last_updated,omitempty"`
	}
	// FeedSystemAlertsAlertTime ...
	FeedSystemAlertsAlertTime struct {
		Start *Timestamp `json:"start"`
		End   *Timestamp `json:"end,omitempty"`
	}
)

// Name ...
func (f *FeedSystemAlerts) Name() string {
	return FeedNameSystemAlerts
}
