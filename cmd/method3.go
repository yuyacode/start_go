package main

import (
	"fmt"
)

type doubleInt struct{X, Y int}
type doubleIntAddressSlice []*doubleInt

func method3Func() {
	di := make([]doubleInt, 5)
	for _, v := range di {
		fmt.Println(v.X, v.Y)
	}
	// 0 0
	// 0 0
	// 0 0
	// 0 0
	// 0 0

	s := doubleIntAddressSlice{}
	s = append(s, &doubleInt{X: 100, Y: 200})
	s = append(s, nil)
	s = append(s, &doubleInt{X: 300, Y: 400})
	res := s.ToString()
	fmt.Println(res)  // [100, 200],nil,[300, 400]
}

func (slice doubleIntAddressSlice) ToString() string {
	str := ""
	for _, s := range slice {
		if str != "" {
			str += ","
		}
		if s == nil {
			str += "nil"
		} else {
			str += fmt.Sprintf("[%d, %d]", s.X, s.Y)
		}
	}
	return str
}
