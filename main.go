package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(Cascades(23000000.00))
	fmt.Println(Randoms(5, 10))
}
func Cascades(num float64) (nm string) {
	numInt := int(num)
	num1 := num
	if num == float64(numInt) {
		num1 += 0.01
	}
	str := strconv.FormatFloat(num1, 'f', -1, 64)

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
	if num == float64(numInt) {
		pointStr[1] = "00"
	}
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

func Randoms(num int, quantity int) []float64 {
	var afterPoint []int
	var beforePoint []int
	for i := 0; i < quantity; i++ {
		beforePoint = append(beforePoint, rand.Intn(int(math.Pow10(num))))
		afterPoint = append(afterPoint, rand.Intn(int(math.Pow10(2))))
	}
	var answer []float64
	for i := 0; i < quantity; i++ {
		answer = append(answer, (float64(afterPoint[i])/100 + float64(beforePoint[i])))

	}
	return answer

}
