package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Cascades(23000789.00))
}
func Cascades(input float64) (result string) {
	var st []string
	var strlen int
	var output []string

	ghf := (int)(input * 100)
	s := fmt.Sprintf("%f", input)
	st = strings.Split(s, ".")
	strlen = len([]rune(st[0]))

	if strlen < 17 {
		if (strlen % 3) == 0 {
			output = strings.SplitAfterN(st[0], "", strlen)
			s1 := strings.Join([]string{output[0], output[1], output[2]}, "")
			result = s1
			for i := 3; i < strlen; i += 3 {
				s1 = strings.Join([]string{output[i], output[i+1], output[i+2]}, "")
				result = result + " " + s1
			}
		} else if (strlen % 3) == 2 {
			output = strings.SplitAfterN(st[0], "", strlen)
			s1 := strings.Join([]string{output[0], output[1]}, "")
			result = s1
			for i := 2; i < strlen; i += 3 {
				s1 = strings.Join([]string{output[i], output[i+1], output[i+2]}, "")
				result = result + " " + s1
			}
		} else if (strlen % 3) == 1 {
			output = strings.SplitAfterN(st[0], "", strlen)
			s1 := strings.Join([]string{output[0]}, "")
			result = s1
			for i := 1; i < strlen; i += 3 {
				s1 = strings.Join([]string{output[i], output[i+1], output[i+2]}, "")
				result = result + " " + s1
			}
		}
		if g := ghf % 100; g != 0 {
			result += ", " + strconv.Itoa(g)
			return result
		} else {
			result += ",00"
			return result
		}
	} else {
		return ("Много букОв")
	}
}
