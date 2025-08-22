package main

func main() {
	zy := ParseZy()
	zy = Filter(zy)
	GenJSFile(zy)
}
