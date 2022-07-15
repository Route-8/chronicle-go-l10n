package localizer

import (
	"time"

	"github.com/xeonx/timeago"
	"golang.org/x/text/language"
)

var timeagoConfigMappings = map[language.Tag]timeago.Config{
	language.AmericanEnglish: timeago.English,
	language.German:          timeago.German,
	language.French:          timeago.French,
	language.Spanish:         timeago.Spanish,
	// language.Dutch:           timeago.,
}

func (l *Localizer) TimeAgo(date time.Time) string {
	if l.timeagoConfig == nil {
		l.timeagoConfig = lookupTimeagoConfig(l.LanguageTag)
	}

	return l.timeagoConfig.Format(date)
}

func lookupTimeagoConfig(tag language.Tag) *timeago.Config {
	// Lookup date locale
	config, ok := timeagoConfigMappings[tag]
	if !ok {
		// If not found, try the parent tag
		config, ok = timeagoConfigMappings[tag.Parent()]
		if !ok {
			config = timeago.English
		}
	}

	config.Max = time.Hour * 24 * 30
	return &config
}
