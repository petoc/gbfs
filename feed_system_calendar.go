package gbfs

type (
	// FeedSystemCalendar ...
	FeedSystemCalendar struct {
		FeedCommon
		Data *FeedSystemCalendarData `json:"data"`
	}
	// FeedSystemCalendarData ...
	FeedSystemCalendarData struct {
		Calendars []*FeedSystemCalendarCalendar `json:"calendars"`
	}
	// FeedSystemCalendarCalendar ...
	FeedSystemCalendarCalendar struct {
		StartMonth int `json:"start_month"`
		StartDay   int `json:"start_day"`
		StartYear  int `json:"start_year,omitempty"`
		EndMonth   int `json:"end_month"`
		EndDay     int `json:"end_day"`
		EndYear    int `json:"end_year,omitempty"`
	}
)

// Name ...
func (f *FeedSystemCalendar) Name() string {
	return FeedNameSystemCalendar
}
