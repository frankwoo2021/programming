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
}

func testDiv(a, b int) (int, error)  {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a/b, nil
}