package localizer

import (
	"time"

	"github.com/xeonx/timeago"
)

var timeagoSwedishConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " sedan",
	FuturePrefix: "på ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "en sekund", Many: "%d sekunder"},
		{D: time.Minute, One: "en minut", Many: "%d minuter"},
		{D: time.Hour, One: "en timme", Many: "%d timmar"},
		{D: timeago.Day, One: "en dag", Many: "%d dagar"},
		{D: timeago.Month, One: "en månad", Many: "%d månad"},
		{D: timeago.Year, One: "ett år", Many: "%d år"},
	},

	Zero:          "",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}
