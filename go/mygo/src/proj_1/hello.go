package main 
import (
	"fmt"
	"errors"
)
func main()  {
	fmt.Println("hello go!")

	a, b := 4, 0
	val, err := testDiv(a,b)
	fmt.Println(val, err)

	f := testFunc(3)
	f(4)

	testDefer(10,2)

	testSlice()

	testMap()

	testStruct()
}

func testDiv(a, b int) (int, error)  {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a/b, nil
}

func testFunc(x int) func (int)  {
	return func (y int)  {
		fmt.Println(x+y)
	}
}  

func testDefer(a, b int)  {
	defer fmt.Println("dispose ...")
	fmt.Println(a/b)
}

func testSlice()  {
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}

func testMap()  {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)
	delete(m, "a")
}

func testStruct()  {
	type user struct {
		name string
		age byte
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