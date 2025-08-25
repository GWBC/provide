package main

import (
	"context"
	"encoding/json"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/net/html"
)

type TvboxFileInfo struct {
	Props struct {
		InitialPayload struct {
			Tree struct {
				Items []struct {
					Name        string `json:"name"`
					Path        string `json:"path"`
					ContentType string `json:"contentType"`
				} `json:"items"`
			} `json:"tree"`
		} `json:"initialPayload"`
	} `json:"props"`
}

type Site struct {
	Name string `json:"name"`
	API  string `json:"api"`
}

type TvboxInfo struct {
	Sites []Site `json:"sites"`
}

func filterByCategory(str string) string {
	var result []rune
	for _, r := range str {
		if !unicode.IsSymbol(r) && !unicode.IsPunct(r) {
			result = append(result, r)
		}
	}

	return string(result)
}

func parseTvbox(url string) map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	datas := map[string]string{}

	rsp := Get(ctx, url)
	if len(rsp) == 0 {
		return datas
	}

	rsp = removeComments(rsp)

	jsonData := TvboxInfo{}
	err := json.Unmarshal([]byte(rsp), &jsonData)
	if err != nil {
		return datas
	}

	re := regexp.MustCompile(`(https?://.*api.php/provide/vod)`)

	for _, site := range jsonData.Sites {
		u := re.FindString(site.API)
		if len(u) == 0 {
			continue
		}

		name := strings.ReplaceAll(site.Name, "*", "")
		datas[u] = filterByCategory(name)
	}

	return datas
}

func Tvbox() map[string]string {
	datas := map[string]string{}

	fileCache := filepath.Join(Pwd(), "tvbox.cache")
	data, _ := os.ReadFile(fileCache)
	if len(data) != 0 {
		err := json.Unmarshal(data, &datas)
		if err == nil {
			return datas
		}
	}

	htmlData := Get(context.Background(), "https://github.com/guxiangbin/tvbox2")
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		return datas
	}

	tvboxFileInfo := TvboxFileInfo{}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, attr := range n.Attr {
				if attr.Key == "data-target" && attr.Val == "react-partial.embeddedData" {
					json.Unmarshal([]byte(n.FirstChild.Data), &tvboxFileInfo)
					if len(tvboxFileInfo.Props.InitialPayload.Tree.Items) != 0 {
						break
					}
				}
			}
		}
	}

	items := tvboxFileInfo.Props.InitialPayload.Tree.Items

	if len(items) == 0 {
		return datas
	}

	rootUrl := "https://raw.githubusercontent.com/guxiangbin/tvbox2/refs/heads/main/"
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	parallelChan := make(chan bool, 20)

	for _, item := range items {
		parallelChan <- true

		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			defer func() {
				<-parallelChan
			}()
			ds := parseTvbox(url)

			lock.Lock()
			defer lock.Unlock()
			maps.Copy(datas, ds)
		}(rootUrl + item.Name)
	}

	jsonData, _ := json.MarshalIndent(datas, "", "    ")
	os.WriteFile(fileCache, jsonData, 0644)

	return datas
}
