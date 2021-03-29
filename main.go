package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Cascades(23000789.01))
}
func Cascades(num float64) (nm string) {
	str := strconv.FormatFloat(num, 'f', -1, 64)
	
	pointStr := strings.Split(str, ".")
	answer := ""
	to := len(pointStr[0])
	step := 3
	for {
		if (to - step) > 0 {
			answer = pointStr[0][to-step:to] + " " + answer
			to = to - step
		} else {
			answer = pointStr[0][0:to] + " " + answer
			break
		}

	}
	lenAns := len(answer)
	answer = answer[:lenAns-1]
	answer = answer + "." + pointStr[1]
	for {
		if answer[0] == '0' {
			answer = answer[1:lenAns]
		} else {
			break
		}
	}
	return answer
}
