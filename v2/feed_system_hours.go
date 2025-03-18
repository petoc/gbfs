package gbfs

type (
	// FeedSystemHours ...
	FeedSystemHours struct {
		FeedCommon
		Data *FeedSystemHoursData `json:"data"`
	}
	// FeedSystemHoursData ...
	FeedSystemHoursData struct {
		RentalHours []*FeedSystemHoursRentalHour `json:"rental_hours"`
	}
	// FeedSystemHoursRentalHour ...
	FeedSystemHoursRentalHour struct {
		UserTypes []string `json:"user_types"`
		Days      []string `json:"days"`
		StartTime *string  `json:"start_time"`
		EndTime   *string  `json:"end_time"`
	}
)

// Name ...
func (f *FeedSystemHours) Name() string {
	return FeedNameSystemHours
}
