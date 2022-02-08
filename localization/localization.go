package localization

import (
	"encoding/json"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitLocalizer(langs ...string) *i18n.Localizer {
	// Create a new i18n bundle with English as default language.
	bundle := i18n.NewBundle(language.English)

	// Register a json unmarshal function for i18n bundle.
	// This is to enable usage of json format
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load source language
	bundle.LoadMessageFile("./locales/en.json")
	bundle.LoadMessageFile("./locales/ru.json")

	// Initialize localizer which will look for phrase keys in passed languages
	// in a strict order (first language is searched first)
	// When no key in any of the languages is found, it fallbacks to default - English language
	localizer := i18n.NewLocalizer(bundle, langs...)

	return localizer
}

func LocaleStringFromLanguage(lang string) string {
	switch strings.ToLower(lang) {
	// Russian
	case "ru":
		fallthrough
	case "russian":
		fallthrough
	case "ru_ru":
		return language.Russian.String()

	default:
		return language.English.String()

	}
}