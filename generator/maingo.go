package main

//
//import (
//	"fmt"
//	"net/http"
//	"sync"
//)
//
//type Result struct {
//	Url        string
//	StatusCode int
//}
//
//func PingUrl(url string, wg *sync.WaitGroup) {
//
//	res, err := http.Get(url)
//	if err != nil {
//
//		print(err.Error())
//	}
//
//	result := Result{
//		Url:        url,
//		StatusCode: res.StatusCode,
//	}
//
//	wg.Done()
//
//}
//
//func GetResults(urls []string) *sync.WaitGroup { // devuelve un canal
//	var wg sync.WaitGroup
//	wg.Add(len(urls))
//
//	//ch := make(chan Result, len(urls))
//
//	for _, url := range urls {
//		go PingUrl(url, ch)
//	}
//
//	return ch
//}
//
//func main() {
//	urls := []string{
//		"https://www.google.com",
//		"https://www.facebook.com",
//		"https://www.lavoz.com.ar",
//	}
//
//	//ch := GetResults(urls)
//	wg := GetResults(urls)
//
//	//for i := 0; i < len(urls); i++ {
//	//	result := <-ch
//	//	fmt.Printf("url: %s - status_code: %d\n",
//	//		result.Url, result.StatusCode)
//	//}
//
//
//	fmt.Println("Proceso terminado.")
//
//}
