package localizer

// Code for this was pulled and modified from the following pages:
// 		https://www.alexedwards.net/blog/i18n-managing-translations
//    https://go.dev/blog/matchlang

// Locale Codes:
//    https://www.science.co.il/language/Locale-codes.php

import (
	// Import the internal/translations so that it's init() function
	// is run. It's really important that we do this here so that the
	// default message catalog is updated to use our translations
	// *before* we initialize the message.Printer instances below.

	_ "github.com/Route-8/chronicle-go/l10n/translations"
	"github.com/goodsign/monday"
	"github.com/xeonx/timeago"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var serverLangs = []language.Tag{
	language.AmericanEnglish, // en-US: first language is fallback
	language.German,          // de
	language.French,          // fr
	language.Spanish,         // es
	language.Swedish,         // sv
	language.Portuguese,      // pt
	language.Japanese,        // ja
	language.Hungarian,       // hu
	language.Russian,         // ru
	language.Czech,           // cs-CZ
	language.Polish,          // pl
}

// Define a Localizer type which stores the relevant locale ID (as used
// in our URLs) and a (deliberately unexported) message.Printer instance
// for the locale.
type Localizer struct {
	ID            string
	LanguageTag   language.Tag
	dateLocale    *monday.Locale
	timeagoConfig *timeago.Config
	printer       *message.Printer
}

var matcher = language.NewMatcher(serverLangs)

// The Get() function accepts a locale ID and returns the
// corresponding Localizer for that locale.
func Get(id string) Localizer {
	userTag := language.Make(id)
	tag, _, _ := matcher.Match(userTag)

	return Localizer{
		ID:          id,
		LanguageTag: tag,
		printer:     message.NewPrinter(tag),
	}
}

// We also add a Translate() method to the Localizer type. This acts
// as a wrapper around the unexported message.Printer's Sprintf()
// function and returns the appropriate translation for the given
// message and arguments.
func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
