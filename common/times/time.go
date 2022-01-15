package times

import "time"

type timesCustomImpl struct {
	gmt int
}

func ProvideNewTimesCustom() Time {
	// default GMT +7 (Asia/Jakarta)
	return &timesCustomImpl{gmt: 7}
}

func (t *timesCustomImpl) Now(timeGMT *int) time.Time {
	newTimeGMT := t.gmt
	if timeGMT != nil {
		newTimeGMT = *timeGMT
	}

	return time.Now().Add(time.Hour * time.Duration(newTimeGMT))
}
