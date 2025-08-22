package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func Get(ctx context.Context, url string) string {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	cli := &http.Client{
		Transport: transport,
	}

	resp, err := cli.Do(req)
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

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
