package gbfs

type (
	FeedSystemAlerts struct {
		FeedCommon
		Data *FeedSystemAlertsData `json:"data"`
	}
	FeedSystemAlertsData struct {
		Alerts []*FeedSystemAlertsAlert `json:"alerts"`
	}
	FeedSystemAlertsAlert struct {
		AlertID     *ID                          `json:"alert_id"`
		Type        *string                      `json:"type"`
		Times       []*FeedSystemAlertsAlertTime `json:"times,omitempty"`
		StationIDs  []*ID                        `json:"station_ids,omitempty"`
		RegionIDs   []*ID                        `json:"region_ids,omitempty"`
		URL         []*LocalizedString           `json:"url,omitempty"`
		Summary     []*LocalizedString           `json:"summary"`
		Description []*LocalizedString           `json:"description,omitempty"`
		LastUpdated *Timestamp                   `json:"last_updated,omitempty"`
	}
	FeedSystemAlertsAlertTime struct {
		Start *Timestamp `json:"start"`
		End   *Timestamp `json:"end,omitempty"`
	}
)

func (f *FeedSystemAlerts) Name() string {
	return FeedNameSystemAlerts
}
