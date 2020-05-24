package riderTrigger

import (
	"fmt"
	"time"
)

func RiderPrint(s string)  {
	fmt.Println("hello from the src!")
	fmt.Println(s)
}

func callPostRequest(url string)  {
	fmt.Println("Calling ", url)
	time.Sleep(time.Duration(1) * time.Second)
}

func timer(interval int)  {
	go callPostRequest("url1")
	fmt.Println("back")
	time.Sleep(time.Duration(interval) * time.Second)
	
	go callPostRequest("url2")
	fmt.Println("back")
	time.Sleep(time.Duration(interval) * time.Second)
	
}

func Run(interval int) {
	fmt.Println("triggering started!")
	timer(interval)
}