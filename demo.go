package main

import "fmt"

func gen() func() int {
	for i := 0; i < 10; i++ {
		return func() int {
			return i
		}
	}
	return nil
}
func testDemo() {
	//arr := [...]int32{1, 2, 3}
	//f := gen()
	//println(f())
	//println(f())
	//var app int32 = 32;

	//for idx, res := range arr {
	//	println(idx,res)
	//}
	//for i := 0; i < len(arr); i++ {
	//	println(i,arr[i])
	//}

	//values := []int{1, 2}
	//values = append(values, 1)
	//print(values[0])
	//const n = 2
	//arr := [n]int{1, 2}
	//var arrptr *[n]int
	//arrptr = &arr
	//for i := 0; i < n; i++ {
	//	println((*arrptr)[i])
	//}
	//var book Books = Books{"Go 语言", "null", "Go", 6495407}
	//changeBook(&book)
	//fmt.Println(book)
	//var av int = 3
	//change(&av)
	//print(av)
	//arr := [3]int{1, 2, 3}
	//res := append(arr[:],4)
	//println(res)
	//s := arr[1:2]
	//s[0] = 3
	//fmt.Println(s, arr)
	//println(len(arr), cap(arr))
	//println(len(s), cap(s))

	//var f func(int) int = nil
	//f = func(a int) int {
	//	return a+1
	//}
	//println(f(1))
	//var dict map[string]string = map[string]string{"key":"value"}
	//dict := map[string]int{"key": 32}
	//dict["newkey"] = 16
	//for key, value := range dict {
	//	fmt.Println(key, value)
	//}
	//v := dict["keys"]
	//println(v)
	//delete(dict,"aaaaa")
	//if _, exist := dict["key"]; exist {
	//	delete(dict, "key")
	//}
	//for key, value := range dict {
	//	fmt.Println(key, value)
	//}
	//arr := []int{1, 2, 3}
	//changeArr(arr, 2, 4)
	//for idx, value := range arr {
	//	println(idx, value)
	//}
	//dict := map[string]int{"1": 1, "2": 2}
	//changeMap(dict,"2",4)
	//for key,value := range dict {
	//	println(key,value)
	//}

	//messages := []*message{
	//	&message{title: "title", subject: "subject", context: "context"},
	//	&message{"title", "subject", "context"},
	//	&message{"title", "subject", "context"},
	//	&message{"title", "subject", "context"},
	//}
	//changeMessageTitle(messages, "t")
	//for _, m := range messages {
	//	fmt.Println(m)
	//}
	msgMap := make(map[string]message)
	msgMap["msg1"] = message{title: "title", subject: "subject", context: "context"}
	msgMap["msg2"] = message{title: "title", subject: "subject", context: "context"}
	changeMapMessage(msgMap, "msg1", message{"t", "s", "c"})
	for key, m := range msgMap {
		fmt.Println(key, m)
	}
}

type message struct {
	title   string
	subject string
	context string
}

func changeMapMessage(dict map[string]message, key string, msg message) {
	if _, exist := dict[key]; exist {
		dict[key] = msg
	}
}
func changeMessageTitle(msgs []*message, title string) {
	for _, m := range msgs {
		m.title = title
	}
}

func changeArr(arr []int, idx int, target int) {
	if idx >= len(arr) {
		println("error idx")
		return
	} else {
		arr[idx] = target
	}
}

func changeMap(dict map[string]int, key string, value int) {
	if _, exist := dict[key]; exist {
		dict[key] = value
	}
}

func change(a *int) {
	*a = 3
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func changeBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	(*book).title = "GO GO GO"
	book.author = "nil"
	book.subject = "go"
	book.book_id = 10
}
