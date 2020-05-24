package riderTrigger

import (
	"fmt"
	"time"
)

// func RiderPrint(s string)  {
// 	fmt.Println("hello from the src!")
// 	fmt.Println(s)
// }

func callProxy(url string) {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("dummy calling")
}

func callPostRequest(urlChan chan string, quit chan int) {

	time.Sleep(time.Duration(1) * time.Second)
	for {
		select {
		case url := <-urlChan:

			callProxy(url)
			fmt.Println("----------------------------------------------------------------------Calling ", url)
		case <-quit:
			fmt.Println("finish callPostResuest")
			return
		}
	}
}

func oneDayArranger(interval int, dateChan chan string, myquit chan int) {

	urlList := [3]string{"url1", "url2", "url3"} // ToDo: dynamically set the url list
	for {
		select {
		case date := <-dateChan:
			fmt.Println("oneDayArranger started! for date ", date)

			go func() {
				urlChan := make(chan string)
				quit := make(chan int)
				go callPostRequest(urlChan, quit)
				for _, oneUrl := range urlList {
					onePostURL := oneUrl + "date" + date
					urlChan <- onePostURL
					fmt.Printf("ODA: sending %s\n", onePostURL)
					time.Sleep(time.Duration(1) * time.Second)
				}
				quit <- 1
			}()

		case <-myquit:
			fmt.Println("finish oneDayArranger")
		}
	}

}

func timer(interval int) {
	fmt.Println("Timer started!")
	// time.Sleep(time.Duration(interval) * time.Second)
	date := make(chan string)
	quit := make(chan int)
	datelist := [5]string{"1", "2", "3", "4", "5"}
	// datelist := []int{1, 2, 3, 4, 5}
	go oneDayArranger(1, date, quit)
	for _, v := range datelist {
		date <- v
		fmt.Println("Timer sended ", v)
		time.Sleep(time.Duration(interval) * time.Second)
	}
	quit <- 1
	fmt.Println("all finished!")

}

func Run(interval int) {
	fmt.Println("triggering started!")
	timer(interval)
}
