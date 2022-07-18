package localizer

import (
	"time"

	"github.com/goodsign/monday"
	"golang.org/x/text/language"
)

var mondayDateMappings = map[language.Tag]string{
	language.AmericanEnglish: monday.LocaleEnUS,
	language.German:          monday.LocaleDeDE,
	language.French:          monday.LocaleFrFR,
	language.Spanish:         monday.LocaleEsES,
	language.Swedish:         monday.LocaleSvSE,
	// language.Dutch:           monday.LocaleNlNL,
}

func (l *Localizer) FormatDate(date time.Time) string {
	if l.dateLocale == nil {
		l.dateLocale = lookupDateLocale(l.LanguageTag)
	}

	return monday.Format(date, monday.FullFormatsByLocale[*l.dateLocale], *l.dateLocale)
}

func lookupDateLocale(tag language.Tag) *monday.Locale {
	// Lookup date locale
	localeString, ok := mondayDateMappings[tag]
	if !ok {
		// If not found, try the parent tag
		localeString, ok = mondayDateMappings[tag.Parent()]
		if !ok {
			localeString = monday.LocaleEnUS
		}
	}

	locale := monday.Locale(localeString)
	return &locale
}
