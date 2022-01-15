package times

import "time"

type Time interface {
	Now(timeGMT *int) time.Time
}
