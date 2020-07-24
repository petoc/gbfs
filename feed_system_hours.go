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
		UserTypes []UserType `json:"user_types"`
		Days      []Day      `json:"days"`
		StartTime Time       `json:"start_time"`
		EndTime   Time       `json:"end_time"`
	}
)

// Name ...
func (f *FeedSystemHours) Name() string {
	return FeedNameSystemHours
}
