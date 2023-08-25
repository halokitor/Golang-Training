package main

import (
	hello "Golang-Training/01_HelloWorld"
	mypkg "Golang-Training/02_ImportPkg"
	variable "Golang-Training/03_Variable"
	function "Golang-Training/04_Function"
	process "Golang-Training/05_Process"
	pointerstruct "Golang-Training/06_PointerStruct"
)

func main() {
	name := "kitor"
	hello.Hello_(name)
	mypkg.PrintPkg()
	variable.PrintVariable()
	function.PrintFunc()
	process.PrintProcess()
	pointerstruct.PrintPointerStruct()
}
