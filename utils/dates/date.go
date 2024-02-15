package dates

import (
	"strings"
	"time"
)

func ToFormat(date string, from string, to string) string {
	t, _ := time.Parse(DateFormat(from), date)

	return t.Format(DateFormat(to))
}

func DateFormat(f string) string {
	f = strings.Replace(f, "YYYY", "2006", 1)
	f = strings.Replace(f, "MM", "01", 1)
	f = strings.Replace(f, "DD", "02", 1)
	f = strings.Replace(f, "HH", "15", 1)
	f = strings.Replace(f, "mm", "04", 1)
	f = strings.Replace(f, "ss", "05", 1)

	return f
}
