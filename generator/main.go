package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Url        string
	StatusCode int
}

func PingUrl(url string, ch chan Result) {

	res, err := http.Get(url)
	if err != nil {

		print(err.Error())
	}

	result := Result{
		Url:        url,
		StatusCode: res.StatusCode,
	}

	ch <- result

}

func GetResults(urls []string) <-chan Result { // devuelve un canal

	ch := make(chan Result, len(urls))

	for _, url := range urls {
		go PingUrl(url, ch)
	}

	return ch
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.lavoz.com.ar",
	}

	ch := GetResults(urls)

	for i := 0; i < len(urls); i++ {
		result := <-ch
		fmt.Printf("url: %s - status_code: %d\n",
			result.Url, result.StatusCode)
	}
	fmt.Println("Proceso terminado.")

}
