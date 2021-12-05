package main

import (
	t "awesomeProject/test"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

func ptr(s string) {
	fmt.Println(s)
}
func swi(num int) {
	if num > 3 {
		defer fmt.Println(3)
	}
	if num > 5 {
		defer fmt.Println(5)
	}
	return
}
func inc(i int) {
	defer func() {
		i++
		fmt.Println(i)
	}()
	i++
	return
}

func testChannnel() {
	for t := 0; t < 100; t++ {
		var ch chan int
		ch = make(chan int)

		go func(c <-chan int) {
			for {
				res, ok := <-c
				if !ok {
					break
				}
				fmt.Println(res)
			}
		}(ch)

		go func(c chan<- int) {
			for i := 0; i < 100; i++ {
				c <- i
			}
		}(ch)

		time.Sleep(time.Second * 1)
	}
}

func throwError() error {
	if true {
		panic("烫烫烫烫烫")
	} else {
		return nil
	}
}

func testError() (string, error) {
	if err := throwError(); err != nil {
		print("error?")
		return "error", err
	} else {
		return "not error occur", nil
	}
}

type msg struct {
	Msg  string
	Code int
}

type data struct {
	Data string
	Code int
}

func copyProperties(from interface{}, to interface{}) {
	typeF := reflect.TypeOf(from).Elem()
	valueF := reflect.ValueOf(from).Elem()
	typeT := reflect.TypeOf(to).Elem()
	valueT := reflect.ValueOf(to).Elem()
	num := typeF.NumField()
	for i := 0; i < num; i++ {
		if field, contains := typeT.FieldByName(typeF.Field(i).Name); contains {
			valueT.FieldByName(field.Name).Set(valueF.FieldByName(field.Name))
		}
	}
}

func testCopy() {
	m := msg{Msg: "3", Code: 500}
	n := data{Data: "2", Code: 300}
	copyProperties(&m, &n)
	fmt.Printf("%v", n)
}

func main() {
	regex := regexp.MustCompile("[0-9]*")
	fmt.Printf("%s", regex.Find([]byte("123as23")))

	println(t.Hello())
	//fmt.Print(regex.FindAllString("12321",-1))
}
