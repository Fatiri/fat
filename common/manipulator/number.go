package manipulator

import (
	"strconv"
)

type PayloadNumber struct {
	NumberOne int
	ListInt64 []int64
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

func (p *PayloadNumber) RemoveDuplicateValues() []int64 {
	keys := make(map[int64]bool)
	list := []int64{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range p.ListInt64 {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
