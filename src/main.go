package main

import "mypackage"

func main() {
	println(mypackage.PublicVar)
	mypackage.PublicFunc()
	println(mypackage.FirstConstant)
}
