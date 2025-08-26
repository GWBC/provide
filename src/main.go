package main

import (
	"flag"
	"maps"
)

func main() {
	enableTvbox := flag.Bool("tvbox", false, "抓取tvbox")
	flag.Parse()

	zy := ParseZy()

	if *enableTvbox {
		tvbox := Tvbox()
		maps.Copy(zy, tvbox)
	}

	zy = Filter(zy)
	GenJSFile(zy)
}
