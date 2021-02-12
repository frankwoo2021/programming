package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go!")

	a, b := 4, 0
	val, err := testDiv(a, b)
	fmt.Println(val, err)

	f := testFunc(3)
	f(4)

	testDefer(10, 2)

	testSlice()

	testMap()

	testStruct()

	go testCur(1)
	go testCur(2)
	time.Sleep(time.Second * 6)

	testCP()
}

func testDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}

func testFunc(x int) func(int) {
	return func(y int) {
		fmt.Println(x + y)
	}
}

func testDefer(a, b int) {
	defer fmt.Println("dispose ...")
	fmt.Println(a / b)
}

func testSlice() {
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}

func testMap() {
	m := make(map[string]int)
	m["a"] = 5
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)
	delete(m, "a")
}

func testStruct() {
	type user struct {
		name string
		age  byte
	}
	type manager struct {
		user
		title string
	}
	var m manager
	m.name = "Tom"
	m.age = 20
	m.title = "CEO"

	fmt.Println(m)
}

func testCur(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("this is %d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func testCP() {
	consumer := func(data chan int, done chan bool) {
		for x := range data {
			fmt.Println("recieve: ", x)
		}
		done <- true
	}
	producer := func(data chan int) {
		for i := 0; i < 4; i++ {
			data <- i
		}
		close(data)
	}
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)
	<-done
}
