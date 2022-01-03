package manipulator

import (
	"strconv"
)

type PayloadNumber struct {
	NumberOne int
}

func (p *PayloadNumber) MakeNumberToBeFloat() float64 {

	numberStr := strconv.Itoa(p.NumberOne)

	if len(numberStr) < 6 {
		return 0
	}
	var result string
	for index := 0; index < 9-len(numberStr); index++ {
		result = result + "0"
	}

	resultdes := result + numberStr
	res := resultdes[:1] + "." + resultdes[1:]
	response, _ := strconv.ParseFloat(res, 64)
	return response

}
