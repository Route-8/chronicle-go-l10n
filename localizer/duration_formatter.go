package localizer

import (
	"strings"
	"time"
)

type durationCount struct {
	Minutes int
	Hours   int
	Days    int
	Months  int
}

func (l *Localizer) TimeDuration(dur time.Duration) string {
	counter := durationCount{
		Minutes: int(dur.Minutes()) % 60,
		Hours:   int(dur.Hours()) % 24,
		Days:    int(dur.Hours()) / 24,
		Months:  int(dur.Hours()) / (24 * 30),
	}

	output := ""

	if counter.Months > 0 {
		output += l.Translate("%d months", counter.Months) + " "
	}

	if counter.Days > 0 {
		output += l.Translate("%d days", counter.Days) + " "
	}

	if counter.Hours > 0 {
		output += l.Translate("%d hours", counter.Hours) + " "
	}

	if counter.Minutes > 0 {
		output += l.Translate("%d minutes", counter.Minutes) + " "
	}

	return strings.TrimSpace(output)
}
