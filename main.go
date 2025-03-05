package main

import (
	"awesomeProject3/errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	a := errors.GetValue()
	fmt.Println(a)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			get, err := http.Get("http://asdf")
			if err != nil {
				fmt.Println(fmt.Sprintf("cannot get http request %v", err))
				time.Sleep(5 * time.Second)
				continue
			}

			if get == nil || get.StatusCode != 200 {
				time.Sleep(5 * time.Second)
				continue
			}

			time.Sleep(5 * time.Hour)
		}
	}()

	wg.Wait()

}
