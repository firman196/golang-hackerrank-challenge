package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConvertion(s string) string {
	am := strings.HasSuffix(s, "AM")
	pm := strings.HasSuffix(s, "PM")
	arrTime := strings.Split(s, ":")
	hh := arrTime[0]
	mm := arrTime[1]
	arrSs := strings.Split(arrTime[2], "")
	ss := arrSs[0] + arrSs[1]
	intHh, _ := strconv.Atoi(hh)

	if am && intHh == 12 {
		intHh = 0
	}

	if pm && intHh != 12 {
		intHh += 12
	}

	return fmt.Sprintf("%02s:%02s:%02s", strconv.Itoa(intHh), mm, ss)
}

/*
func main() {
	result1 := timeConvertion("12:40:22AM")
	fmt.Println(result1)
	result2 := timeConvertion("07:05:45PM")
	fmt.Println(result2)
}*/
