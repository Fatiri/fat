package times

import "time"

type Time interface {
	Now(timeGMT *int) time.Time
	TimpStampToDateStr(timeStr, layout string) string
	TimpStampToDate(timeStr, layout string) time.Time
}
