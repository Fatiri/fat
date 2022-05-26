package times

import (
	"strconv"
	"time"
)

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

	location, _ := time.LoadLocation("Asia/Jakarta")

	return time.Now().In(location).Add(time.Hour * time.Duration(newTimeGMT))
}

func (t *timesCustomImpl) TimpStampToDateStr(timeStr, layout string) string {

	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)

	return tm.Format(layout)
}

func (t *timesCustomImpl) TimpStampToDate(timeStr, layout string) time.Time {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)

	tmStr := tm.Format(layout)

	parsed, _ := time.Parse(layout, tmStr)
	return parsed
}
