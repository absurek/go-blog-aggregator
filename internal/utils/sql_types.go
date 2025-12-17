package utils

import (
	"database/sql"
	"time"
)

// some common layouts
var layouts = [...]string{
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

const printLayout = time.RFC3339

func ParseNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{
			Valid: false,
		}
	}

	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

func ParseNullTime(str string) sql.NullTime {
	for _, layout := range layouts {
		if t, err := time.Parse(layout, str); err == nil {
			return sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
	}

	return sql.NullTime{
		Valid: false,
	}
}

func FormatNullTime(nt sql.NullTime) string {
	if !nt.Valid {
		return ""
	}

	return nt.Time.Format(printLayout)
}
