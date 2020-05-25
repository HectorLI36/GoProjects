package main

import (
	"fmt"
)
import "strconv"

func getDateList() []string {
	baseDate := "2020-05-%s"
	var dateList []string
	for i := 1; i <= 9; i++ {
		dateList = append(dateList, fmt.Sprintf(baseDate, "0"+strconv.Itoa(i)))
	}
	for i := 10; i <= 24; i++ {
		dateList = append(dateList, fmt.Sprintf(baseDate, strconv.Itoa(i)))
	}
	return dateList
}

func main() {
	fmt.Println("I am main")
	////riderTrigger.Run(5)
	//a := getDateList()
	//print(a)
	//riderTrigger.RiderPrint("hello!")
	print(a)
	//riderTrigger.CallProxy("http://127.0.0.1:8000/I-am-go")
}
