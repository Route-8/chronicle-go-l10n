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
	PastSuffix:   "経過",
	FuturePrefix: "",
	FutureSuffix: "前",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "1秒", Many: "%d秒"},
		{D: time.Minute, One: "1分", Many: "%d分"},
		{D: time.Hour, One: "1時間", Many: "%d時間"},
		{D: timeago.Day, One: "1日", Many: "%d日"},
		{D: timeago.Month, One: "1ヵ月", Many: "%dヵ月"},
		{D: timeago.Year, One: "1年", Many: "%d年"},
	},

	Zero:          "中",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoHungarianConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " telt el azóta",
	FuturePrefix: "",
	FutureSuffix: " múlva",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "egy másodperc", Many: "%d másodperc"},
		{D: time.Minute, One: "egy perc", Many: "%d perc"},
		{D: time.Hour, One: "egy óra", Many: "%d óra"},
		{D: timeago.Day, One: "egy nap", Many: "%d nap"},
		{D: timeago.Month, One: "egy hónap", Many: "%d hónap"},
		{D: timeago.Year, One: "egy év", Many: "%d év"},
	},

	Zero:          "most",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoRussianConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " назад",
	FuturePrefix: "через ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "одна секунда", Many: "%d секунд"},
		{D: time.Minute, One: "одна минута", Many: "%d минут"},
		{D: time.Hour, One: "один час", Many: "%d часов"},
		{D: timeago.Day, One: "один день", Many: "%d дней"},
		{D: timeago.Month, One: "один месяц", Many: "%d месяцев"},
		{D: timeago.Year, One: "один год", Many: "%d лет"},
	},

	Zero:          "ныне",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoCzechConfig = timeago.Config{
	PastPrefix:   "před ",
	PastSuffix:   "",
	FuturePrefix: "za ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "jedna sekunda", Many: "%d sekund"},
		{D: time.Minute, One: "jedna minuta", Many: "%d minut"},
		{D: time.Hour, One: "jedna hodina", Many: "%d hodin"},
		{D: timeago.Day, One: "jednoho dne", Many: "%d dní"},
		{D: timeago.Month, One: "jeden měsíc", Many: "%d měsíců"},
		{D: timeago.Year, One: "jeden rok", Many: "%d let"},
	},

	Zero:          "nyní",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoPolishConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " temu",
	FuturePrefix: "za ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "jedna sekunda", Many: "%d sekund"},
		{D: time.Minute, One: "jedna minuta", Many: "%d minut"},
		{D: time.Hour, One: "jedna godzina", Many: "%d godzin"},
		{D: timeago.Day, One: "jeden dzień", Many: "%d dni"},
		{D: timeago.Month, One: "jeden miesiąc", Many: "%d miesięcy"},
		{D: timeago.Year, One: "jeden rok", Many: "%d lata"},
	},

	Zero:          "zero",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "2006-01-02",
}

var timeagoFinnishConfig = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " sitten",
	FuturePrefix: "",
	FutureSuffix: " päästä",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "yksi sekunti", Many: "%d sekuntia"},
		{D: time.Minute, One: "yksi minuutti", Many: "%d minuuttia"},
		{D: time.Hour, One: "yksi tunti", Many: "%d tuntia"},
		{D: timeago.Day, One: "yksi päivä", Many: "%d päivää"},
		{D: timeago.Month, One: "yksi kuukausi", Many: "%d kuukautta"},
		{D: timeago.Year, One: "yksi vuosi", Many: "%d vuotta"},
	},

	Zero:          "nolla",
	Max:           30 * 24 * time.Hour,
	DefaultLayout: "02.01.2006",
}
