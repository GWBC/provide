package main

import (
	"flag"
	"maps"
)

func main() {
	isTvbox := flag.Bool("tvbox", false, "抓取tvbox")
	zy := ParseZy()

	if *isTvbox {
		tvbox := Tvbox()
		maps.Copy(zy, tvbox)
	}

	zy = Filter(zy)
	GenJSFile(zy)
}
