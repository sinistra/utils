package utils

import (
	"time"

	"golang.org/x/text/message"
)

const (
	defaultLanguage = "en"
	defaultTimeZone = "Australia/Sydney" // IANA time zone name
)

// AddThousandsSeparator formats an integer to a string with the thousands' separator.
func AddThousandsSeparator(n int64) string {
	p := message.NewPrinter(message.MatchLanguage(defaultLanguage))
	return p.Sprint(n)
}

// FormatTime into hh:mm XM.
func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("3:04 PM")
}

// FormatDate into dd/mm/yyyy.
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02/01/2006")
}

// FormatDateTimeInDefaultZone formats a timestamp to "hh:mm XM, dd/mm/yyyy" in the default timezone.
func FormatDateTimeInDefaultZone(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	loc, err := time.LoadLocation(defaultTimeZone)
	if err != nil {
		return ""
	}
	return t.In(loc).Format("3:04 PM, 02/01/2006 MST")
}

// FormatBool to "Yes" or "No".
func FormatBool(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
