package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"
)

func check(name string, url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rsp := Get(ctx, url+"?ac=list&wd=.")
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
	domain := map[string]bool{}

	for strUrl, name := range data {
		wg.Add(1)
		go func(name string, strUrl string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			if check(name, strUrl) {
				urlObj, err := url.Parse(strUrl)
				if err != nil {
					fmt.Println("地址解析失败:", strUrl)
					return
				}

				lock.Lock()
				defer lock.Unlock()

				if domain[urlObj.Host] {
					return
				}

				ret[strUrl] = name
				domain[urlObj.Host] = true
			}
		}(name, strUrl)

		semaphore <- true
	}

	wg.Wait()

	return ret
}
