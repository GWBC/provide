package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"
)

func check(name string, strUrl string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := url.Values{}
	params.Add("ac", "videolist")
	params.Add("wd", "侏罗纪")

	strUrl += "?" + params.Encode()
	rsp := Get(ctx, strUrl, 2)
	if len(rsp) == 0 {
		return false
	}

	data := map[string]any{}
	err := json.Unmarshal([]byte(rsp), &data)
	if err != nil {
		return false
	}

	_, ok := data["code"]
	if ok {
		count := AnyToNumber(data["total"])
		if count > 200 {
			ok = false
		}
	} else {
		fmt.Println(name, strUrl, "数据错误")
	}

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
