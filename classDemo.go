package main

import (
	"fmt"
	"reflect"
)

type DataDAO interface {
	getMessage() string
	setMessage(string)
	toString() string
}
type Data struct {
	message string
}

func (this *Data) getMessage() string {
	return this.message
}

func (this *Data) setMessage(msg string) {
	this.message = msg
}
func (this *Data) toString() string {
	return "['message':'" + this.message + "']"
}

type MessageData struct {
	message string
	info    string
}

func (this *MessageData) getMessage() string {
	return this.message
}

func (this *MessageData) setMessage(message string) {
	this.message = message
}

func (this *MessageData) toString() string {
	return "['message':'" + this.message + "'," + "'info':'" + this.info + "']"
}

func (this *MessageData) setInfo(info string) {
	this.info = info
}
func (this *MessageData) getInfo() string {
	return this.info
}

func jsonStr(dao DataDAO) string {
	val := dao.toString()
	fmt.Println(val)
	return val
}

func getType(dao DataDAO) string {
	return "tp:" + reflect.TypeOf(dao).String()
}
func testClassDemo() {
	data := Data{"test"}
	data.setMessage("test data")
	fmt.Println(data.getMessage())
	jsonStr(&data)
	dataMsg := MessageData{
		message: "ok",
		info:    "info",
	}
	jsonStr(&dataMsg)

	fmt.Println(getType(&data))
	fmt.Println(getType(&dataMsg))
	panic("runtime exception")
}
