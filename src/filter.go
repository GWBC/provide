package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func check(name string, url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rsp := Get(ctx, url+"?ac=list")
	if len(rsp) == 0 {
		return false
	}

	data := map[string]any{}
	err := json.Unmarshal([]byte(rsp), &data)
	if err != nil {
		return false
	}

	_, ok := data["code"]
	fmt.Println(name, "检测:", ok)
	return ok
}

func Filter(data map[string]string) map[string]string {
	wg := sync.WaitGroup{}
	semaphore := make(chan bool, 10)

	lock := sync.Mutex{}
	ret := map[string]string{}

	for url, name := range data {
		wg.Add(1)
		go func(name string, url string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			if check(name, url) {
				lock.Lock()
				ret[url] = name
				lock.Unlock()
			}
		}(name, url)

		semaphore <- true
	}

	wg.Wait()

	return ret
}
