package main

import "maps"

func main() {
	zy := ParseZy()
	tvbox := Tvbox()
	maps.Copy(zy, tvbox)
	zy = Filter(zy)
	GenJSFile(zy)
}
