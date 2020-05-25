package riderTrigger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

func RiderPrint(s string) {
	fmt.Println("hello from the src!")
	fmt.Println(s)
}

func callProxy(url string) {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("\tCalling API %s", url)
	res, err := http.Post(url, "application/json;charset=utf-8", nil) // will make the call
	if err != nil {
		fmt.Println("Calling got FATAL error", err.Error())
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Calling got FATAL error", err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	fmt.Println(*str)

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

func oneDayArranger(interval int, dateChan chan string, myquit chan int) {

	for {
		select {
		case date := <-dateChan:
			fmt.Println("oneDayArranger started! for date ", date)

			go func() {
				urlChan := make(chan string)
				quit := make(chan int)
				go callPostRequest(urlChan, quit)

				// start to send
				// xiashuo
				urlChan <- fmt.Sprintf(urlList[0], date)
				time.Sleep(time.Duration(30) * time.Minute)

				//3dm
				urlChan <- fmt.Sprintf(urlList[1], date)
				time.Sleep(time.Duration(30) * time.Minute)
				//youxia
				urlChan <- fmt.Sprintf(urlList[2], date)
				time.Sleep(time.Duration(3) * time.Hour)

				//rider
				urlChan <- fmt.Sprintf(urlList[3], date)
				time.Sleep(time.Duration(10) * time.Second)

				//
				//for _, oneUrl := range urlList {
				//	onePostURL := oneUrl + "date" + date  // ToDo: make the url here, and set different intervals
				//	urlChan <- onePostURL
				//	fmt.Printf("ODA: sending %s\n", onePostURL)
				//	time.Sleep(time.Duration(1) * time.Second)
				//}
				quit <- 1
				fmt.Printf("Finish oneDayArranger for one date %s\n", date)
			}()

		case <-myquit:
			fmt.Println("Finish oneDayArranger")
		}
	}

}

func timer(interval int) {
	fmt.Println("Timer started!")
	// time.Sleep(time.Duration(interval) * time.Second)
	date := make(chan string)
	quit := make(chan int)
	datelist := getDateList()
	// datelist := []int{1, 2, 3, 4, 5}
	go oneDayArranger(1, date, quit)
	for _, v := range datelist {
		date <- v
		fmt.Printf("Starting to wash data for date %s \n", v)
		time.Sleep(time.Duration(interval) * time.Hour)
	}
	quit <- 1
	fmt.Println("all finished!")

}

func Run(interval int) {
	fmt.Println("triggering started!")
	timer(interval)
}
