package localizer

import (
	"time"

	"github.com/xeonx/timeago"
)

var timeagoSwedishConfig = timeago.Config{
	PastPrefix:   "för ",
	PastSuffix:   " sedan",
	FuturePrefix: "om ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "en sekund", Many: "%d sekunder"},
		{D: time.Minute, One: "en minut", Many: "%d minuter"},
		{D: time.Hour, One: "en timme", Many: "%d timmar"},
		{D: timeago.Day, One: "en dag", Many: "%d dagar"},
		{D: timeago.Month, One: "en månad", Many: "%d månader"},
		{D: timeago.Year, One: "ett år", Many: "%d år"},
	},

	Zero:          "en sekund",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoJapaneseConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   "前",
	FuturePrefix: "",
	FutureSuffix: "で",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "1秒", Many: "%d秒"},
		{D: time.Minute, One: "1分", Many: "%d分"},
		{D: time.Hour, One: "1時間", Many: "%d時間"},
		{D: timeago.Day, One: "1日", Many: "%d日"},
		{D: timeago.Month, One: "1ヵ月", Many: "%dヵ月"},
		{D: timeago.Year, One: "1年", Many: "%d年"},
	},

	Zero:          "今",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}
